package ports

import "github.com/kinetikpod/gomicro/order/internal/application/core/domain"

// DBPort → kontrak untuk database
type DBPort interface {
    Get(id string) (domain.Order, error)
    Save(order *domain.Order) error
}

