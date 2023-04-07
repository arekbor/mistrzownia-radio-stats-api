package main

import (
	"log"
	"os"
	"time"

	"github.com/arekbor/mistrzownia-radio-stats-api/api"
	"github.com/arekbor/mistrzownia-radio-stats-api/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalln(err)
		return
	}

	sc := store.New(
		os.Getenv("PQ_HOST"),
		os.Getenv("PQ_PORT"),
		os.Getenv("PQ_USER"),
		os.Getenv("PQ_PWD"),
		os.Getenv("PQ_DBNAME"),
		10*time.Second,
	)
	store, err := sc.Init()
	if err != nil {
		log.Fatalln(err)
		return
	}

	a := api.New(os.Getenv("API_ADDR"), store, 10*time.Second, 10*time.Second)
	err = a.Init()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
