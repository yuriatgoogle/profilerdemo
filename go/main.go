package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/profiler"
)

const (
	serviceName = "golang-profiler"
)

var (
	projectID = "thegrinch-project"
)

func main() {

	log.Printf("[golang-profiler:main] ProjectID: %v", projectID)

	if err := profiler.Start(profiler.Config{
		ProjectID:      projectID,
		Service:        serviceName,
		ServiceVersion: "0.0.1",
		DebugLogging:   true}); err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[golang-profiler:handle] Entered")
	// spin CPU
	blockCPU()
	fmt.Fprintln(w, "Hello Henry!")
	log.Printf("[golang-profiler:handle] Exited")
}

func blockCPU() {
	<-time.After(time.Duration(10))
}
