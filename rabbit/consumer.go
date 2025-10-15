package rabbit

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
	chanel  *amqp.Channel
	username string
	password string
	mac string
}

func NewConsumer(username string, password string, mac string) (*Consumer, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@localhost:5672/", username, password))
	if err != nil {
		return nil, err
	}
	
	chanel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	
	c := &Consumer{
		conn: conn,
		chanel: chanel,
		username: username,
		password: password,
		mac: mac,
	}
	return c, nil
}

func (c *Consumer) NewQueue(name string) error {
	_, err := c.chanel.QueueDeclare(
		name,
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	
	return nil
}

func (c *Consumer) Consume(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := c.chanel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (c *Consumer) Close() {
	c.chanel.Close()
	c.conn.Close()
}