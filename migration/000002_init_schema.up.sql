

CREATE TABLE IF NOT EXISTS servers (
        id INT AUTO_INCREMENT PRIMARY KEY,
        ip VARCHAR(255),
        cloud_name   VARCHAR(255),
        cloud_type   VARCHAR(255),
        cloud_status VARCHAR(255) NOT NULL,
        cloud_state  VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP DEFAULT NULL
);
    