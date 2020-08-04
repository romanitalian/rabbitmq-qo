package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

const queName = "testQueue"

func main() {
	fmt.Println("Rabbitmq go demo")

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

	q, err := ch.QueueDeclare(queName, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)

	msg := "Hello world. Time:  " + time.Now().String()
	err = ch.Publish("", queName, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(msg)})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully publish msg to queue")
}
