package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//ADD env variable and persistent storage

func main() {
	println("Starting webserver")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World through the Web from Go!"))
	})

	router.Get("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Aktueller Zeitpunkt: " + time.Now().Local().Format("01.02.2006 15:04:05")))
	})

	http.ListenAndServe(":10101", router)
}
