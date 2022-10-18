

CREATE TABLE IF NOT EXISTS reports (
	id VARCHAR ( 255 ) PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);


CREATE TABLE IF NOT EXISTS health_info (
	id VARCHAR ( 255 ) PRIMARY KEY,
    age INT NOT NULL,
    weight INT NOT NULL, 
    height INT NOT NULL,
    gender VARCHAR(1),
    activity_insentity VARCHAR(20),
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);


CREATE TABLE IF NOT EXISTS imc (
	id VARCHAR ( 255 ) PRIMARY KEY,
    result REAL NOT NULL,
    class VARCHAR ( 255 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS basal (
	id VARCHAR ( 255 ) PRIMARY KEY,
    bmr REAL NOT NULL,
    necessity REAL NOT NULL,
    unit VARCHAR ( 10 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS recomendations_protein (
	id VARCHAR ( 255 ) PRIMARY KEY,
    value REAL NOT NULL,
    unit VARCHAR ( 10 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS recomendations_water (
	id VARCHAR ( 255 ) PRIMARY KEY,
    value REAL NOT NULL,
    unit VARCHAR ( 10 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS recomendations_calories (
	id VARCHAR ( 255 ) PRIMARY KEY,
    maintain REAL NOT NULL,
    loss REAL NOT NULL,
    maintain REAL NOT NULL,
    fain VARCHAR ( 10 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL
);