package auth

import (
	"database/sql"
	"fmt"

	"github.com/fsufitch/r9kd/db"
)

// APIKey represents one key entry in the database
type APIKey struct {
	id                int
	Key               string
	Admin             bool
	ChannelModify     *int // using pointers since these are nullable
	ChannelAddMessage *int
}

func getAPIKeysFromRows(rows *sql.Rows) (keys []APIKey, err error) {
	keys = []APIKey{}
	for rows.Next() {
		key := APIKey{}
		if err = rows.Scan(
			&key.id,
			&key.Key,
			&key.Admin,
			&key.ChannelModify,
			&key.ChannelAddMessage,
		); err != nil {
			keys = []APIKey{}
			return
		}
		keys = append(keys, key)
	}

	if err = rows.Err(); err != nil {
		keys = []APIKey{}
		return
	}
	return
}

// GetAPIKeyByID does what it says on the can
func GetAPIKeyByID(id string) (key APIKey, err error) {
	conn := db.GetCachedConnection().Conn
	rows, err := conn.Query(`
    SELECT id, key, admin, channel_modify, channel_add_message
    FROM api_keys
    WHERE id=$1
  `, id)
	if err != nil {
		return
	}
	defer rows.Close()

	keys, err := getAPIKeysFromRows(rows)
	if err != nil {
		return
	}
	if len(keys) == 0 {
		err = sql.ErrNoRows
		return
	}

	key = keys[0]
	return
}

// GetAPIKeyByKey does what it says on the can
func GetAPIKeyByKey(keyStr string) (key APIKey, err error) {
	conn := db.GetCachedConnection().Conn
	rows, err := conn.Query(`
    SELECT id, key, admin, channel_modify, channel_add_message
    FROM api_keys
    WHERE key=$1;
  `, keyStr)
	if err != nil {
		return
	}
	defer rows.Close()

	keys, err := getAPIKeysFromRows(rows)
	if err != nil {
		return
	}
	if len(keys) == 0 {
		err = sql.ErrNoRows
		return
	}

	key = keys[0]
	return
}

// AddAPIKey does what it says on the can
func AddAPIKey(key APIKey) (err error) {
	conn := db.GetCachedConnection().Conn
	_, err = conn.Exec(`
    INSERT INTO api_keys (key, admin, channel_modify, channel_add_message)
    VALUES ($1, $2, $3, $4);
  `, key.Key, key.Admin, key.ChannelModify, key.ChannelAddMessage)
	return
}

// UpdateAPIKey does what it says on the can
func UpdateAPIKey(key APIKey) (err error) {
	conn := db.GetCachedConnection().Conn
	_, err = conn.Exec(`
    UPDATE api_keys
    SET (key, admin, channel_modify, channel_add_message) = ($2, $3, $4, $5)
    WHERE id=$1;
  `, key.id, key.Key, key.Admin, key.ChannelModify, key.ChannelAddMessage)
	return
}

// RequireAdminPermissions returns whether the given key has global admin rights
func RequireAdminPermissions(key string) bool {
	apiKey, err := GetAPIKeyByKey(key)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		fmt.Printf("Error checking permissions: %v\n", err)
		return false
	}
	return apiKey.Admin
}
