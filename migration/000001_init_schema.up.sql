-- This SQL script creates the "clouds" table
-- with the specified columns.

-- Define the "clouds" table

-- INSERT statements
INSERT INTO servers (ip, cloud_name, cloud_type, cloud_status, cloud_state, created_at, updated_at, deleted_at)
VALUES
    ('192.168.0.106', 'dcdell3', 'handler', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('192.168.0.12', 'dcdell4', 'handler',  'inactive', 'offline', NOW(), NOW(), NULL),
    ('192.168.0.129', 'dcdell5', 'handler', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('192.168.0.207', 'dcmac1', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.157', 'dcmac2', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.144', 'dcmac3', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.55', 'dcmac4', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.138', 'dcmac5', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.62', 'dcmac6', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.100', 'dcmac7', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.162', 'dcmac8', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.112', 'dcmac9', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.195', 'dcmac10', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.221', 'dcmac11', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.108', 'dcmac13', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.189', 'dcmac15', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.147', 'dcmac17', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.177', 'dcmac19', 'device_worker','inactive', 'online',  NOW(), NOW(), NULL),
    ('192.168.0.128', 'dcmac21', 'device_worker','inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.28', 'dcmac23', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('192.168.0.29', 'dcmac25', 'device_worker', 'inactive', 'online', NOW(), NOW(), NULL),
    ('10.57.3.87', 'dcmac27', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('10.57.1.58', 'dcmac12', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('10.57.3.204', 'dcmac29', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('10.57.1.114', 'dcmac14', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('10.57.1.21', 'dcmac31', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('10.57.0.51', 'dcmac16', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('10.57.2.152', 'dcmac33', 'device_worker', 'inactive', 'offline',  NOW(), NOW(), NULL),
    ('10.57.2.247', 'dcmac18', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('10.57.3.217', 'dcmac35', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL),
    ('10.57.2.139', 'dcmac37', 'device_worker', 'inactive', 'offline', NOW(), NOW(), NULL);

    