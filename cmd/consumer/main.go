package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

const queName = "testQueue"

func main() {
	fmt.Println("Consumer application run")

	con, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer con.Close()
	fmt.Println("Succesfully connected to our rabbitmq instance")

	ch, err := con.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(queName, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}


	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Printf("Received msg: %s\n", msg.Body)
		}
	}()
	fmt.Println("Successfully connected to rabbitmq")
	fmt.Println(" [*] - waiting for messages")
	<-forever
}
