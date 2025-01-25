package graph

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	endpointCounter sync.Map
)

// addOne increments the counter for the given endpoint
func addOne(endpoint string) {
	// Increment the counter atomically
	value, _ := endpointCounter.LoadOrStore(endpoint, new(uint32))
	counter := value.(*uint32)
	*counter++

	// Check if the counter reached 100
	if *counter >= 3 {
		*counter = 0               // Reset the counter
		go saveTimestamp(endpoint) // Save the timestamp asynchronously
	}
}

// saveTimestamp writes the endpoint and timestamp to a file
func saveTimestamp(endpoint string) {
	timestamp := time.Now().Format(time.RFC3339)
	filename := "ql_timestamps.log"

	// Format the log entry
	logEntry := fmt.Sprintf("%s - Endpoint: %s\n", timestamp, endpoint)

	// Append the log entry to the file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(logEntry); err != nil {
		fmt.Printf("Error writing log entry: %v\n", err)
	}
}
