package message

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fsufitch/r9kd/db"
	"github.com/fsufitch/r9kd/model"
	"github.com/fsufitch/r9kd/server/util"
	"github.com/gorilla/mux"
)

const channelNotFound = "Channel not found"
const banDurationMultiplier = 2

type postMessageRequestJSON struct {
	Body           string `json:"body"`
	SenderStringID string `json:"sender"`
}

type postMessageResponseJSON struct {
	MessageID      int    `json:"message_id"`
	SenderStringID string `json:"sender"`
	Banned         bool   `json:"banned"`
	BanExpireTime  int64  `json:"ban_expire"`
	BanDuration    int    `json:"ban_duration"`
}

func (r postMessageResponseJSON) Serialize() ([]byte, error) {
	return json.Marshal(r)
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	dryRun := (r.URL.Query().Get("dryRun") != "")
	channelStringID, ok := mux.Vars(r)["channel_id"]

	if !ok {
		util.WriteHTTPErrorResponse(w, 404, channelNotFound)
		return
	}

	channel, err := db.GetChannelByStringID(channelStringID)
	if err == sql.ErrNoRows {
		util.WriteHTTPErrorResponse(w, 404, channelNotFound)
		return
	} else if err != nil {
		util.WriteHTTPErrorResponse(w, 500, err.Error())
		return
	}

	// TODO: Add channel access control

	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.WriteHTTPErrorResponse(w, 400, "Failed reading request body: "+err.Error())
		return
	}

	var parsedJSON postMessageRequestJSON
	json.Unmarshal(bodyData, &parsedJSON)

	if err != nil {
		util.WriteHTTPErrorResponse(w, 400, "Failed parsing JSON: "+err.Error())
		return
	}

	newMessage := model.Message{
		Body:            parsedJSON.Body,
		Timestamp:       time.Now(),
		ChannelStringID: channel.StringID,
		SenderStringID:  parsedJSON.SenderStringID,
	}
	newMessage.CalculateHash()

	sender, err := safeGetSender(newMessage.SenderStringID, newMessage.ChannelStringID, dryRun)
	if err != nil {
		util.WriteHTTPErrorResponse(w, 500, "Error sourcing message sender: "+err.Error())
		return
	}

	bs, err := detectBanState(newMessage, sender)
	if err != nil {
		util.WriteHTTPErrorResponse(w, 500, "Error detecting ban state: "+err.Error())
		return
	}

	if bs.AlreadyBanned {
		util.WriteHTTPErrorResponse(w, 409, "Sender is already banned and ban has not expired")
		return
	}

	if !dryRun {
		var newMessageID int
		newMessageID, err = db.AddMessage(newMessage)
		if err != nil {
			util.WriteHTTPErrorResponse(w, 500, "Failed adding row: "+err.Error())
			return
		}
		newMessage, err = db.GetMessage(newMessageID)
		if err != nil {
			util.WriteHTTPErrorResponse(w, 500, "Failed re-fetching row: "+err.Error())
			return
		}
	}

	response := postMessageResponseJSON{
		MessageID:      newMessage.ID,
		SenderStringID: sender.StringID,
	}

	if bs.NewBan {
		banExpireTime := time.Now().Add(time.Duration(bs.NewBanLength) * time.Second)
		response.Banned = true
		response.BanExpireTime = banExpireTime.Unix()
		response.BanDuration = bs.NewBanLength

		if !dryRun {
			sender.Banned = true
			sender.BanExpiration = banExpireTime
			sender.LastBanLength = bs.NewBanLength

			err = db.SaveSender(sender)
			if err != nil {
				util.WriteHTTPErrorResponse(w, 500, "Failed saving banned sender: "+err.Error())
				return
			}
		}
	}

	util.WriteSerializableJSON(w, 200, response)
}

func safeGetSender(senderStringID string, channelStringID string, dryRun bool) (sender model.Sender, err error) {
	if dryRun {
		sender, err = db.GetSender(senderStringID, channelStringID)
		if err == sql.ErrNoRows {
			sender = model.Sender{
				StringID:        senderStringID,
				Banned:          false,
				BanExpiration:   time.Unix(0, 0),
				LastBanLength:   0,
				ChannelStringID: channelStringID,
			}
			err = nil
		}
	} else {
		sender, err = db.GetOrCreateSender(senderStringID, channelStringID)
	}
	return
}

type banState struct {
	AlreadyBanned bool
	NewBan        bool
	NewBanLength  int // seconds

}

func detectBanState(message model.Message, sender model.Sender) (banState, error) {
	sender.UpdateBanned()
	if sender.Banned {
		return banState{AlreadyBanned: true}, nil
	}

	dupes, err := db.CountMessagesWithSameHash(message.Hash, message.ChannelStringID)
	if err != nil {
		return banState{}, err
	}

	if dupes == 0 {
		return banState{}, nil
	}

	banLength := 1
	if sender.LastBanLength > 0 {
		banLength = sender.LastBanLength * banDurationMultiplier
	}
	return banState{
		NewBan:       true,
		NewBanLength: banLength,
	}, nil
}
