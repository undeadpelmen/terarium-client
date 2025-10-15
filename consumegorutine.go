package main

import "terarium-client/rabbit"

func Consume(queueName string, mac string, consumer *rabbit.Consumer, outch chan string, errch chan error) {
	msgs, err := consumer.Consume(queueName + "/" + mac)
	if err != nil {
		errch <- err
		return
	}
	
	for msg := range msgs {
		outch <- string(msg.Body)
	}
}