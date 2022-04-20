CREATE TABLE IF NOT EXISTS service_parameters (
    id INT PRIMARY KEY AUTO_INCREMENT,
    variable VARCHAR(100) NOT NULL,
    value TEXT NOT NULL,
    description VARCHAR(255)NOT NULL,
    created_by varchar(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(100),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_by varchar(100),
    deleted_at_unix INT (10) DEFAULT 0,
    CONSTRAINT service_parameters_variable_deleted_at_unix_unique UNIQUE (variable, deleted_at_unix)
);