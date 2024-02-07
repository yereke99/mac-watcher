package domain

import "time"

type Cloud struct {
	IP        string
	ID        int
	Name      string
	Type      string
	Status    string // статус из бд
	State     string // статус пинга
	Reboot    bool
	EventTime time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type SendMessagePayload struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

type Command struct {
	Command string `json:"command"`
	Macs    []int  `json:"macs"`
}
