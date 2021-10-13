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
	flagSyncDbInterval = flag.Duration("bgsave", time.Duration(30*time.Second), "dump memory to file intervally")
)

var store *Store

func main() {

	flag.Parse()

	router := newRouter()

	storeDb, err := os.OpenFile("./db/store.db", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Storedb file not opened: %s", err.Error())
		os.Exit(1)
	}

	store = NewStore()
	store.Load(storeDb)
	storeDb.Close()

	go syncDb()

	server := &http.Server{
		Addr:         *flagAddr,
		Handler:      Logging(QueryParams(router)),
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln("Could not start server")
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	<-c
	bgSave()
}

// NewServer returns a new Redcon server configured on "tcp" network net.
func syncDb() {
	syncTicker := time.NewTicker(*flagSyncDbInterval)
	for _ = range syncTicker.C {
		bgSave()
	}
}

// NewServer returns a new Redcon server configured on "tcp" network net.
func bgSave() {
	tempFileName := fmt.Sprintf("./db/store.db-%d", rand.Intn(2000))
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

	err = os.Rename(tempFileName, "./db/store.db")
	if err != nil {
		log.Fatalln("Storedb rename process failed")
	}
}
