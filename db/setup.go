package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/fsufitch/r9kd/db/migrations" // Import migrations for goose to find
	_ "github.com/lib/pq"                      // Postgres driver
	"github.com/pressly/goose"
)

// ConnectionAttempt encapsulates an attempt to connect to Postgres, and either
// the connection or error that reulted
type ConnectionAttempt struct {
	Conn      *sql.DB
	Error     error
	attempted bool
	dburl     string
}

var cachedConnection ConnectionAttempt

// GetCachedConnection returns the last connection attempt
func GetCachedConnection() ConnectionAttempt {
	return cachedConnection
}

// Connect connects to the given database URL and recycles the connection
func Connect(dburl string, force bool) ConnectionAttempt {
	if cachedConnection.dburl == dburl && cachedConnection.attempted && !force {
		return cachedConnection
	}

	if cachedConnection.Conn != nil {
		cachedConnection.Conn.Close()
	}

	newConn := ConnectionAttempt{}

	if dburl == "" {
		dburl = os.Getenv("DATABASE_URL")
	}

	if dburl == "" {
		newConn.Error = errors.New("No database URL received")
		fmt.Println(newConn.Error)
	} else {
		newConn.Conn, newConn.Error = sql.Open("postgres", dburl)
		newConn.attempted = true

		if newConn.Error != nil {
		} else if newConn.Conn == nil {
			newConn.Error = errors.New("No connection from sql.Open() but no error either")
		} else {
			newConn.Error = newConn.Conn.Ping()
		}
	}

	if newConn.Error == nil {
		fmt.Println("Database connection successful")
	} else {
		fmt.Println(newConn.Error)
	}

	cachedConnection = newConn
	return newConn
}

// Migrate attempts to run database migrations on the given DB connection
func Migrate(db *sql.DB) (err error) {
	fmt.Println("Running database migrations...")
	err = goose.SetDialect("postgres")
	if err != nil {
		return
	}
	err = goose.Run("up", db, ".")
	return
}
