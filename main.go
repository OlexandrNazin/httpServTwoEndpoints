package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"
)

type control struct {
	Time                time.Time
	count               int
	id                  int
	mapIdAndTimeConnect map[int]time.Duration
	tim                 time.Timer
}

func (c *control) myHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	c.count = c.count + 1
	switch r.Method {
	case "GET":
		sortMapValue(w, c.mapIdAndTimeConnect)

	case "POST":
		{
			connectUser(w, c)
		}
	default:
		_, err := fmt.Fprintf(w, "Sorry, only GET and POST.")
		if err != nil {
			log.Println(err)
		}
	}
}

func connectUser(w http.ResponseWriter, cont *control) {

	switch cont.count {
	case 1:
		cont.tim = *time.NewTimer(5 * time.Second)
		cont.Time = time.Now()
		<-cont.tim.C
		if cont.count != 0 {
			_, err := fmt.Fprint(w, "timeOut")
			if err != nil {
				log.Println(err)
			}
			cont.count = 0
		} else {
			_, err := fmt.Fprintln(w, "User 1 connect")
			if err != nil {
				log.Println(err)
			}
		}
	case 2:
		cont.tim.Reset(0)

		cont.count = 0
		t := time.Since(cont.Time)
		cont.id++
		cont.mapIdAndTimeConnect[cont.id] = t
		fmt.Println(cont.id, "--", t)
		_, err := fmt.Fprint(w, "User 2 connect, waiting time : ", t)
		if err != nil {
			log.Println(err)
		}
	}
}

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
	//return ss
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
