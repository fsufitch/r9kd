package channels

import (
	"database/sql"
	"net/http"

	"github.com/fsufitch/r9kd/db"
	"github.com/gorilla/mux"
)
import "github.com/fsufitch/r9kd/server/util"

const channelNotFound = "Channel not found"

func getChannel(w http.ResponseWriter, r *http.Request) {
	stringID, ok := mux.Vars(r)["string_id"]
	if !ok {
		util.WriteHTTPErrorResponse(w, 404, channelNotFound)
		return
	}

	channel, err := db.GetChannelByStringID(stringID)
	if err == sql.ErrNoRows {
		util.WriteHTTPErrorResponse(w, 404, channelNotFound)
		return
	} else if err != nil {
		util.WriteHTTPErrorResponse(w, 500, err.Error())
	}

	util.WriteSerializableJSON(w, 200, channel)
}
