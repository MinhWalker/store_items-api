package app

import (
	"github.com/MinhWalker/store_items-api/src/clients/elasticsearch"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var(
	router = mux.NewRouter()
)

func StartApplication()  {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		//Good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout: 2 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
