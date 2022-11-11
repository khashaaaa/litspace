package paymethod

type PayMethod struct {
	Mark         string `json:"mark"`
	Owner        string `json:"owner"`
	HolderName   string `json:"holder_name"`
	CardNumber   string `json:"card_number"`
	Expiry       string `json:"expiry"`
	Cvv          string `json:"cvv"`
	Region       string `json:"region"`
	CardProvider string `json:"card_provider"`
}
