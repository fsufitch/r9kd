package server

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/r9kd/db"
	"github.com/gorilla/mux"
)

func createRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/status", http.StatusMovedPermanently)
	})
	router.Handle("/status", newStatusHandler())
	return
}

// RunServer starts the r9kd server listening on the given port
func RunServer(port string) (err error) {
	serveStr := fmt.Sprintf(":%s", port)

	err = db.Connect("", false).Error
	if err != nil {
		return
	}

	err = db.Migrate(db.GetCachedConnection().Conn)
	if err != nil {
		return
	}

	fmt.Printf("Starting r9kd server on address: %s\n", serveStr)
	return http.ListenAndServe(serveStr, createRouter())
}
