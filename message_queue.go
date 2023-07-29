package main

import (
	"github.com/streadway/amqp"
)

type MessageQueue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// NewMessageQueue creates a new MessageQueue instance.
func NewMessageQueue() (*MessageQueue, error) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MessageQueue{
		conn: conn,
		ch:   ch,
	}, nil
}

// SendMessage sends a message to the message queue.
func (mq *MessageQueue) SendMessage(message string) error {
	err := mq.ch.Publish(
		"",           	 // exchange
		"crawlerQueue", // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return err
}

// ReceiveMessage receives a message from the message queue.
func (mq *MessageQueue) ReceiveMessage() (string, error) {
	msgs, err := mq.ch.Consume(
		"crawlerQueue", // queue name
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		return "", err
	}

	for d := range msgs {
		return string(d.Body), nil
	}

	return "", nil
}

//implement
func (mq *MessageQueue) MoveToDeadLetterQueue(message string) error {
	return nil
}