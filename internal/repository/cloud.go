package repository

import (
	"database/sql"
	"mac-watcher/internal/domain"
)

type Cloud struct {
	Client *sql.DB
}

func NewCloud(client *sql.DB) *Cloud {

	return &Cloud{
		Client: client,
	}
}

func (r *Cloud) GetListClouds() ([]*domain.Cloud, error) {

	query := `SELECT ip, id, cloud_name, cloud_type, cloud_status, cloud_state 
				FROM servers 
				WHERE cloud_type = "device_worker" 
  				AND cloud_state = "online"`

	rows, err := r.Client.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clouds []*domain.Cloud
	for rows.Next() {
		var cloud domain.Cloud

		if err := rows.Scan(&cloud.IP, &cloud.ID, &cloud.Name, &cloud.Type, &cloud.Status, &cloud.State); err != nil {
			return nil, err
		}

		clouds = append(clouds, &cloud)
	}
	return clouds, nil
}

func (r *Cloud) SetCloudStatus(status, state string, mac int) error {

	query := "UPDATE servers SET status = ?, state = ? WHERE id = ?"

	_, err := r.Client.Exec(query, status, state, mac)
	if err != nil {
		return err
	}

	return nil
}
