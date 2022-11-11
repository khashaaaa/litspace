package porder

type POrder struct {
	Mark     string  `json:"mark"`
	Items    *uint   `json:"items"`
	Sum      *uint   `json:"sum"`
	Amount   *uint   `json:"amount"`
	InStatus *string `json:"in_status"`
	Issued   *string `json:"issued,omitempty"`
}
