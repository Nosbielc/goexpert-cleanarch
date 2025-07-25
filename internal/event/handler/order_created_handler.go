package handler

import (
	"encoding/json"
	"fmt"

	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface) {
	jsonOutput, _ := json.Marshal(event.GetPayload())
	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish("amq.direct", "", false, false, msgRabbitmq)
	fmt.Println("OrderCreatedHandler: ", string(jsonOutput))
}
