package sender

import (
	"database/sql"
	"net/http"

	"github.com/fsufitch/r9kd/db"
	"github.com/fsufitch/r9kd/server/util"
	"github.com/gorilla/mux"
)

const senderNotFound = "Sender not found"

func getSender(w http.ResponseWriter, r *http.Request) {
	channelStringID, ok1 := mux.Vars(r)["channel_id"]
	stringID, ok2 := mux.Vars(r)["sender_id"]
	if !(ok1 && ok2) {
		util.WriteHTTPErrorResponse(w, 404, senderNotFound)
		return
	}

	sender, err := db.GetSender(stringID, channelStringID)

	if err == sql.ErrNoRows {
		util.WriteHTTPErrorResponse(w, 404, senderNotFound)
		return
	} else if err != nil {
		util.WriteHTTPErrorResponse(w, 500, err.Error())
	}

	util.WriteSerializableJSON(w, 200, sender)
}
