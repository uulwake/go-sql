package models

type Order struct {
	ID               int    `json:"id"`
	RecipientName    string `json:"recipient_name"`
	RecipientAddress string `json:"recipient_address"`
	Shipper          string `json:"shipper"`
}
