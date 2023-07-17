package schemas

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id          string `gorm:"primaryKey"`
	Description string
	Designation string
	Company     string
	Price       decimal.Decimal
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.Id = uuid.NewString()
	return
}

type ProductResponse struct {
	Id          string          `json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   time.Time       `json:"deletedAt,omitempty"`
	Description string          `json:"description"`
	Designation string          `json:"designation"`
	Company     string          `json:"company"`
	Price       decimal.Decimal `json:"price"`
}
