package review

type Review struct {
	Mark       uint    `json:"mark"`
	Subjective *string `json:"subjective"`
	Objective  *string `json:"objective"`
	Comment    *string `json:"comment"`
	Rate1      *uint   `json:"rate_1"`
	Rate2      *uint   `json:"rate_2"`
	Rate3      *uint   `json:"rate_3"`
	Rate4      *uint   `json:"rate_4"`
	Rate5      *uint   `json:"rate_5"`
	Rate6      *uint   `json:"rate_6"`
	Rate7      *uint   `json:"rate_7"`
	Rate8      *uint   `json:"rate_8"`
	Rate9      *uint   `json:"rate_9"`
	Rate10     *uint   `json:"rate_10"`
	Created    *string `json:"created"`
}
