package models

type Card struct {
	ID    string `json:"-"`
	Suit  string `json:"suit"`
	Value string `json:"value"`
	Code  string `json:"code"`
}
