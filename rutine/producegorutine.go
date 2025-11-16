package rutine

import (
	"terarium-client/rabbit"
	"terarium-client/rabbit/dto/terarium"
)

func Produce(queueName string, mac string, producer *rabbit.Producer, inch chan terarium.TerariumOutDto, errch chan error) {
	
	producer.NewQueue(queueName + "/" + mac)
	
	for message := range inch {
		data := message.JsonToString()
		if data == "" {
			continue
		}
		
		err := producer.Publish("", queueName + "/" + mac, data)
		if err != nil {
			errch <- err
		}
	}
}