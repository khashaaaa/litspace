package provider

type Provider struct {
	Mark          string  `json:"mark"`
	Founder       *string `json:"founder"`
	EntityName    *string `json:"entity_name"`
	Email         *string `json:"email"`
	Mobile        *string `json:"mobile"`
	Address       *string `json:"address"`
	OriginCountry *string `json:"origin_country"`
	InStatus      string  `json:"in_status"`
	Type          string  `json:"type"`
	Created       *string `json:"created"`
	Updated       *string `json:"updated"`
}
