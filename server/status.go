package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fsufitch/r9kd/db"
)

type statusResponse struct {
	Ok      bool  `json:"ok"`
	Uptime  int64 `json:"uptimeNano"`
	DbState struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	} `json:"dbState"`
}

type statusHandler struct {
	startTimeNano int64
}

func nowNano() int64 {
	return time.Now().UTC().UnixNano()
}

func newStatusHandler() statusHandler {
	return statusHandler{nowNano()}
}

func (h statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := statusResponse{
		Uptime: nowNano() - h.startTimeNano,
	}

	dbError := db.GetCachedConnection().Error
	if dbError != nil {
		response.DbState.Ok = false
		response.DbState.Error = dbError.Error()
	} else {
		response.DbState.Ok = true
	}

	response.Ok = response.DbState.Ok && true

	data, _ := json.MarshalIndent(response, "", "  ")

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
