package controllers

import (
	"net/http"

	"github.com/blackriper/payment/models"
	"github.com/blackriper/payment/repository"
	"github.com/gin-gonic/gin"
)

type ControllerPayment struct {
	repoPayment repository.Payment
}

func NewControllerPayment(repository repository.Payment) *ControllerPayment {
	return &ControllerPayment{
		repoPayment: repository,
	}
}

func (con ControllerPayment) HomeTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, models.HOME, nil)
}

func (con ControllerPayment) ProductsTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, models.PRODUCTS, gin.H{
		"products": con.repoPayment.GetCatalogProducts(),
	})
}

func (con ControllerPayment) CheckoutTemplate(c *gin.Context) {
	priceId := c.Param("priceId")
	productId := c.Query("prodId")

	c.HTML(http.StatusOK, models.CHECKOUT, gin.H{
		"PriceId": priceId,
		"ProdId":  productId,
	})
}

func (con ControllerPayment) SuccessTemplate(c *gin.Context) {
	paymentId := c.Query("payment_intent")
	amount := con.repoPayment.GetDataPayment(paymentId)
	c.HTML(http.StatusOK, models.SUCCESS, gin.H{
		"PaymentId": paymentId,
		"Amount":    amount,
	})
}

func (con ControllerPayment) NewPaymentIntent(c *gin.Context) {
	var request models.RequestIntent
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reponseIntent, err := con.repoPayment.NewPaymentIntent(request.PriceId, request.ProductId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"body": reponseIntent})
}
