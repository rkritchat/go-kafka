package model

type RegisterRq struct {
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
}
