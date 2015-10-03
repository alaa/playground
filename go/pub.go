package main

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
	"time"
)

func main() {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind("tcp://127.0.0.1:5556")
	time.Sleep(time.Second)

	for {
			fmt.Println("publishing message!")
		_, err := publisher.SendMessage("users", "Logged in user B")
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)


		_, err = publisher.SendMessage("questions", "Add question 1")
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}
