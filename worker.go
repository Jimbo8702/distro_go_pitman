package main

import (
	"fmt"
	"log"
	"time"
)

// Worker represents a worker node responsible for crawling tasks.
type Worker struct {
	ID          	int
	crawler     	*WebCrawler
	queue       	*MessageQueue
	database    	*Database
	maxRetries  	int
	retryInterval   time.Duration
	retryCounts 	map[string]int // Map to track the number of retries for each crawling task
}

// NewWorker creates a new Worker instance.
func NewWorker(id int, crawler *WebCrawler, queue *MessageQueue, database *Database) *Worker {
	return &Worker{
		ID:          id,
		crawler:     crawler,
		queue:       queue,
		database:    database,
		retryCounts: make(map[string]int),
	}
}

// Start starts the worker to listen for tasks from the message queue and process them.
func (w *Worker) Start() {
	for {
		message, err := w.queue.ReceiveMessage()
		if err != nil {
			log.Printf("Worker %d: Failed to receive message: %v", w.ID, err)
			// Implement retry logic by sleeping for a short duration and then retrying.
			time.Sleep(5 * time.Second)
			continue
		}	// Process the data using the web crawler with retry mechanism
		data, err := w.processWithRetry(message)
		if err != nil {
			log.Printf("Worker %d: Error processing data: %v", w.ID, err)
			// Move the failed message to the dead letter queue if all retries are exhausted
			if w.retryCounts[message] >= w.maxRetries {
				log.Printf("Worker %d: Failed to process task after max retries. Moving to dead letter queue.", w.ID)
				err = w.queue.MoveToDeadLetterQueue(message)
				if err != nil {
					log.Printf("Worker %d: Failed to move message to dead letter queue: %v", w.ID, err)
				}
			}
			// Implement retry logic by sleeping for a short duration and then retrying.
			time.Sleep(w.retryInterval)
			continue
		}

		// Save data to the database
		err = w.database.SaveData(data)
		if err != nil {
			log.Printf("Worker %d: Failed to save data to the database: %v", w.ID, err)
			// Implement retry logic by sleeping for a short duration and then retrying.
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

// processWithRetry processes the crawling task with retry mechanism.
func (w *Worker) processWithRetry(message string) (data string, err error) {
	for w.retryCounts[message] < w.maxRetries {
		// Process the crawling task
		data, err = w.crawler.ProcessData(message)
		if err == nil {
			// Successfully processed the task, reset the retry count
			w.retryCounts[message] = 0
			return data, nil
		}

		// Increment the retry count and wait before the next retry
		w.retryCounts[message]++
		time.Sleep(w.retryInterval)
	}

	return "", fmt.Errorf("failed to process task after max retries")
}