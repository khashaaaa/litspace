package service

type Service struct {
	Mark     string  `json:"mark"`
	Provider string  `json:"provider"`
	Category uint    `json:"category"`
	Title    *string `json:"title"`
	Descr    *string `json:"descr"`
	CostFrom *uint   `json:"cost_from"`
	CostUp   *uint   `json:"cost_up"`
	Opened   *string `json:"opened"`
	Closed   *string `json:"closed"`
	Created  *string `json:"created"`
	Updated  *string `json:"updated"`
}
