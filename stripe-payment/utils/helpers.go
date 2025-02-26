package utils

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// load env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Error("error to load .env file", "error:", err)
	}
}

// get currency symbols
func GetCurrencySymbol(currency string) string {
	switch strings.ToUpper(currency) {
	case "USD":
		return "$"
	case "EUR":
		return "€"
	case "GBP":
		return "£"
	case "JPY":
		return "¥"
	default:
		return "$"
	}
}

func ZipLogger(appName string) *zap.Logger {
	logger, _ := zap.NewDevelopment()
	logger.With(
		zap.String("name", appName),
	)
	return logger
}

// configure static files , templates and logger
func InicializeRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")
	router.Static("/static", "static")
	return router
}
