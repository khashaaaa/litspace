package service

type Service struct {
	Mark     string  `json:"mark"`
	Provider *string `json:"provider"`
	Category *uint   `json:"category"`
	Title    *string `json:"title"`
	Descr    *string `json:"descr"`
	CostFrom *uint   `json:"cost_from"`
	CostUp   *uint   `json:"cost_up"`
	Opened   *string `json:"opened,omitempty"`
	Closed   *string `json:"closed,omitempty"`
	Created  *string `json:"created,omitempty"`
	Updated  *string `json:"updated,omitempty"`
}
