package main

import (
	"fmt"
	"terarium-client/rabbit"
)

func Produce(queueName string, mac string, producer *rabbit.Producer, inch chan string, errch chan error) {
	
	producer.NewQueue(queueName + "/" + mac)
	
	for message := range inch {
		err := producer.Publish("", queueName + "/" + mac, fmt.Sprintf("New message: %s", message))
		if err != nil {
			errch <- err
		}
	}
}