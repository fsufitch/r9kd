package channels

import "net/http"

import "github.com/fsufitch/r9kd/auth"

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
	w.Write([]byte("Create new channel success!\n"))
}
