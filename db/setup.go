package db

import (
	"database/sql"

	_ "github.com/lib/pq" // Install postgres driver
)

// migrationStep describes a step in a database migration/setup
type migrationStep interface {
	StepNumber() int
	Check() bool
	Apply()
}

func foo() {
	_, _ = sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
}
