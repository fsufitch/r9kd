package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// GetAPIKeyFromAuthorizationHeader extracts an API key from a HTTP request
func GetAPIKeyFromAuthorizationHeader(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return ""
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		fmt.Println("Bad authorization prefix: ", auth)
		return ""
	}

	keyBase64 := auth[7:]
	keyBytes, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		fmt.Println("Error decoding base64: ", err)
		return ""
	}

	return string(keyBytes)
}

// Write401Response does what it says on the can
func Write401Response(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("WWW-Authenticate", "Bearer realm=\"API key authentication required\"")
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized -- Please specify a valid API key via the Authorization header\n"))
}

// Write403Response does what it says on the can
func Write403Response(w http.ResponseWriter, deniedAction string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(403)
	if deniedAction == "" {
		deniedAction = "<unspecified>"
	}
	message := fmt.Sprintf("403 Forbidden -- Your API key does not allow you to: %s\n", deniedAction)
	w.Write([]byte(message))
}
