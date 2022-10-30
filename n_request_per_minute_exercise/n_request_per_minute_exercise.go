package nRequestPerMinuteExercise

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

const serverHostname = "localhost:8080"
const fileName = "n_request_per_minute_exercise.rec"

var timeValues = make([]time.Time, 0)
var recordFile *os.File
var mu sync.Mutex

const slidingWindowTimeSpan = -time.Minute

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	readTimeValues()

	var err error
	recordFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0777)
	checkError(err)
	defer recordFile.Close()

	server := http.Server{
		Addr:    serverHostname,
		Handler: mux,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		// close(idleConnectionsClosed)
	}()

	server.ListenAndServe()

	log.Printf("Bye bye")
}

func readTimeValues() {
	allBytes, err := os.ReadFile(fileName)
	checkError(err)
	fileData := strings.ReplaceAll(string(allBytes), "\r", "")

	lines := strings.Split(fileData, "\n")

	valuesFromFile := make([]time.Time, 0)
	newFileBuilder := strings.Builder{}

	currentTime := time.Now()
	timeLimit := currentTime.Add(slidingWindowTimeSpan)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		timeValue, err := time.Parse(time.UnixDate, line)
		if err != nil {
			log.Printf("[ERROR] Unable to parse line \"%v\"\n", line)
		}

		if timeValue.Before(currentTime) && timeValue.After(timeLimit) {
			valuesFromFile = append(valuesFromFile, timeValue)
			newFileBuilder.WriteString(line)
			newFileBuilder.WriteByte('\n')
		}
	}

	timeValues = valuesFromFile

	err = os.WriteFile(fileName, []byte(newFileBuilder.String()), 0777)
	checkError(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestCount := getUpdatedRequestCount()

	writeSerialized(requestCount, w, r)
}

func getUpdatedRequestCount() int {
	currentTime := time.Now()
	timeLimit := currentTime.Add(slidingWindowTimeSpan)

	nToRemove := 0
	for _, value := range timeValues {
		if value.Before(timeLimit) {
			nToRemove++
		}
	}

	mu.Lock()
	_, err := recordFile.WriteString(fmt.Sprintf("%s\n", currentTime.Format(time.UnixDate)))
	mu.Unlock()

	if err != nil {
		log.Printf("[ERROR] Unable to write to file (%s)", err)
	}

	timeValues = append(timeValues[nToRemove:], currentTime)

	return len(timeValues)
}

func writeSerialized(v any, w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(v)
	checkError(err)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
