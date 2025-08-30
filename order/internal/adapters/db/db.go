package db

import (
	"fmt"
	"time"

	"github.com/kinetikpod/gomicro/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ---------------------------
// GORM Models (untuk DB layer)
// ---------------------------

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

// ---------------------------
// Adapter struct
// ---------------------------

type Adapter struct {
	db *gorm.DB
}

// NewAdapter → buat koneksi ke MySQL
func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	// AutoMigrate bikin tabel berdasarkan struct
	err := db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

// ---------------------------
// Implementasi DBPort
// ---------------------------

// Get → ambil Order by ID
func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.Preload("OrderItems").First(&orderEntity, id) // include OrderItems

	var orderItems []domain.OrderItem
	for _, item := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.Unix(), // pakai detik biar konsisten
	}

	return order, res.Error
}

// Save → simpan Order
func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
		order.CreatedAt = time.Now().Unix()
	}
	return res.Error
}
