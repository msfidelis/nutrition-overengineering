package migration

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"database-store-service/pkg/db"
	"database-store-service/pkg/envs"
)

func Migrate() {
	fmt.Println("Migrating")

	conn, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	driver, err := GetDriver(conn)
	if err != nil {
		panic(err)
	}

	path := envs.Getenv("MIGRATIONS_PATH", "./migrations")
	migrationsPath := fmt.Sprintf("file://%s", path)
	fmt.Println(migrationsPath)

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		panic(err)
	}
	m.Steps(1000)
}

func GetDriver(conn *sql.DB) (database.Driver, error) {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return driver, nil

}