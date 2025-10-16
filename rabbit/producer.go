package rabbit

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	chanel *amqp.Channel
	conn   *amqp.Connection
	mac string
}

func NewProducer(username string, password string, host string, port string, mac string) (*Producer, error) {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port))
	if err != nil {
		return nil, err
	}
	
	chanel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	
	return &Producer{
		conn: connection,
		chanel: chanel,
		mac: mac,
	}, nil
}

func (p *Producer) NewQueue(name string) error {
	_, err := p.chanel.QueueDeclare(
		name,
		true,
		true,
		false,
		false,
		amqp.Table{
			
		},
	)
	if err != nil {
		return err
	}
	
	return nil
}

func (p *Producer) Publish(exchange string, queueName string, message string) error {
	return p.chanel.Publish(
		exchange,
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)
}

func (p *Producer) Close() {
	p.chanel.Close()
	p.conn.Close()
}