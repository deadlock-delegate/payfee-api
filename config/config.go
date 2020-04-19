package config

type WebhookConfig struct {
	Token           string
	VerificationKey string
}

var Webhook WebhookConfig
