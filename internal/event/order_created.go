package event

import "time"

type OrderCreated struct {
	Name     string
	DateTime string
	Payload  interface{}
}

func NewOrderCreated() *OrderCreated {
	return &OrderCreated{
		Name:     "OrderCreated",
		DateTime: time.Now().Format(time.RFC3339),
	}
}

func (e *OrderCreated) GetName() string {
	return e.Name
}

func (e *OrderCreated) GetDateTime() string {
	return e.DateTime
}

func (e *OrderCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}
