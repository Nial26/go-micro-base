package order

import (
	"fmt"

	"github.com/segmentio/ksuid"
	"go-micro-base/item"
)

type Service interface {
	CreateOrder(items []item.Item) (Order, error)
	GetOrder(orderId string) Order
}

type service struct {
	orderRepository Repository
	itemService     item.Service
}

func (s service) CreateOrder(items []LineItem) (Order, error) {
	totalOrderPrice := 0.0
	for _, oi := range items {
		// Verify the item is still present
		itm, err := s.itemService.FindItem(oi.ItemId)

		if err != nil {
			return Order{}, err
		}

		if itm.AvailableQuantity-oi.Count < 0 {
			return Order{}, fmt.Errorf("can't buy more items than available stock")
		}

		// Add up the total price
		totalOrderPrice += itm.Price * float64(oi.Count)
	}

	_ = Order{
		Id:        ksuid.New().String(),
		LineItems: items,
		Price:     totalOrderPrice,
	}
	// s.orderRepository.Create()
	return Order{}, nil
}

func (s service) GetOrder(orderId string) Order {
	panic("implement me")
}
