package schemas

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID
	Description string
	Designation string
	Company     string
	Price       decimal.Decimal
}
