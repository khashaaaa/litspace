package ssubcategory

type SSubCategory struct {
	Mark     uint    `json:"mark"`
	Name     string  `json:"name"`
	Category uint    `json:"category"`
	Created  *string `json:"created,omitempty"`
	Updated  *string `json:"updated,omitempty"`
}
