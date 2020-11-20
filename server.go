package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var responseCodePtr *int
var sleepPtr *int

func main() {
	portPtr := flag.String("port", "8080", "server port")
	responseCodePtr = flag.Int("responseCode", 200, "http status code")
	sleepPtr = flag.Int("sleep", 0, "sleep duration in seconds")

	flag.Parse()

	fmt.Printf("Starting server on port %s ...\n", *portPtr)

	http.HandleFunc("/", handle)
	http.ListenAndServe(":"+*portPtr, nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if *sleepPtr > 0 {
		time.Sleep(time.Duration(*sleepPtr) * time.Second)
	}

	w.WriteHeader(*responseCodePtr)
	w.Write([]byte("Response body"))
}
