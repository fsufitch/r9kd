package channels

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/fsufitch/r9kd/auth"
	"github.com/fsufitch/r9kd/db"
	"github.com/fsufitch/r9kd/model"
	"github.com/fsufitch/r9kd/server/util"
)

var validIDRegexp = regexp.MustCompile("^[a-z0-9._-]{5,}$")

func postChannel(w http.ResponseWriter, r *http.Request) {
	key := auth.GetAPIKeyFromAuthorizationHeader(r)
	if key == "" {
		auth.Write401Response(w)
		return
	}
	if !auth.RequireAdminPermissions(key) {
		auth.Write403Response(w, "create new channel")
		return
	}

	rawData, _ := ioutil.ReadAll(r.Body)
	var inputMC model.MessageChannel
	if err := model.SerializedMessageChannel(rawData).Deserialize(&inputMC); err != nil {
		util.WriteHTTPErrorResponse(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if inputMC.Name == "" {
		util.WriteHTTPErrorResponse(w, 400, "Property `name` missing")
		return
	}

	if inputMC.StringID == "" {
		util.WriteHTTPErrorResponse(w, 400, "Property `string_id` missing")
		return
	}

	if !validIDRegexp.MatchString(inputMC.StringID) {
		util.WriteHTTPErrorResponse(w, 400, "Property `string_id` has invalid value. 5+ characters that are a-z, 0-9, ., _, or - required")
		return

	}

	if _, err := db.GetChannelByStringID(inputMC.StringID); err != sql.ErrNoRows {
		msg := fmt.Sprintf("A channel by that string_id already exists (%v)", err)
		util.WriteHTTPErrorResponse(w, 400, msg)
		return
	}

	if err := db.AddChannel(inputMC.Name, inputMC.StringID); err != nil {
		util.WriteHTTPErrorResponse(w, 500, "Server error writing entry to DB: "+err.Error())
		return
	}

	channel, err := db.GetChannelByStringID(inputMC.StringID)
	if err != nil {
		util.WriteHTTPErrorResponse(w, 500, "Server error retrieving new entry: "+err.Error())
		return
	}

	util.WriteSerializableJSON(w, 200, channel)
}
