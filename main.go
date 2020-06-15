package main

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
	"weatherapp/pkg/endpoints"
)

func main() {
	var wait time.Duration
	wait = 60

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	restyClient := resty.New()
	server := endpoints.NewServer(*restyClient)
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/weather", server.GetWeatherByName)

	srv := &http.Server {
		Addr:         fmt.Sprintf("0.0.0.0:%s", server.GetConfig().APIPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
