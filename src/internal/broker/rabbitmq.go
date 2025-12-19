package broker

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection

func Connect() error {
	var err error
	conn, err = amqp.Dial(os.Getenv("RABBITMQ_CONNECTION"))
	if err != nil {
		return err
	}

	return nil
}

func Channel() (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func Disconnect() error {
	if err := conn.Close(); err != nil {
		return err
	}

	return nil
}
