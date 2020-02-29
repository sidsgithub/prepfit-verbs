package mqConsumer

import (
	"encoding/json"
	"log"
	"github.com/streadway/amqp"
	"github.com/user/prepFitness/mqClient"
	"github.com/user/prepFitness/followers"
	"strconv"
)

func MqConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	msgrcv := mqClient.Message{}
	go func() {
		for d := range msgs{
			log.Printf("Received a message: %s", d.Body)
			json.Unmarshal(d.Body, &msgrcv)
			// log.Println(msgrcv.USERID)
			ursid,err := strconv.Atoi(msgrcv.USERID)
			log.Print(err)
			mqClient.Followers_helper(ursid,followers.Followers(ursid))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
