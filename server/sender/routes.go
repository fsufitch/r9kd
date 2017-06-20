package sender

import "github.com/gorilla/mux"

// RegisterSenderRoutes registers the appropriate routes for REST channel endpoints
func RegisterSenderRoutes(channelStringID string, router *mux.Router) {
	sub := router.PathPrefix("/sender").Subrouter()

	sub.HandleFunc("/channel/{channel_id}/sender/{sender_id}", getSender).Methods("GET")
	sub.HandleFunc("/channel/{channel_id}/sender/{sender_id}/clear", clearSenderBans).Methods("POST")
}
