package provider

type Provider struct {
	Mark          string  `json:"mark"`
	Founder       *string `json:"founder"`
	EntityName    *string `json:"entity_name"`
	Email         *string `json:"email"`
	Mobile        *string `json:"mobile"`
	Address       *string `json:"address"`
	OriginCountry *string `json:"origin_country,omitempty"`
	InStatus      *string `json:"in_status"`
	Type          *string `json:"type,omitempty"`
	Created       *string `json:"created,omitempty"`
	Updated       *string `json:"updated,omitempty"`
}
