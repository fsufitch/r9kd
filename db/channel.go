package db

import "github.com/fsufitch/r9kd/model"

// AddChannel adds a new channel to the database
func AddChannel(name, stringID string) error {
	if GetCachedConnection().Error != nil {
		return GetCachedConnection().Error
	}
	conn := GetCachedConnection().Conn
	_, err := conn.Exec(`INSERT INTO channels (name, string_id) VALUES ($1, $2);`, name, stringID)
	return err
}

// GetChannelByStringID does what it says on the can
func GetChannelByStringID(stringID string) (channel model.MessageChannel, err error) {
	if GetCachedConnection().Error != nil {
		err = GetCachedConnection().Error
		return
	}
	conn := GetCachedConnection().Conn
	row := conn.QueryRow(`
    SELECT id, name, string_id
    FROM channels
    WHERE string_id=$1`, stringID)

	err = row.Scan(&channel.ID, &channel.Name, &channel.StringID)
	return
}

// DeleteChannel does what it says on the can
func DeleteChannel(id int) (err error) {
	if GetCachedConnection().Error != nil {
		err = GetCachedConnection().Error
		return
	}
	conn := GetCachedConnection().Conn
	_, err = conn.Exec(`
    DELETE FROM channels
    WHERE id=$1`, id)
	return
}
