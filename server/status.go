package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type statusResponse struct {
	Uptime int64 `json:"uptimeNano"`
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
	response := statusResponse{nowNano() - h.startTimeNano}
	data, _ := json.Marshal(response)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
