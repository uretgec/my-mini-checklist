package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Logging middleware handler collects request data
func Logging(rlogger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var rBody []byte
		if r.Body != nil {
			rBody, _ = ioutil.ReadAll(r.Body)

			// Restore Body content
			r.Body = ioutil.NopCloser(bytes.NewBuffer(rBody))
		}

		next.ServeHTTP(w, r)

		rlogger.Printf("%s %s %s %s", r.Method, r.RequestURI, string(rBody), time.Since(start))
	})
}

/*
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				rlogger.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}*/

// QueryParams middleware handler makes POST and GET methods equal
func QueryParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		switch r.Method {
		case "GET":

			/*for k, v := range r.URL.Query() {
				fmt.Printf("Query: %s: %s\n", k, v)
			}*/

			r.Body = ioutil.NopCloser(strings.NewReader(r.URL.RawQuery))
			r.ContentLength = int64(len(r.URL.RawQuery))

		case "POST":

			r.ParseForm()
			for k1, v1 := range r.Form {
				//fmt.Printf("RFORM: %s: %s\n", k1, v1)

				r.URL.Query().Add(k1, v1[0])
			}

		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		}

		next.ServeHTTP(w, r)
	})
}
