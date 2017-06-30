package model

import "testing"
import "github.com/stretchr/testify/assert"

var testChannel = MessageChannel{
	ID:       12345,
	Name:     "Test Channel Name",
	StringID: "test_channel",
}

var testChannelSerialized = SerializedMessageChannel(`{
  "id": 12345,
  "name": "Test Channel Name",
  "string_id": "test_channel"
}`)

func TestChannelSerialize(t *testing.T) {
	data, err := testChannel.Serialize()
	assert.Nil(t, err)
	assert.JSONEq(t, string(testChannelSerialized), string(data))
}

func TestChannelDeserialize(t *testing.T) {
	tc := MessageChannel{}
	err := testChannelSerialized.Deserialize(&tc)
	assert.Nil(t, err)
	assert.Equal(t, testChannel, tc)
}

func TestChannelDeserialize_InvalidTarget(t *testing.T) {
	tc := struct{}{}
	assert.Panics(t, func() {
		testChannelSerialized.Deserialize(&tc)
	})
}
