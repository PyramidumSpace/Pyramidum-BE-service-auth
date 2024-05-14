package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var postgresUser, postgresPass, postgresHost, postgresPort, dbName string
	var migrationsPath, migrationsTable string
	var forceVersion int
	var dropBase bool

	flag.StringVar(&postgresUser, "user", "service-auth-owner", "postgres user")
	flag.StringVar(&postgresPass, "password", "service-auth-owner", "postgres password")
	flag.StringVar(&postgresHost, "host", "localhost", "postgres host")
	flag.StringVar(&postgresPort, "port", "5434", "postgres port")
	flag.StringVar(&dbName, "db", "service-auth", "postgres database")
	flag.IntVar(&forceVersion, "force", 0, "postgres database")
	flag.BoolVar(&dropBase, "drop", false, "postgres database")

	flag.StringVar(&migrationsPath, "migrations-path", "./migrations", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?x-migrations-table=%s&sslmode=disable",
			postgresUser,
			postgresPass,
			postgresHost,
			postgresPort,
			dbName,
			migrationsTable,
		),
	)
	if err != nil {
		panic(err)
	}

	if forceVersion != 0 {
		if err := m.Force(forceVersion); err != nil {
			panic(err)
		}
	}

	if dropBase {
		if err := m.Drop(); err != nil {
			panic(err)
		}
	}

	// run migrations
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(err)
	}
}
