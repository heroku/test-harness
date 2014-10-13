package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	sleepTime := r.Form["after"][0]
	sleepDuration, _ := time.ParseDuration(sleepTime)
	fmt.Printf("Sleeping for %q\n", sleepDuration)
	time.Sleep(sleepDuration)

	fmt.Fprint(w, "OK")
}

func main() {
	http.HandleFunc("/OK", handler)
	http.ListenAndServe(":8080", nil)
}
