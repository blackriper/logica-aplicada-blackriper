package routes

import (
	"github.com/blackriper/payment/controllers"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router     *gin.Engine
	controller *controllers.ControllerPayment
}

func NewRoutes(router *gin.Engine, controller *controllers.ControllerPayment) *Routes {
	return &Routes{
		router:     router,
		controller: controller,
	}
}

func (r *Routes) CreateRoutes() {
	r.router.GET("/", r.controller.HomeTemplate)
	r.router.GET("/products", r.controller.ProductsTemplate)
	r.router.GET("/checkout/:priceId", r.controller.CheckoutTemplate)
	r.router.GET("/success", r.controller.SuccessTemplate)
	r.router.POST("/create-payment-intent", r.controller.NewPaymentIntent)
}
