package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170614143948, Down20170614143948)
}

// Up20170614143948 updates the database to the new requirements
func Up20170614143948(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE SEQUENCE channel_id_seq;
		ALTER TABLE channels ALTER id SET DEFAULT NEXTVAL('channel_id_seq');

		INSERT INTO channels (name, string_id)
		  VALUES ('Default Channel', 'default');
	`)
	return err
}

// Down20170614143948 should send the database back to the state it was from before Up was ran
func Down20170614143948(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DELETE FROM channels where string_id='default';
		ALTER TABLE channels ALTER id DROP DEFAULT;
		DROP SEQUENCE channel_id_seq;
	`)
	return err
}
