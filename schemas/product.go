package schemas

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Description string
	Designation string
	Company     string
	Price       decimal.Decimal
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.NewString()
	return
}

type ProductResponse struct {
	ID          string          `json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   time.Time       `json:"deletedAt,omitempty"`
	Description string          `json:"description"`
	Designation string          `json:"designation"`
	Company     string          `json:"company"`
	Price       decimal.Decimal `json:"price"`
}
