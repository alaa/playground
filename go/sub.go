package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	subscriber.Connect("tcp://localhost:5556")
	subscriber.SetSubscribe("users")

	for {
		msg, err := subscriber.RecvMessage(0)
		if err != nil {
			break
		}

		topic := msg[0]
		data := msg[1]

		if topic != "users" {
			panic("topic != users")
		}

		fmt.Println(data)
	}
}
