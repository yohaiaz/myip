package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	graceFullShutDown := make(chan os.Signal, 1)

	signal.Notify(graceFullShutDown)

	fmt.Println("starting server ...")

	http.HandleFunc("/", serverHandler)

	http.ListenAndServe(":5000", nil)

	<-graceFullShutDown
}

func serverHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "welcome, \n")

	fmt.Fprintf(w, "now: %v utc \n", time.Now().UTC())

	fmt.Fprintf(w, "your address is: %s \n", r.RemoteAddr)
}
