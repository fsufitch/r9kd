package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/fsufitch/r9kd/server/util"
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

type err40XResponse struct {
	Success      bool   `json:"success"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}

// Write401Response does what it says on the can
func Write401Response(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", "Bearer realm=\"API key authentication required\"")
	util.WriteHTTPErrorResponse(w, 401, "Authorization header not specified, or contained invalid base64 Bearer key")
}

// Write403Response does what it says on the can
func Write403Response(w http.ResponseWriter, deniedAction string) {
	if deniedAction == "" {
		deniedAction = "<unspecified>"
	}
	message := fmt.Sprintf("Your API key does not allow you to: %s\n", deniedAction)
	util.WriteHTTPErrorResponse(w, 403, message)
}
