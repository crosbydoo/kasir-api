package database

import (
	"database/sql"
	"kasir-api/internal/pkg"

	// _ "github.com/lib/pq" // postgres driver
	_ "github.com/jackc/pgx/v5/stdlib" // postgres driver
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func InitDB(connectionString string) (*sql.DB, error) {
	// Open connection
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	pkg.Log.Info("Database connection opened")
	return db, nil
}
