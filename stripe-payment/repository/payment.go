package repository

import "github.com/blackriper/payment/models"

type Payment interface {
	GetCatalogProducts() []models.ProductDto
	NewPaymentIntent(priceId, prodId string) (models.ResponseIntent, error)
	GetDataPayment(intentId string) string
}
