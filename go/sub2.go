package main

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
)

func main() {
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	subscriber.Connect("tcp://localhost:5556")
	subscriber.SetSubscribe("questions")

	for {
		msg, err := subscriber.RecvMessage(0)
		if err != nil {
			break
		}

		topic := msg[0]
		data := msg[1]

		if topic != "questions" {
			panic("topic != questions")
		}

		fmt.Println(data)
	}
}
