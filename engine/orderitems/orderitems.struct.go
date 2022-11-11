package orderitems

type OrderItems struct {
	Mark      uint    `json:"mark"`
	PMark     *uint   `json:"p_mark"`
	PName     *string `json:"p_name"`
	PQuantity *uint   `json:"p_quantity"`
	PPrice    *uint   `json:"p_price"`
	Seller    *string `json:"seller"`
	Buyer     *string `json:"buyer"`
}
