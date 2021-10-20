// My Mini Checklist Api
//
// How to make simple key value based checklist api with golang stdlib for not master users?
//
//	Schemes: http
//	Host: localhost:3000
//	BasePath: /
//	Version: 1.0.3
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: Uretgec<iletisim@uretgec.com>
//
//	Consumes:
//	- application/x-www-form-urlencoded
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	flagAddr           = flag.String("addr", ":3000", "server listen addr")
	flagDbPath         = flag.String("dbpath", "./store.db", "db file name with folder path")
	flagHttpLogPath    = flag.String("logpath", "./service-http.log", "db file name with folder path")
	flagBuildVersion   = flag.String("version", "1.0", "build version of service")
	flagSyncDbInterval = flag.Duration("bgsave", time.Duration(1*time.Minute), "dump memory to file periodly")
)

var store *Store

func main() {
	// Flags Parse: flagAddr, flagDbPath, flagSyncDbInterval
	flag.Parse()

	// logger init
	logfile, err := os.OpenFile(*flagHttpLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		log.Fatal(err)
	}

	defer logfile.Close()

	rlogger := log.New(logfile, "http: ", log.LstdFlags)

	// All routes
	router := newRouter()

	// Create or open file database
	storeDb, err := os.OpenFile(*flagDbPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Storedb file not opened: %s", err.Error())
		os.Exit(1)
	}

	// Start new store and if possible database import from file db to memory
	store = NewStore()
	store.Load(storeDb)
	storeDb.Close()

	// Sync memory to file db
	go syncDb()

	// Http server options init
	server := &http.Server{
		Addr:         *flagAddr,
		Handler:      Logging(rlogger, QueryParams(router)),
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Http server start
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln("Could not start server")
			os.Exit(1)
		}
	}()

	// Listen server quit or something happened and notify channel
	// And sync memory data to file database
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	<-c
	bgSave()
}

// Periodly sync from in-memory data to flat file
func syncDb() {
	syncTicker := time.NewTicker(*flagSyncDbInterval)
	for _ = range syncTicker.C {
		bgSave()
	}
}

// Background save process
// First all mem data saves to temp file
// After renames temp filename to real db filename
func bgSave() {
	tempFileName := fmt.Sprintf("%s-%d", *flagDbPath, rand.Intn(2000))
	storeDb, err := os.OpenFile(tempFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("Storedb file not opened")
		os.Exit(1)
	}

	err = store.Sync(storeDb, false)
	if err != nil {
		log.Fatalln("Storedb sync process failed")
		os.Exit(1)
	}

	storeDb.Close()

	err = os.Rename(tempFileName, *flagDbPath)
	if err != nil {
		log.Fatalln("Storedb rename process failed")
	}
}
