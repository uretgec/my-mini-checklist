package main

import "net/http"

func newRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", home())
	router.Handle("/version", buildVersion())

	router.Handle("/api", apiHome())
	router.Handle("/api/set", apiStoreSet())
	router.Handle("/api/get", apiStoreGet())
	router.Handle("/api/del", apiStoreDel())
	router.Handle("/api/list", apiStoreList())
	router.Handle("/api/stats", apiStoreStats())
	router.Handle("/api/flush", apiStoreFlush())

	return router
}
