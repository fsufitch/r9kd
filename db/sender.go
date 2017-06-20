package db

import "github.com/fsufitch/r9kd/model"

// AddSender adds a new sender to the database
func AddSender(stringID string, channelStringID string) (err error) {
	if GetCachedConnection().Error != nil {
		return GetCachedConnection().Error
	}
	conn := GetCachedConnection().Conn

	channel, err := GetChannelByStringID(channelStringID)
	if err != nil {
		return
	}

	_, err = conn.Exec(`INSERT INTO senders (string_id, channel) VALUES ($1, $2);`, stringID, channel.StringID)
	return err
}

// GetSender retrieves a sender from the database
func GetSender(senderStringID string, channelStringID string) (sender model.Sender, err error) {
	if GetCachedConnection().Error != nil {
		err = GetCachedConnection().Error
		return
	}
	conn := GetCachedConnection().Conn
	row := conn.QueryRow(`
    SELECT s.id, s.string_id, s.banned, s.ban_expire_time, s.last_ban_length, c.string_id
    FROM senders AS s
    INNER JOIN channels AS c ON s.channel=c.id
    WHERE s.string_id=$1 AND c.string_id=$2`,
		senderStringID, channelStringID,
	)

	err = row.Scan(
		&sender.ID,
		&sender.StringID,
		&sender.Banned,
		&sender.BanExpiration,
		&sender.LastBanLength,
		&sender.ChannelStringID,
	)
	return
}

// SaveSender updates the entry for a sender
// **Only updates ban status, not IDs or relationships!**
func SaveSender(sender model.Sender) (err error) {
	if GetCachedConnection().Error != nil {
		err = GetCachedConnection().Error
		return
	}
	conn := GetCachedConnection().Conn

	_, err = conn.Exec(`
    UPDATE senders
    SET (banned, ban_expire_time, last_ban_length) = ($1, $2, $3)
    WHERE id=$4;`,
		sender.Banned, sender.BanExpiration, sender.LastBanLength, sender.ID,
	)
	return
}
