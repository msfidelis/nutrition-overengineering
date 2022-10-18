package db

import (
	"database/sql"
	"fmt"
	"database-store-service/pkg/envs"
	"database-store-service/pkg/logger"

	_ "github.com/lib/pq"
)

var connection *sql.DB

func GetConnection() (*sql.DB, error) {
	log := logger.Instance()

	if connection == nil {
		log.Info().
			Str("component", "database").
			Msg("Creating a first connection with postgres database")

		host := envs.Getenv("DB_HOST", "0.0.0.0")
		port := envs.Getenv("DB_PORT", "5432")
		user := envs.Getenv("DB_USER", "user")
		password := envs.Getenv("DB_PASSWORD", "pass")
		database := envs.Getenv("DB_SCHEMA", "open-prr")
		sslMode := envs.Getenv("POSTGRES_SSL_MODE", "disable")
		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, database, sslMode)
		fmt.Println(connectionString)
		conn, err := sql.Open("postgres", connectionString)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		connection = conn
	}

	log.Debug().
		Str("component", "database").
		Msg("Retrieving already existing database connection")

	return connection, nil
}