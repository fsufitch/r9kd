package server

import (
	"errors"
	"fmt"
	"net/http"
        "os"

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

// RunServer starts the r9kd server listening on the port from the environment
func RunServer() (err error) {
	port := os.Getenv("PORT")
        if port == "" {
                err = errors.New("PORT environment variable not set")
                return
        }
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
