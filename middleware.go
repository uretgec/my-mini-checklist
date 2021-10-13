package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func QueryParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		switch r.Method {
		case "GET":

			for k, v := range r.URL.Query() {
				fmt.Printf("Query: %s: %s\n", k, v)
			}

			r.Body = ioutil.NopCloser(strings.NewReader(r.URL.RawQuery))
			r.ContentLength = int64(len(r.URL.RawQuery))

		case "POST":

			r.ParseForm()
			for k1, v1 := range r.Form {
				fmt.Printf("RFORM: %s: %s\n", k1, v1)

				r.URL.Query().Add(k1, v1[0])
			}

		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		}

		next.ServeHTTP(w, r)
	})
}
