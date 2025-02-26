package repository

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/blackriper/payment/models"
	"github.com/blackriper/payment/utils"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/v81/price"
	"github.com/stripe/stripe-go/v81/product"
	"go.uber.org/zap"
)

type StripeRepository struct {
	Logger *zap.Logger
}

func (s StripeRepository) GetCatalogProducts() []models.ProductDto {
	var products []models.ProductDto

	params := &stripe.ProductListParams{}
	params.Limit = stripe.Int64(7)
	result := product.List(params)

	for result.Next() {
		p := result.Product()
		formattLabel := s.formattedPrice(p.DefaultPrice.ID)
		pr := models.ProductDto{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			ImageUrl:    p.Images[0],
			Price:       formattLabel,
			PriceId:     p.DefaultPrice.ID,
		}
		products = append(products, pr)

	}
	if err := result.Err(); err != nil {
		s.Logger.Error(fmt.Sprintf("error to load product error:%v", err))
	}
	return products
}

func (s StripeRepository) NewPaymentIntent(priceId, prodId string) (models.ResponseIntent, error) {
	productPrice, err := price.Get(priceId, &stripe.PriceParams{})
	if err != nil {
		s.Logger.Error(fmt.Sprintf("error to find price %v", err))
		return models.ResponseIntent{}, err
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(productPrice.UnitAmount),
		Currency: stripe.String(string(productPrice.Currency)),
		Metadata: map[string]string{
			"product_id": prodId,
			"price_id":   priceId,
		},
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	payment, err := paymentintent.New(params)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("error to create payment intent %v ", err))
		return models.ResponseIntent{}, err
	}

	result := models.ResponseIntent{
		ClientSecret: payment.ClientSecret,
		StripeKey:    os.Getenv("STRIPE_PUBLIC_KEY"),
		Amount:       s.formattedPrice(priceId),
	}

	return result, nil
}

func (s StripeRepository) GetDataPayment(intentId string) string {
	result, err := paymentintent.Get(intentId, &stripe.PaymentIntentParams{})
	if err != nil {
		s.Logger.Error(fmt.Sprintf("error to find payment %v", err))
	}
	return s.formattedPrice(result.Metadata["price_id"])
}

// auxilar function  for formatting price
func (s StripeRepository) formattedPrice(priceId string) string {
	priceParams := &stripe.PriceParams{}

	priceProduct, err := price.Get(priceId, priceParams)
	if err != nil {
		slog.Error("error to find price data ", "%v", err)
	}

	return fmt.Sprintf(
		"%s %.2f %s",
		utils.GetCurrencySymbol(string(priceProduct.Currency)),
		float64(priceProduct.UnitAmount)/100,
		strings.ToUpper(string(priceProduct.Currency)),
	)
}
