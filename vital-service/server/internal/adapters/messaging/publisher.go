package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/GabrielEValenzuela/RefuCare/internal/core/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queue     amqp.Queue
	queueName string
}

func NewPublisher(amqpURL string, queueName string) (*Publisher, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot open channel: %w", err)
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot declare queue: %w", err)
	}

	return &Publisher{
		conn:      conn,
		channel:   ch,
		queue:     q,
		queueName: queueName,
	}, nil
}

func (p *Publisher) PublishVitals(v *domain.Vitals) error {
	payload, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("cannot marshal vitals: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.channel.PublishWithContext(ctx,
		"",          // default exchange
		p.queueName, // routing key = queue name
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
		},
	)

	return err
}

func (p *Publisher) Close() {
	_ = p.channel.Close()
	_ = p.conn.Close()
}
