package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if (r.URL.Path[1:] == "date") {
		fmt.Fprintf(w, time.Now().Local().Format("2006-01-02T15:04:05-0700\n"))
	} else {
		if (r.URL.Path[1:] == "wait") {
			keys := r.URL.Query()["delay"]
			key := keys[0]
			intkey, err := strconv.Atoi(key)
			if err == nil {
				time.Sleep(time.Duration(intkey) * time.Millisecond)
			}
		}
	}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
