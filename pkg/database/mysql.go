package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"mac-watcher/config"
	"mac-watcher/internal/domain"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

var (
	sqls = []string{"./migration/000002_init_schema.up.sql", "./migration/000001_init_schema.up.sql"}
)

func NewDatabase(databaseConfig *config.DatabaseConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to MySQL database")

	return db, nil
}

func InitCloudsTable(db *sql.DB) error {
	// SQL statement to create the "clouds" table
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS clouds (
			id INT AUTO_INCREMENT PRIMARY KEY,
			ip VARCHAR(255),
			cloud_name VARCHAR(255),
			cloud_type VARCHAR(255),
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			deleted_at TIMESTAMP
		);
	`

	// Execute the SQL statement to create the table
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	log.Println("created table")

	return nil
}

func InsertCloudsData(db *sql.DB) error {
	data := []domain.Cloud{
		{IP: "192.168.0.106", Name: "dcdell3", Type: "handler", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.12", Name: "dcdell4", Type: "handler", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.129", Name: "dcdell5", Type: "handler", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.207", Name: "dcmac1", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.157", Name: "dcmac2", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.144", Name: "dcmac3", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.55", Name: "dcmac4", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.138", Name: "dcmac5", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.62", Name: "dcmac6", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.100", Name: "dcmac7", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.162", Name: "dcmac8", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.112", Name: "dcmac9", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.195", Name: "dcmac10", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.221", Name: "dcmac11", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.108", Name: "dcmac13", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.189", Name: "dcmac15", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.147", Name: "dcmac17", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.177", Name: "dcmac19", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.128", Name: "dcmac21", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.28", Name: "dcmac23", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "192.168.0.29", Name: "dcmac25", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.3.87", Name: "dcmac27", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.1.58", Name: "dcmac12", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.3.204", Name: "dcmac29", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.1.114", Name: "dcmac14", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.1.21", Name: "dcmac31", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.0.51", Name: "dcmac16", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.2.152", Name: "dcmac33", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.2.247", Name: "dcmac18", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.3.217", Name: "dcmac35", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
		{IP: "10.57.2.139", Name: "dcmac37", Type: "device_worker", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}},
	}

	for _, entry := range data {
		insertQuery := "INSERT INTO servers (ip, cloud_name, cloud_type, created_at, updated_at, deleted_at) VALUES (?, ?, ?, NOW(), NOW(), NULL)"
		_, err := db.Exec(insertQuery, entry.IP, entry.Name, entry.Type)
		if err != nil {
			return err
		}

	}

	log.Println("Data inserted into the 'clouds' table successfully")

	selectQuery := "SELECT * FROM servers"
	rows, err := db.Query(selectQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	// ...
	var result domain.Cloud
	for rows.Next() {
		err := rows.Scan(&result.IP, &result.Name, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
		if err != nil {
			return err
		}
		log.Printf("Selected data: IP=%s, Name=%s, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s\n", result.IP, result.Name, result.CreatedAt, result.UpdatedAt, result.DeletedAt)
	}
	// ...

	return nil
}

func Migrate(db *sql.DB, zapLogger *zap.Logger) error {

	query := `SELECT ip FROM servers WHERE cloud_name='dcmac1'`

	_, err := db.Query(query)
	if err == nil {
		return domain.ErrExistsTable
	}

	for _, sql := range sqls {

		sqlFile, err := os.Open(sql)
		if err != nil {
			break
		}
		defer sqlFile.Close()

		sqlBytes, err := ioutil.ReadAll(sqlFile)
		if err != nil {
			log.Fatal(err)
		}

		sqlQuery := string(sqlBytes)

		_, err = db.Exec(sqlQuery)
		if err != nil {
			break
		}
	}

	zapLogger.Info("migrated")

	return nil
}
