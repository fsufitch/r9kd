package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testPastTime = time.Date(1970, time.January, 1, 0, 1, 0, 0, time.UTC)
var testFutureTime = time.Date(2070, time.January, 1, 0, 1, 0, 0, time.UTC)

var testSender = Sender{
	ID:              98765,
	StringID:        "test_sender",
	Banned:          false,
	BanExpiration:   testPastTime,
	LastBanLength:   0,
	ChannelStringID: "abcd",
}

var testSenderSerialized = `{
  "id": 98765,
  "string_id": "test_sender",
  "banned": false,
  "ban_expire_time": 60,
  "last_ban_length": 0,
  "channel": "abcd"
}`

func TestSenderSerialize(t *testing.T) {
	data, err := testSender.Serialize()
	assert.Nil(t, err)
	assert.JSONEq(t, testSenderSerialized, string(data))
}

func TestSenderUpdateBanned_BanExpired(t *testing.T) {
	senderCopy := testSender
	senderCopy.Banned = true
	senderCopy.BanExpiration = testPastTime

	senderCopy.UpdateBanned()
	assert.False(t, senderCopy.Banned)

	// Doing it again makes no difference
	senderCopy.UpdateBanned()
	assert.False(t, senderCopy.Banned)
}

func TestSenderUpdateBanned_BanUpheld(t *testing.T) {
	senderCopy := testSender
	senderCopy.Banned = false
	senderCopy.BanExpiration = testFutureTime

	senderCopy.UpdateBanned()
	assert.True(t, senderCopy.Banned)

	// Doing it again makes no difference
	senderCopy.UpdateBanned()
	assert.True(t, senderCopy.Banned)
}
