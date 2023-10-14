package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	return channel, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"minhafila",   // queue name
		"go-consumer", // application consumer name
		false,         // if can delete message after consume automatic
		false,         // if is a exclusive queue
		false,         // if its no local, that means production
		false,         // if its no wait
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ch *amqp.Channel, body string) error {
	err := ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
