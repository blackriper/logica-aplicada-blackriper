package main

import (
	"log/slog"
	"os"

	"github.com/blackriper/payment/controllers"
	"github.com/blackriper/payment/repository"
	"github.com/blackriper/payment/routes"
	"github.com/blackriper/payment/utils"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v81"
)

var (
	router        *gin.Engine
	stripeRepo    repository.Payment
	controllerPay *controllers.ControllerPayment
)

func init() {
	// load env  file and load stripe key
	utils.LoadEnv()
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// create router and render templates and static files
	router = utils.InicializeRouter()
	// load business logis
	logger := utils.ZipLogger("stripe_repository")
	stripeRepo = repository.StripeRepository{Logger: logger}
	controllerPay = controllers.NewControllerPayment(stripeRepo)
}

func main() {
	// load routes
	routesPay := routes.NewRoutes(router, controllerPay)
	// create routes and start server
	routesPay.CreateRoutes()

	if err := router.Run(":3000"); err != nil {
		slog.Error("error to connect to server ", "%v", err)
	}
}
