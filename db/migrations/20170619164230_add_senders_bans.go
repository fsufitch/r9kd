package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170619164230, Down20170619164230)
}

// Up20170619164230 updates the database to the new requirements
func Up20170619164230(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE senders (
			id        			INT PRIMARY KEY,
			string_id 			VARCHAR(64) UNIQUE,
			banned					BOOLEAN,
			ban_expire_time TIMESTAMP,
			last_ban_length INT,
			channel 				INT REFERENCES channels (id) ON DELETE CASCADE
			UNIQUE 					(string_id, channel)
		);

		CREATE SEQUENCE senders_id_seq;
		ALTER TABLE senders ALTER id SET DEFAULT NEXTVAL('senders_id_seq');

		ALTER TABLE messages ADD COLUMN channel INT REFERENCES channels (id) ON DELETE CASCADE;
		ALTER TABLE messages ADD COLUMN sender INT REFERENCES senders (id) ON DELETE CASCADE;
	`)
	return err
}

// Down20170619164230 should send the database back to the state it was from before Up was ran
func Down20170619164230(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE messages DROP COLUMN sender;
		ALTER TABLE messages DROP COLUMN channel;

		DROP TABLE senders;
		DROP SEQUENCE senders_id_seq;
	`)
	return err
}
