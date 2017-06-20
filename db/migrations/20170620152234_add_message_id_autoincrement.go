package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170620152234, Down20170620152234)
}

// Up20170620152234 updates the database to the new requirements
func Up20170620152234(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE SEQUENCE message_id_seq;
		ALTER TABLE messages ALTER id SET DEFAULT NEXTVAL('message_id_seq');
	`)
	return err
}

// Down20170620152234 should send the database back to the state it was from before Up was ran
func Down20170620152234(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE messages ALTER id DROP DEFAULT;
		DROP SEQUENCE message_id_seq;
	`)
	return err
}
