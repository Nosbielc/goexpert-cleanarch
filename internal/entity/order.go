package entity

import (
	"errors"
)

type OrderRepositoryInterface interface {
	Save(order *Order) error
	FindAll() ([]*Order, error)
}

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	order.CalculateFinalPrice()
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}
