package main

import (
	"fmt"
	"log"
	"net/http"
)

func (c *control) myHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	c.count += 1
	switch r.Method {
	case "GET":
		c.mu.Lock()
		var m = c.mapIdAndTimeConnect
		c.mu.Unlock()
		sortMapValue(w, m)

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
