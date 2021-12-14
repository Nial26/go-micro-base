package order

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	LineItems []LineItem
	Price     float64
}

type LineItem struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	ItemId    string
	Count     int
}

type Entity struct {
	Id         string `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	OrderItems []ItemEntity
	Price      float64
}

type ItemEntity struct {
	Id        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ItemId    string
	Count     string
}
