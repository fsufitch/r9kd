package sender

import "github.com/gorilla/mux"

// RegisterSenderRoutes registers the appropriate routes for REST channel endpoints
func RegisterSenderRoutes(router *mux.Router) {
	router.HandleFunc("/channel/{channel_id}/sender/{sender_id}", getSender).Methods("GET")
	router.HandleFunc("/channel/{channel_id}/sender/{sender_id}/clear", clearSenderBans).Methods("POST")
}
