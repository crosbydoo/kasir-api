package database

import (
	"database/sql"
	"kasir-api/internal/pkg"
	"net/url"
	"strings"

	// _ "github.com/lib/pq" // postgres driver
	_ "github.com/jackc/pgx/v5/stdlib" // postgres driver
	// Note: When using Supabase transaction pooler (port 6543), connect with default_query_exec_mode=simple_protocol
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func InitDB(connectionString string) (*sql.DB, error) {
	// Parse the connection string to check for Supabase pooler
	// and force simple protocol if needed
	if u, err := url.Parse(connectionString); err == nil {
		if strings.Contains(u.Host, ".pooler.supabase.com") {
			q := u.Query()
			q.Set("default_query_exec_mode", "simple_protocol")
			u.RawQuery = q.Encode()
			connectionString = u.String()
			pkg.Log.Info("Detected Supabase pooler, forcing default_query_exec_mode=simple_protocol")
		}
	}

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
