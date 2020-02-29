package mqClient

import(
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
)
func SendHelper(message Message){
	// connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// a channel, which is where most of the API for getting things done resides
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	body := message
	b,err :=json.Marshal(body)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	log.Printf(" [x] Sent %s", body)
	FailOnError(err, "Failed to publish a message")
}