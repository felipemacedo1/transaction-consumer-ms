package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"transaction_queue", // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	// go func() {
	// 	for d := range msgs {
	// 		var transaction dto.Transaction
	// 		if err := json.Unmarshal(d.Body, &transaction); err != nil {
	// 			log.Printf("Error decoding JSON: %v", err)
	// 			continue
	// 		}

	// 		// Process the transaction here
	// 		fmt.Printf("Received transaction: %+v\n", transaction)
	// 	}
	// }()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
