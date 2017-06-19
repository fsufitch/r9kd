package util

import (
	"encoding/json"
	"net/http"

	"github.com/fsufitch/r9kd/model"
)

type httpErrorResponse struct {
	Success      bool   `json:"success"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}

// WriteHTTPErrorResponse is a boilerplate function for writing a plain JSON error
func WriteHTTPErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	data, _ := json.MarshalIndent(httpErrorResponse{false, code, message}, "", "  ")
	w.Write(data)
}

// WriteSerializableJSON is a boilerplate function to serialize and write a JSON object
func WriteSerializableJSON(w http.ResponseWriter, code int, obj model.Serializable) {
	data, _ := obj.Serialize()
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

type httpBasicSuccessResponse struct {
	Success bool `json:"success"`
}

func (r httpBasicSuccessResponse) Serialize() ([]byte, error) {
	return json.Marshal(r)
}

//HTTPBasicSuccessResponse is a shortcut for printing a simple {"success": true} JSON response
var HTTPBasicSuccessResponse = httpBasicSuccessResponse{true}
