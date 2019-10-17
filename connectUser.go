package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

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
		cont.mu.Lock()
		cont.mapIdAndTimeConnect[cont.id] = t
		cont.mu.Unlock()
		fmt.Println(cont.id, "--", t)
		_, err := fmt.Fprint(w, "User 2 connect, waiting time : ", t)
		if err != nil {
			log.Println(err)
		}
	}
}
