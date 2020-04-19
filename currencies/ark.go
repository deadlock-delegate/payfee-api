package ark

type Transaction struct {
	SenderPublicKey string
	RecipientID     string
	Amount          string
	TransactionFee  string
	Vendorfield     string
	Type            int
}

type WebhookData struct {
	Height             string
	Timestamp          string
	GeneratorPublicKey string
	Transactions       []Transaction
}
