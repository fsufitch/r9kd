package model

import (
	"encoding/json"
	"hash/adler32"
	"strings"
	"time"
)

// Message represents one message logged by r9kd
type Message struct {
	ID              int
	Body            string
	Timestamp       time.Time
	Hash            uint32
	ChannelStringID string
	SenderStringID  string
}

type messageJSON struct {
	ID              int    `json:"id"`
	Body            string `json:"body"`
	Timestamp       int64  `json:"timestamp"`
	Hash            uint32 `json:"hash"`
	ChannelStringID string `json:"channel"`
	SenderStringID  string `json:"sender"`
}

// Serialize fulfills the Serializable interface
func (m Message) Serialize() ([]byte, error) {
	return json.Marshal(messageJSON{
		ID:              m.ID,
		Body:            m.Body,
		Timestamp:       m.Timestamp.Unix(),
		Hash:            m.Hash,
		ChannelStringID: m.ChannelStringID,
		SenderStringID:  m.SenderStringID,
	})
}

// CalculateHash populates the message's hash for easy comparison/lookup in DB
func (m *Message) CalculateHash() {
	data := m.Body
	data = strings.ToLower(data)
	data = strings.TrimSpace(data)
	m.Hash = adler32.Checksum([]byte(data))
}
