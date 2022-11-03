package consumer

type Consumer struct {
	Mark          string  `json:"mark"`
	FirstName     *string `json:"first_name"`
	LastName      *string `json:"last_name"`
	Email         *string `json:"email"`
	Mobile        *string `json:"mobile"`
	OriginCountry *string `json:"origin_country"`
	Pass          string  `json:"pass"`
	Type          string  `json:"type"`
	Created       *string `json:"created"`
	Updated       *string `json:"updated"`
}
