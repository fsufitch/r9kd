package model

import "encoding/json"

// MessageChannel is a one-to-many for messages
type MessageChannel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	StringID string `json:"string_id"`
}

// SerializedMessageChannel holds a channel's JSON serialized data
type SerializedMessageChannel []byte

// Deserialize turns a serialized message channel back into a real object
func (smc SerializedMessageChannel) Deserialize(target interface{}) error {
	if _, ok := target.(*MessageChannel); !ok {
		panic("Type assertion of target to MessageChannel failed")
	}
	mc, _ := target.(*MessageChannel)
	err := json.Unmarshal(smc, &mc)
	return err
}

// Serialize turns a message channel into its serialized form
func (mc MessageChannel) Serialize() ([]byte, error) {
	return json.Marshal(mc)
}
