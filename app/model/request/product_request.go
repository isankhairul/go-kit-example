package request

import (
	"go-kit-example/helper/message"

	validation "github.com/itgelo/ozzo-validation/v4"
)

// swagger:model ProductListRequest
type ProductListRequest struct {
	// Maximun records per page
	// in: int32
	Limit int `schema:"limit" binding:"omitempty,numeric,min=1,max=100"`

	// Page No
	// in: int32
	Page int `schema:"page" binding:"omitempty,numeric,min=1"`

	// Sort fields, example: name asc, uom desc
	// in: string
	Sort string `schema:"sort" binding:"omitempty"`

	// Name keyword of the product
	// in: string
	Name string `schema:"name" binding:"omitempty"`

	// Sku of the product
	// in: string

	Sku string `schema:"sku" binding:"omitempty"`
	// Sku of the product
	// in: string
	UOM string `schema:"uom" binding:"omitempty"`
}

// swagger:parameters SaveProductRequest
type ReqProductBody struct {
	//  in: body
	Body SaveProductRequest `json:"body"`
}

type SaveProductRequest struct {
	// Name of the product
	// in: string
	Name string `json:"name"`

	// Sku of the product
	// in: string
	Sku string `json:"sku"`

	// Uom of the product
	// in: string
	Uom string `json:"uom"`

	// Weight of the product
	// in: int32
	Weight int32 `json:"weight"`

	// Uid of the product, use this on UPDATE function
	// in: int32
	Uid string `json:"uid" binding:"omitempty"`
}

func (req SaveProductRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required.Error(message.ErrReq.Message)),
		validation.Field(&req.Uom, validation.Required.Error(message.ErrReq.Message)),
	)
}
