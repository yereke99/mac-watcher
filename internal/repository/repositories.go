package repository

import (
	"database/sql"
	"mac-watcher/internal/domain"
)

type ICloud interface {
	GetListClouds() ([]*domain.Cloud, error)
	SetCloudStatus(status, state string, mac int) error
}

type Repositories struct {
	Database ICloud
}

func NewRepositories(db *sql.DB) *Repositories {

	return &Repositories{
		Database: NewCloud(db),
	}
}
