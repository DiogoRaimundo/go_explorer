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
const slidingWindowTimeSpan = -time.Minute
const fileName = "n_request_per_minute_exercise.rec"

type PersistedTimeValues struct {
	lock   sync.Mutex
	file   *os.File
	values []time.Time
}

func CreatePersistedTimeValues(filePath string) *PersistedTimeValues {
	timeValues := PersistedTimeValues{}

	var err error
	timeValues.file, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	timeValues.loadAndUpdateValues()

	return &timeValues
}

func (timeValues *PersistedTimeValues) GetValues() []time.Time {
	return timeValues.values
}

func (timeValues *PersistedTimeValues) Close() error {
	return timeValues.file.Close()
}

func (timeValues *PersistedTimeValues) loadAndUpdateValues() {
	allBytes, err := os.ReadFile(timeValues.file.Name())
	if err != nil {
		log.Printf("[ERROR] Unable to read file (%v)\n", err)
		return
	}

	fileData := strings.ReplaceAll(string(allBytes), "\r", "")

	lines := strings.Split(fileData, "\n")

	timeValues.values = make([]time.Time, 0)
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
			continue
		}

		if timeValue.Before(currentTime) && timeValue.After(timeLimit) {
			timeValues.values = append(timeValues.values, timeValue)
			newFileBuilder.WriteString(line)
			newFileBuilder.WriteByte('\n')
		}
	}

	err = os.WriteFile(timeValues.file.Name(), []byte(newFileBuilder.String()), 0777)
	if err != nil {
		log.Printf("[ERROR] Unable to update file (%v)\n", err)
	}
}

func (timeValues *PersistedTimeValues) AddValueAndUpdate() int {
	currentTime := time.Now()
	timeLimit := currentTime.Add(slidingWindowTimeSpan)

	nToRemove := 0
	for _, value := range timeValues.values {
		if value.Before(timeLimit) {
			nToRemove++
		}
	}

	timeValues.lock.Lock()
	_, err := timeValues.file.WriteString(fmt.Sprintf("%s\n", currentTime.Format(time.UnixDate)))
	timeValues.values = append(timeValues.values[nToRemove:], currentTime)
	valuesCount := len(timeValues.values)
	timeValues.lock.Unlock()

	if err != nil {
		log.Printf("[ERROR] Unable to write to file (%s)", err)
	}

	return valuesCount
}

func Run() {
	timeValues := CreatePersistedTimeValues(fileName)
	defer timeValues.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCount := timeValues.AddValueAndUpdate()

		writeSerialized(requestCount, w, r)
	})

	server := http.Server{
		Addr:    serverHostname,
		Handler: mux,
	}

	serverShutdownErrorSignal := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("[ERROR] Unable to shutdown server (%v)", err)
		}
		close(serverShutdownErrorSignal)
	}()

	server.ListenAndServe()

	<-serverShutdownErrorSignal
}

func writeSerialized(v any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	jsonBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("[ERROR] Unable to serialize response (%s)", err)
		return
	}

	w.Write(jsonBytes)
}
