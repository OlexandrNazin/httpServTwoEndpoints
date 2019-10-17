package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type control struct {
	mu                  sync.Mutex
	Time                time.Time
	count               int
	id                  int
	mapIdAndTimeConnect map[int]time.Duration
	tim                 time.Timer
}

func main() {
	c := &control{
		mapIdAndTimeConnect: make(map[int]time.Duration),
	}

	http.HandleFunc("/", c.myHandler)
	fmt.Printf("Starting server for port :2019...\n")
	if err := http.ListenAndServe(":2019", nil); err != nil {
		log.Fatal(err)
	}
}
