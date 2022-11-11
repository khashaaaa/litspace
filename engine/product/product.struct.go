package product

type Product struct {
	Mark       string      `json:"mark"`
	Merchant   string      `json:"merchant"`
	Category   uint        `json:"category"`
	Name       *string     `json:"name"`
	Descr      *string     `json:"descr"`
	Price      *uint       `json:"price"`
	Stock      *uint       `json:"stock"`
	Attrs      *Attributes `json:"attrs"`
	ImagePaths *[]string   `json:"image_paths"`
	Created    *string     `json:"created"`
	Updated    *string     `json:"updated"`
}
