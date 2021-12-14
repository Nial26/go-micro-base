package item

import (
	"time"

	"github.com/ulule/deepcopier"
	"gorm.io/gorm"
)

type Item struct {
	Id                string  `json:"id"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	AvailableQuantity int     `json:"available_quantity"`
}

type Entity struct {
	Id                string `gorm:"primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	Name              string
	Description       string
	Price             float64
	AvailableQuantity int `gorm:"column:stock_count"`
}

func (e Entity) TableName() string {
	return "items"
}

func (i Item) ToEntity() Entity {
	var e Entity
	deepcopier.Copy(i).To(&e)
	return e
}

func (e Entity) ToDTO() Item {
	var i Item
	deepcopier.Copy(e).To(&i)
	i.CreatedAt = e.CreatedAt.String()
	i.UpdatedAt = e.UpdatedAt.String()
	return i
}
