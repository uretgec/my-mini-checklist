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

// Request Obj
//
// swagger:parameters ApiSet
type RequestObj struct {

	// in: formData
	// required: true
	// default: test
	Key string `json:"key"`

	// in: formData
	// required: true
	// default: 12
	Value string `json:"value"`
}

// Response obj
//
// swagger:response ResultObj
type ResultObj struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// swagger:operation GET / home Home
// Home Page Handler
// ---
// produces:
//  - application/json
// responses:
//	 '200':
//     description: Welcome
//     schema:
//	     "$ref": "#/responses/ResultObj"
func home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(&ResultObj{resultSuccess, "Welcome Home", nil})
	})
}

// swagger:operation GET /version version Version
// Build Version Page Handler
// ---
// produces:
//  - application/json
// responses:
//	 '200':
//     description: Version Number
//     schema:
//	     "$ref": "#/responses/ResultObj"
func buildVersion() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(&ResultObj{resultSuccess, "Build version found", *flagBuildVersion})
	})
}

// swagger:operation GET /api api ApiHome
// Api Home Page Handler
// ---
// produces:
//  - application/json
// responses:
//	 '200':
//     description: Welcome Api Home
//     schema:
//	     "$ref": "#/responses/ResultObj"
func apiHome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(&ResultObj{resultSuccess, "Api Home", nil})
	})
}

// swagger:operation POST /api/set apiSet ApiSet
// Api Set/Update Handler
// ---
// consumes:
//  - application/x-www-form-urlencoded
// produces:
//  - application/json
// parameters:
//	- name: key
//	  description: Name of key
//	  required: true
//	  type: string
//	- name: value
//	  description: Value of key
//	  required: true
//	  type: string
// responses:
//	 '200':
//     description: Key updated successfully!
//     schema:
//	     "$ref": "#/responses/ResultObj"
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
