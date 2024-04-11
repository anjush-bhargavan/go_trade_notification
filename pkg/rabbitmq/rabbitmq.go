package rabbitmq

import (
	"encoding/json"
	"github.com/anjush-bhargavan/go_trade_notification/pkg/config"
	"github.com/anjush-bhargavan/go_trade_notification/pkg/service"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func ConsumeNotificationMessages(cnfg *config.Config) {


	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	// Declare the queue
	q, err := ch.QueueDeclare(
		"notification_queue", // queue name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Process incoming messages
	for msg := range msgs {
		// Extract notification details from the message
		var notification service.Messages
		err := json.Unmarshal(msg.Body, &notification)
		if err != nil {
			log.Printf("Error decoding message body: %v", err)
			continue
		}

		// Send notification email using the details
		err = service.SendEmail(cnfg, notification)
		if err != nil {
			log.Printf("Error sending notification email: %v", err)
		}
	}
}