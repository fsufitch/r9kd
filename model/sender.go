package model

import (
	"encoding/json"
	"time"
)

// Sender encapsulates the idea of a unique sender to a particular channel
type Sender struct {
	ID              int
	StringID        string
	Banned          bool
	BanExpiration   time.Time
	LastBanLength   int
	ChannelStringID string
}

type senderJSON struct {
	ID              int    `json:"id"`
	StringID        string `json:"string_id"`
	Banned          bool   `json:"banned"`
	BanExpiration   int64  `json:"ban_expire_time"`
	LastBanLength   int    `json:"last_ban_length"`
	ChannelStringID string `json:"channel"`
}

// Serialize conforms to the Serializer interface to let Sender objects be serialized
func (s Sender) Serialize() ([]byte, error) {
	actual := senderJSON{
		ID:              s.ID,
		StringID:        s.StringID,
		Banned:          s.Banned,
		BanExpiration:   s.BanExpiration.Unix(),
		LastBanLength:   s.LastBanLength,
		ChannelStringID: s.ChannelStringID,
	}

	return json.Marshal(actual)
}

// UpdateBanned updates the data in sender.Banned
func (s *Sender) UpdateBanned() {
	s.Banned = s.BanExpiration.After(time.Now())
}
