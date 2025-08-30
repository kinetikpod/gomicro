package grpc

import (
	"context"

	"github.com/huseyinbabal/microservices-proto/golang/order"
	"github.com/kinetikpod/gomicro/order/internal/application/core/domain"
)

// Implementasi Create endpoint gRPC
func (a Adapter) Create(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem

	// convert dari gRPC request ke domain object
	for _, item := range req.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	// buat domain order
	newOrder := domain.NewOrder(req.UserId, orderItems)

	// panggil business logic lewat APIPort
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	// kembalikan response gRPC
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

