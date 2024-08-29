package transport

type PrompPayQRCodeRequest struct {
	Target string  `json:"target"`
	Amount float64 `json:"amount"`
}
