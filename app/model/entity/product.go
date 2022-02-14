package entity

import (
	"go-kit-example/app/model/base"
)

// swagger:model product
type Product struct {
	base.BaseIDModel
	// Name of the product
	// in: string
	Name string `json:"name" faker:"first_name"`

	// Sku of the product
	// in: string
	Sku string `json:"sku" faker:"time_period"`

	// Uom of the product
	// in: string
	Uom string `json:"uom" faker:"uuid_digit"`

	// Weight of the product
	// in: int32
	Weight int32 `json:"weight" faker:"oneof: 15, 27, 61"`
}
