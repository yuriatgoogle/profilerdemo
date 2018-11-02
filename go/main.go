package main

import (
	"fmt"
	"log"
	"math/rand"
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
	// spin CPU for that many seconds
	blockCPU(rand.Intn(5))
	fmt.Fprintln(w, "blocked CPU!")
	log.Printf("[golang-profiler:handle] Exited")
}

func blockCPU(delay int) {
	log.Printf("blocking CPU")
	result := 0
	timeToExit := time.Now().Local().Add(time.Second * time.Duration(delay))
	for true {
		r := rand.New(rand.NewSource(99))
		result += r.Int() * r.Int()
		if time.Now().After(timeToExit) {
			log.Print("exiting loop")
			return
		}
	}
}
