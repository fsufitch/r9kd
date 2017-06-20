package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170614142902, Down20170614142902)
}

// Up20170614142902 updates the database to the new requirements
func Up20170614142902(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE channels (
			id        INT PRIMARY KEY,
			name      VARCHAR(128),
			string_id VARCHAR(64) UNIQUE
		);

		ALTER TABLE messages
			ADD channel	VARCHAR(64) REFERENCES channels (string_id) ON DELETE CASCADE;
	`)
	return err
}

// Down20170614142902 should send the database back to the state it was from before Up was ran
func Down20170614142902(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE messages DROP channel_id;
		DROP TABLE channels;
	`)
	return err

}
