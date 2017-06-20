package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170620150122, Down20170620150122)
}

// Up20170620150122 updates the database to the new requirements
func Up20170620150122(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE api_keys ALTER key SET NOT NULL;
		ALTER TABLE api_keys ALTER admin SET DEFAULT false;

		ALTER TABLE channels ALTER name SET NOT NULL;
		ALTER TABLE channels ALTER string_id SET NOT NULL;

		ALTER TABLE messages ALTER body SET NOT NULL;
		ALTER TABLE messages ALTER "timestamp" SET NOT NULL;
		ALTER TABLE messages ALTER hash SET DEFAULT 0;

		ALTER TABLE senders ALTER string_id SET NOT NULL;
		ALTER TABLE senders ALTER banned SET DEFAULT false;
		ALTER TABLE senders ALTER ban_expire_time SET DEFAULT to_timestamp(0);
		ALTER TABLE senders ALTER last_ban_length SET DEFAULT 0;
	`)
	return err
}

// Down20170620150122 should send the database back to the state it was from before Up was ran
func Down20170620150122(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE api_keys ALTER key DROP NOT NULL;
		ALTER TABLE api_keys ALTER admin DROP DEFAULT;

		ALTER TABLE channels ALTER name DROP NOT NULL;
		ALTER TABLE channels ALTER string_id DROP NOT NULL;

		ALTER TABLE messages ALTER body DROP NOT NULL;
		ALTER TABLE messages ALTER "timestamp" DROP NOT NULL;
		ALTER TABLE messages ALTER hash DROP DEFAULT;

		ALTER TABLE senders ALTER string_id DROP NOT NULL;
		ALTER TABLE senders ALTER banned DROP DEFAULT;
		ALTER TABLE senders ALTER last_ban_length DROP DEFAULT;
	`)
	return err
}
