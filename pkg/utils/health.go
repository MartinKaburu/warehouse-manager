package utils

import (
	"encoding/json"
	"github.com/martinkaburu/warehouse-manager/pkg/errors"
	"net/http"
)

type HealthCheck struct {
	Message string `json:"message"`
}

func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {

		SetResponseContentType(w)

		response, err := json.Marshal(&HealthCheck{Message: "OK"})
		if err != nil {

			SendHTTPError(w, errors.ServerError{
				Message: "Could not marshal the response",
				Code:    "RESPONSE_MARSHAL_ERROR",
			}, http.StatusInternalServerError)
			return
		}

		WriteResponse(w, response)
	}
}
