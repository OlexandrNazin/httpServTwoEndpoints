package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"
)

func sortMapValue(w http.ResponseWriter, m map[int]time.Duration) {

	type kv struct {
		Key   int
		Value time.Duration
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for _, kv := range ss {
		_, err := fmt.Fprintf(w, "%v, %v\n", kv.Key, kv.Value)
		if err != nil {
			log.Println(err)
		}
	}
	if len(ss) > 50000 {
		ss = ss[50000:]
	}
}
