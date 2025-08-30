package api

import (
    "github.com/kinetikpod/gomicro/order/internal/application/core/domain"
    "github.com/kinetikpod/gomicro/order/internal/ports"
)

// Application struct yang bergantung pada DBPort
type Application struct {
    db ports.DBPort
}

// NewApplication → constructor untuk Application
func NewApplication(db ports.DBPort) *Application {
    return &Application{
        db: db,
    }
}

// PlaceOrder → simpan order via DBPort
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
    err := a.db.Save(&order)
    if err != nil {
        return domain.Order{}, err
    }
    return order, nil
}

