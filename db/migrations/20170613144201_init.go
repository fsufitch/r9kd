package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170613144201, Down20170613144201)
}

// Up20170613144201 updates the database to the new requirements
func Up20170613144201(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE messages (
			id        INT PRIMARY KEY,
			body      TEXT,
			timestamp TIMESTAMP,
			hash      BIGINT
		);

		CREATE INDEX messages_hash_idx
		  ON messages (hash);
	`)
	return err
}

// Down20170613144201 should send the database back to the state it was from before Up was ran
func Down20170613144201(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE messages;
		DROP INDEX IF EXISTS messages_hash_idx;
	`)
	return err
}
