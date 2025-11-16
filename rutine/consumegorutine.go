package rutine

import (
	"terarium-client/rabbit"
	"terarium-client/rabbit/dto/terarium"
)

func Consume(queueName string, ter *terarium.Tererarium, consumer *rabbit.Consumer, outch chan string, errch chan error) {
	
	consumer.NewQueue(queueName + "/" + ter.Mac)
	
	msgs, err := consumer.Consume(queueName + "/" + ter.Mac)
	if err != nil {
		errch <- err
		return
	}
	
	for msg := range msgs {
		body := string(msg.Body)
		
		message := terarium.TerariumInDto{}
		err := message.JsonFromString(body)
		if err != nil {
			errch <- err
			return
		}
		
		ter.Animal = message.Animal
		ter.Id = message.TerariumId
		ter.Mac = message.Mac
		ter.Name = message.Name
		
		outch <- message.Message.Message
	}
}