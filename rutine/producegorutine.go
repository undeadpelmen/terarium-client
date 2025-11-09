package rutine

import (
	"terarium-client/rabbit"
)

func Produce(queueName string, mac string, producer *rabbit.Producer, inch chan OutTerMsg, errch chan error) {
	
	producer.NewQueue(queueName + "/" + mac)
	
	for message := range inch {
		data, err := message.JsonToString()
		if err != nil {
			errch <- err
			continue
		}
		
		err = producer.Publish("", queueName + "/" + mac, data)
		if err != nil {
			errch <- err
		}
	}
}