package validations

type CreateTransactionRequest struct {
	MerchantID string `json:"merchant_id" validate:"required"`
	TerminalID string `json:"terminal_id" validate:"required"`
	Amount     int    `json:"amount" validate:"required,gt=0"`
	CardNumber string `json:"card_number" validate:"required"`
	Timestamp  string `json:"timestamp" validate:"required"`
}
