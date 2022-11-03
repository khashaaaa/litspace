package sorder

type Sorder struct {
	Mark          string  `json:"mark"`
	Service       *string `json:"service"`
	Amount        *uint   `json:"amount"`
	Demandant     *string `json:"demandant"`
	Executor      *string `json:"executor"`
	DemandantType *string `json:"demandant_type"`
	ExecutorType  *string `json:"executor_type"`
	Issued        *string `json:"issued,omitempty"`
}
