package schemas

import (
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
