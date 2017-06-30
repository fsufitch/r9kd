package model

import (
	"testing"
	"time"
)

import "github.com/stretchr/testify/assert"

var testTime1 = time.Date(1970, time.January, 1, 0, 0, 1, 0, time.UTC)
var testTime2 = time.Date(1971, time.January, 1, 0, 0, 1, 0, time.UTC)

var testMessage = Message{
	ID:              54321,
	Body:            "This is a test message",
	Timestamp:       testTime1,
	Hash:            0,
	ChannelStringID: "abc",
	SenderStringID:  "def",
}

var testMessageSerialized = `{
  "id": 54321,
  "body": "This is a test message",
  "timestamp": 1,
  "hash": 0,
  "channel": "abc",
  "sender": "def"
}`

func TestMessageSerialize(t *testing.T) {
	data, err := testMessage.Serialize()
	assert.Nil(t, err)
	assert.JSONEq(t, testMessageSerialized, string(data))
}

func TestMessageHash_DependsOnBody(t *testing.T) {
	origMsg := testMessage
	origMsg.CalculateHash()

	copyMsg := origMsg
	copyMsg.Body = "Different body"
	copyMsg.CalculateHash()
	assert.NotEqual(t, origMsg.Hash, copyMsg.Hash)
}

func TestMessageHash_DoesNotDependOnOtherProps(t *testing.T) {
	origMsg := testMessage
	origMsg.CalculateHash()

	copyMsg := origMsg
	copyMsg.ID = 123456
	copyMsg.Timestamp = testTime2
	copyMsg.ChannelStringID = "xyz"
	copyMsg.SenderStringID = "zyx"
	copyMsg.CalculateHash()
	assert.Equal(t, origMsg.Hash, copyMsg.Hash)
}
