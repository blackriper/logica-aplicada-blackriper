package models

type RequestIntent struct {
	PriceId   string `form:"price_id" json:"price_id"  binding:"required"`
	ProductId string ` form:"product_id" json:"product_id"  binding:"required"`
}
