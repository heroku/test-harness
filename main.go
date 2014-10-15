package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	sleepTime := r.Form["after"][0]

requestHeader := r.Header
fmt.Printf("%s", requestHeader)

	fmt.Printf("Sleeping for %s\n", sleepTime)
	sleepDuration, _ := strconv.Atoi(sleepTime)
	for i := 0; i < sleepDuration; i++ {
		time.Sleep(time.Second)
		fmt.Println("tick")
		fmt.Fprint(w, "tick\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}

	fmt.Fprint(w, "OK\n")
}

func main() {
	http.HandleFunc("/OK", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
