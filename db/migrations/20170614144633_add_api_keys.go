package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170614144633, Down20170614144633)
}

// Up20170614144633 updates the database to the new requirements
func Up20170614144633(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE api_keys (
			id        INT PRIMARY KEY,
			key      	VARCHAR(128) UNIQUE,
			admin									BOOLEAN,
			channel_modify 				INT REFERENCES channels (id) ON DELETE CASCADE,
			channel_add_message 	INT REFERENCES channels (id) ON DELETE CASCADE
		);
	`)
	return err
}

// Down20170614144633 should send the database back to the state it was from before Up was ran
func Down20170614144633(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE api_keys;
	`)
	return err
}
