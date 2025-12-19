package main

import (
	"log"

	"feed/internal/broker"
)

func main() {
	if err := broker.Connect(); err != nil {
		log.Panic(err)
	}
	ch, err := broker.Channel()
	if err != nil {
		log.Panic(err.Error())
	}
	defer func() {
		if err := ch.Close(); err != nil {
			log.Panic(err.Error())
		}
		if err := broker.Disconnect(); err != nil {
			log.Panic(err.Error())
		}
	}()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Panic(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err.Error())
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-forever
}
