package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mac-watcher/config"
	"mac-watcher/internal/domain"
	"mac-watcher/internal/repository"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/go-ping/ping"
	"go.uber.org/zap"
)

type MacWatcherService struct {
	ctx          context.Context
	appConfig    *config.Config
	zapLogger    *zap.Logger
	repositories *repository.Repositories
	rebootChan   chan int
}

func NewMacWatcherService(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repositories *repository.Repositories) *MacWatcherService {

	zapLogger.Info("started service")

	macWatcherService := &MacWatcherService{
		ctx:          ctx,
		appConfig:    appConfig,
		zapLogger:    zapLogger,
		repositories: repositories,
		rebootChan:   make(chan int),
	}

	clouds, err := macWatcherService.repositories.Database.GetListClouds()
	if err != nil {
		return nil
	}

	go func(clouds []*domain.Cloud) {
		for {
			if macWatcherService.ctx.Err() != nil {
				return
			}
			macWatcherService.listenMacReboot(clouds)
		}
	}(clouds)

	go func(clouds []*domain.Cloud) {
		for {
			if macWatcherService.ctx.Err() != nil {
				return
			}
			macWatcherService.listenMacRecovery(clouds)
		}
	}(clouds)
	return macWatcherService
}

func (s *MacWatcherService) listenMacReboot(clouds []*domain.Cloud) {

	time.Sleep(20 * time.Minute)

	for _, cloud := range clouds {
		if err := s.macReboot(cloud.ID); err != nil {
			s.zapLogger.Error("error in reboot mac", zap.Error(err))
			continue
		}
		cloud.Reboot = true
	}

	time.Sleep(2 * time.Minute)

	for _, cloud := range clouds {

		// TODO: Why?
		if err := s.sendCommandRequest("/opt/homebrew/bin/adb devices && sudo launchctl unload -w /Library/LaunchDaemons/cloud.plist && sudo launchctl load -w /Library/LaunchDaemons/cloud.plist", cloud.ID); err != nil {
			s.zapLogger.Error("error in send request", zap.Error(err))
			continue
		}

		if err := s.sendCommandRequestToOCP("sudo launchctl unload -w /Library/LaunchDaemons/ImageProcessor.plist && sudo launchctl load -w /Library/LaunchDaemons/ImageProcessor.plist", domain.OCPClouds); err != nil {
			s.zapLogger.Error("error in send request to OCP", zap.Error(err))
			continue
		}
		cloud.Reboot = false
	}
}

func (s *MacWatcherService) listenMacRecovery(clouds []*domain.Cloud) {

	time.Sleep(10 * time.Minute)

	for _, cloud := range clouds {
		if cloud.Reboot {
			continue
		}
		if err := s.macRecovery(cloud.ID); err != nil {
			s.zapLogger.Error("error in recovery mac", zap.Error(err))
			continue
		}
	}
}

func (s *MacWatcherService) sendMessage(message string) error {

	messagePaylod := domain.SendMessagePayload{
		Channel: s.appConfig.ChannelName,
		Message: message,
	}

	payload, err := json.Marshal(messagePaylod)
	if err != nil {
		return err
	}

	resp, err := http.Post(s.appConfig.SlackBotUrl+"/v1/slack/send-message", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.ErrStatusCode
	}

	return nil
}

func (s *MacWatcherService) macRecovery(mac int) error {

	if mac%2 == 0 {
		if err := s.sendCommandRequest("sudo launchctl unload -w /Library/LaunchDaemons/cloud.plist", mac); err != nil {
			return err
		}

		if err := s.sendCommandRequest("/bin/sh/ Users/badmin/reboothubs.sh", mac); err != nil {
			if err.Error() == "No hubs found" {
				return domain.ErrNoHubs
			} else {
				return err
			}
		}

		if err := s.sendCommandRequest("sudo launchctl load -w /Library/LaunchDaemons/cloud.plist", mac); err != nil {
			return err
		}

		return nil
	}

	if err := s.sendCommandRequest("sudo launchctl unload -w /Library/LaunchDaemons/cloud.plist", mac); err != nil {
		return err
	}

	if err := s.sendCommandRequest("/opt/homebrew/bin/adb kill-server", mac); err != nil {
		s.rebootChan <- mac
		return err
	}

	if err := s.sendCommandRequest("/bin/sh/ Users/badmin/reboothubs.sh", mac); err != nil {
		if err.Error() == "No hubs found" {
			return domain.ErrNoHubs
		} else {
			return err
		}
	}

	if err := s.sendCommandRequest("/opt/homebrew/bin/adb start-server", mac); err != nil {
		if err.Error() == "no devices/emulators found" {
			return domain.ErrNoDevices
		} else {
			return err
		}
	}

	if err := s.sendCommandRequest("/opt/homebrew/bin/adb devices", mac); err != nil {
		return err
	}

	if err := s.sendCommandRequest("sudo launchctl load -w /Library/LaunchDaemons/cloud.plist", mac); err != nil {
		return err
	}

	return nil
}

func (s *MacWatcherService) sendCommandRequest(command string, mac int) error {

	curlCommand := fmt.Sprintf("curl --location 'http://192.168.0.106:7071/cloud-control/command' --header 'Content-Type: application/json' --data '{\"command\":\"%s\",\"macs\":[%d]}'", command, mac)

	cmdArgs := []string{"bash", "-c", curlCommand}

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (s *MacWatcherService) sendCommandRequestToOCP(command string, macs []int) error {

	curlCommand := fmt.Sprintf("curl --location 'http://192.168.0.106:7071/cloud-control/command' --header 'Content-Type: application/json' --data '{\"command\":\"%s\",\"macs\":%v}'", command, macs)

	cmdArgs := []string{"bash", "-c", curlCommand}

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (s *MacWatcherService) macReboot(mac int) error {

	curlCommand := fmt.Sprintf("curl --location 'http://192.168.0.106:7071/cloud-control/command' --header 'Content-Type: application/json' --data '{\"command\":\"sudo shutdown -r now\",\"macs\":[%d]}'", mac)

	cmdArgs := []string{"bash", "-c", curlCommand}

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (s *MacWatcherService) macPing(cloud *domain.Cloud) error {

	pinger, err := ping.NewPinger(cloud.IP)
	if err != nil {
		return err
	}

	pinger.Count = 3
	pinger.Timeout = 5 * time.Second

	pinger.OnRecv = func(pkt *ping.Packet) {
		s.zapLogger.Info("received from", zap.Any("addr", pkt.IPAddr), zap.Any("rtt", pkt.Rtt))
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		if stats.PacketsRecv == 0 {
			s.zapLogger.Info("no response, server might be down.")
		} else {
			s.zapLogger.Info("Ping statistics: %d packets transmitted, %d packets received, %v%% packet loss", zap.Any("packetSent", stats.PacketsSent))
		}
	}

	if err = pinger.Run(); err != nil {
		return err
	}

	return nil
}
