package app

import (
	"github.com/MinhWalker/store_items-api/src/controllers"
	"net/http"
)

func mapUrls()  {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
