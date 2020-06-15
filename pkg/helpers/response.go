package helpers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// RespondWithError forms required JSON error payload together with HTTP error code.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON forms required JSON payload together with HTTP code.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		errMsg := fmt.Sprintf("Cannot form JSON payload! Details: %v", err)
		log.Error(errMsg)
		code = http.StatusInternalServerError
		response = []byte(`"error": "Cannot form JSON payload. Please contact system administrator"`)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
