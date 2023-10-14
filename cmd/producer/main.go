package main

import "github.com/luksrocha/events-golang/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, "Hello world")
}
