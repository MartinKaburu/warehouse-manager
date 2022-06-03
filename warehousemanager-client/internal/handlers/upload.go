package handlers

import (
	"encoding/csv"
	"encoding/json"
	"github.com/jszwec/csvutil"
	"github.com/martinkaburu/warehouse-manager/pkg/errors"
	"github.com/martinkaburu/warehouse-manager/pkg/models"
	"github.com/martinkaburu/warehouse-manager/pkg/utils"
	"github.com/martinkaburu/warehouse-manager/warehousemanager-client/internal/service"
	"net/http"
)

//UploadHandler ...
func UploadHandler(port *int) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		file, _, err := request.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		csvReader := csv.NewReader(file)

		userHeader, _ := csvutil.Header(models.OrderEntry{}, "csv")
		decoder, _ := csvutil.NewDecoder(csvReader, userHeader...)

		service.Stream(decoder, port)

		utils.SetResponseContentType(w)

		response, err := json.Marshal(&utils.HealthCheck{Message: "OK"})
		if err != nil {
			utils.SendHTTPError(w, errors.ServerError{
				Message: "Could not marshal the response",
				Code:    "RESPONSE_MARSHAL_ERROR",
			}, http.StatusInternalServerError)
			return
		}

		utils.WriteResponse(w, response)
		return
	}
}
