package db

import "github.com/fsufitch/r9kd/model"

// AddMessage does what it says on the can
func AddMessage(message model.Message) (newMessageID int, err error) {
	if GetCachedConnection().Error != nil {
		err = GetCachedConnection().Error
		return
	}
	conn := GetCachedConnection().Conn
	row := conn.QueryRow(`
    INSERT INTO messages (body, timestamp, hash, channel, sender)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id;
    `,
		message.Body,
		message.Timestamp,
		message.Hash,
		message.ChannelStringID,
		message.SenderStringID,
	)

	err = row.Scan(&newMessageID)
	return
}

// GetMessage does what it says on the can
func GetMessage(id int) (message model.Message, err error) {
	if GetCachedConnection().Error != nil {
		err = GetCachedConnection().Error
		return
	}
	conn := GetCachedConnection().Conn
	row := conn.QueryRow(`
    SELECT id, body, timestamp, hash, channel, sender
    FROM messages
    WHERE id=$1`, id,
	)

	message = model.Message{}
	err = row.Scan(
		&message.ID,
		&message.Body,
		&message.Timestamp,
		&message.Hash,
		&message.ChannelStringID,
		&message.SenderStringID,
	)

	return
}

// CountMessagesWithSameHash counts how many messages with the same hash as was
// given exist within a particular channel
func CountMessagesWithSameHash(hash uint32, channelStringID string) (dupes int, err error) {
	if GetCachedConnection().Error != nil {
		err = GetCachedConnection().Error
		return
	}
	conn := GetCachedConnection().Conn
	row := conn.QueryRow(`
    SELECT COUNT(*)
    FROM messages
    WHERE hash=$1 AND channel=$2`,
		hash, channelStringID,
	)

	err = row.Scan(&dupes)

	return
}
