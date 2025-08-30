package ports

import "github.com/kinetikpod/gomicro/order/internal/application/core/domain"

// APIPort → kontrak untuk fungsionalitas aplikasi
type APIPort interface {
    PlaceOrder(order domain.Order) (domain.Order, error)
}

