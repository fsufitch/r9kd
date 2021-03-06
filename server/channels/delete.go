package channels

import (
	"database/sql"
	"net/http"

	"github.com/fsufitch/r9kd/auth"
	"github.com/fsufitch/r9kd/db"
	"github.com/fsufitch/r9kd/server/util"
	"github.com/gorilla/mux"
)

func deleteChannel(w http.ResponseWriter, r *http.Request) {
	key := auth.GetAPIKeyFromAuthorizationHeader(r)
	if key == "" {
		auth.Write401Response(w)
		return
	}
	if !auth.RequireAdminPermissions(key) {
		auth.Write403Response(w, "delete channel")
		return
	}

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
		return
	}

	err = db.DeleteChannel(channel.ID)

	if err != nil {
		util.WriteHTTPErrorResponse(w, 500, err.Error())
		return
	}

	util.WriteSerializableJSON(w, 200, util.HTTPBasicSuccessResponse)
}
