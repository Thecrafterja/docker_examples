package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	println("Starting webserver")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World through the Web from Go!"))
	})

	timeRoute := os.Getenv("TIME_ROUTE")

	router.Get("/"+timeRoute, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Aktueller Zeitpunkt: " + time.Now().Local().Format("01.02.2006 15:04:05")))
	})

	router.Get("/count", func(w http.ResponseWriter, r *http.Request) {
		data, readErr := os.ReadFile("./data/count.txt")
		if readErr != nil {
			http.Error(w, "Could not read count.txt", http.StatusInternalServerError)
			return
		}

		var count int
		if string(data) == "" {
			count = 0
		} else {
			var convErr error
			count, convErr = strconv.Atoi(string(data))
			if convErr != nil {
				http.Error(w, "Could not convert file content into integer", http.StatusInternalServerError)
				return
			}
		}
		count++

		writeErr := os.WriteFile("./data/count.txt", []byte(strconv.Itoa(count)), 0666)
		if writeErr != nil {
			http.Error(w, "Could not save count.txt", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Count: " + strconv.Itoa(count)))
	})

	http.ListenAndServe(":10101", router)
}
