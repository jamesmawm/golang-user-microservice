package core

import (
	"encoding/json"
	"net/http"
)

func OnPing(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"ok":   true,
		"pong": true,
	}

	responseBytes, _ := json.Marshal(response)
	_, _ = w.Write(responseBytes)
	return
}
