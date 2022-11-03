package pcategory

type PCategory struct {
	Mark    uint    `json:"mark"`
	Name    *string `json:"name"`
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
}
