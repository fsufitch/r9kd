package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170614152417, Down20170614152417)
}

// Up20170614152417 updates the database to the new requirements
func Up20170614152417(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE SEQUENCE api_keys_id_seq;
		ALTER TABLE api_keys ALTER id SET DEFAULT NEXTVAL('api_keys_id_seq');
	`)
	return err
}

// Down20170614152417 should send the database back to the state it was from before Up was ran
func Down20170614152417(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE api_keys ALTER id DROP DEFAULT;
		DROP SEQUENCE api_keys_id_seq;
	`)
	return err
}
