package handler

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateProductRequest struct {
	Description string          `json:"description"`
	Designation string          `json:"designation"`
	Company     string          `json:"company"`
	Price       decimal.Decimal `json:"price"`
}

func (r *CreateProductRequest) Validate() error {
	if r.Description == "" && r.Designation == "" && r.Company == "" && r.Price.LessThanOrEqual(decimal.NewFromInt(0)) {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Designation == "" {
		return errParamIsRequired("designation", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Price.LessThanOrEqual(decimal.NewFromInt(0)) {
		return errParamIsRequired("price", "decimal")
	}
	return nil
}
