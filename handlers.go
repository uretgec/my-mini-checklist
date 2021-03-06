package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	resultSuccess = "success"
	resultError   = "error"
)

type ResultObj struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Home Page Handler
func home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(&ResultObj{resultSuccess, "Welcome Home", nil})
	})
}

// Build Version Page Handler
func buildVersion() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(&ResultObj{resultSuccess, "Build version found", *flagBuildVersion})
	})
}

// Api Home Page Handler
func apiHome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(&ResultObj{resultSuccess, "Api Home", nil})
	})
}

// Api Set/Update Handler
func apiStoreSet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		k := r.FormValue("key")
		v := r.FormValue("value")

		result := &ResultObj{resultError, k + " Not found!", nil}
		if k != "" || v != "" {
			store.Set(k, v)
		}

		val := store.Get(k)
		if val != "" {
			result.Status = resultSuccess
			result.Message = k + " key updated"
			result.Result = val
		}

		json.NewEncoder(w).Encode(result)
	})
}

// Api Get Handler
func apiStoreGet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		k := r.FormValue("key")

		result := &ResultObj{resultError, k + " Not found!", nil}
		val := store.Get(k)
		if val != "" {
			result.Status = resultSuccess
			result.Message = k + " key found"
			result.Result = val
		}

		json.NewEncoder(w).Encode(result)
	})
}

// Api Del Handler
func apiStoreDel() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		k := r.FormValue("key")

		store.Del(k)

		result := &ResultObj{resultError, k + " Not found!", nil}
		val := store.Get(k)
		if val == "" {
			result.Status = resultSuccess
			result.Message = k + " key deleted"
		}

		json.NewEncoder(w).Encode(result)
	})
}

// Api List Handler
func apiStoreList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := &ResultObj{resultSuccess, "", nil}
		result.Message = strconv.Itoa(store.Stats()) + " items found"
		result.Result = store.GetAll()

		json.NewEncoder(w).Encode(result)
	})
}

// Api Stats Handler
func apiStoreStats() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := &ResultObj{resultSuccess, "", nil}
		result.Result = strconv.Itoa(store.Stats())

		json.NewEncoder(w).Encode(result)
	})
}

// Api Flush Handler
func apiStoreFlush() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		store.Flush()

		result := &ResultObj{resultError, "Not flushed!", nil}
		if store.Stats() == 0 {
			result.Status = resultSuccess
			result.Message = "Flushed!"
			result.Result = 0
		}

		json.NewEncoder(w).Encode(result)
	})
}
