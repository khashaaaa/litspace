package product

type Attributes struct {
	Color   *string  `json:"color"`
	Width   *float32 `json:"width"`
	Height  *float32 `json:"height"`
	Weight  *float32 `json:"weight"`
	Size    *string  `json:"size"`
	Feature *string  `json:"feature"`
}
