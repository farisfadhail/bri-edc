package validations

type AuthorizationRequest struct {
	MerchantID string  `json:"merchant_id"`
	TerminalID string  `json:"terminal_id"`
	Amount     float64 `json:"amount"`
	CardNumber string  `json:"card_number"`
	Timestamp  string  `json:"timestamp"`
}
