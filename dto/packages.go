package dto

type PackageRequest struct {
	Name      string      `json:"name" validate:"required"`
	Price     int         `json:"price" validate:"required"`
	List      []List      `json:"list"`
	Aditional []Aditional `json:"aditional"`
}

type List struct {
	Name string `json:"name"`
}

type Aditional struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type CreateListRequest struct {
	List []List `json:"list"`
}

type CreateAditionalRequest struct {
	Aditional []Aditional `json:"aditional"`
}
