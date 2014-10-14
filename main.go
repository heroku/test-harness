package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	sleepTime := r.Form["after"][0]
	sleepDuration, _ := time.ParseDuration(sleepTime)
	fmt.Printf("Sleeping for %q\n", sleepDuration)
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(w, "tick\n")
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			}
		}
	}()
	time.Sleep(sleepDuration)
	ticker.Stop()

	fmt.Fprint(w, "OK")
}

func main() {
	http.HandleFunc("/OK", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
