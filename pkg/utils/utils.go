package utils

import (
	"github.com/martinkaburu/warehouse-manager/pkg/errors"
	"net/http"
	"os"
)

func WriteResponse(w http.ResponseWriter, response []byte) {
	_, err := w.Write(response)
	if err != nil {
		WriteErrorResponse(w, err, errors.ServerError{
			Message: "Could not encode the response",
			Code:    "RESPONSE_ENCODING_ERROR",
		})
		return
	}
	return
}

func WriteErrorResponse(w http.ResponseWriter, err error, errDescription errors.ServerError) {
	SendHTTPError(w, errDescription, http.StatusInternalServerError)
	return
}

func SendHTTPError(w http.ResponseWriter, err errors.ServerError, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}

func SetResponseContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
