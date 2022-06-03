package main

import (
	"fmt"
	"github.com/martinkaburu/warehouse-manager/pkg/errors"
	"github.com/martinkaburu/warehouse-manager/warehousemanager-client/internal/server"
)

func main() {
	err := server.RunServer()
	if err != nil {
		fmt.Println(errors.ServerError{Message: "Unable to start server", Code: "500"})
	}
}
