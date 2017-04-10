package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
	}

	errorResource struct {
		Data appError `json:"data"`
	}
)

/*DisplayError para retornar el error como json*/
func DisplayError(w http.ResponseWriter, handleError error, message string, code int) {
	errorObj := appError{
		Error:      handleError.Error(),
		Message:    message,
		HTTPStatus: code,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	j, err := json.Marshal(errorResource{Data: errorObj})
	if err != nil {
		log.Printf("[Error marshal json]: %s\n", err)
	}
	w.Write(j)
}
