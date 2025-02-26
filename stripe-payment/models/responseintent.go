package models

type ResponseIntent struct {
	StripeKey    string `json:"stripe_key"`
	ClientSecret string `json:"client_secret"`
	Amount       string `json:"amount"`
}
