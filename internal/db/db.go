package db

import (
	"context"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB holds the database connection pool
type DB struct {
	SQL *pgxpool.Pool
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates database pool for Postgres
func ConnectSQL() (*DB, error) {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		panic("DATABASE_URL environment variable is required")
	}

	d, err := NewDatabase(DATABASE_URL)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// testDB tries to ping the database
func testDB(d *pgxpool.Pool) error {
	err := d.Ping(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}

// GetDB returns a connection to the database
func GetDB() *pgxpool.Pool {
	return dbConn.SQL
}
