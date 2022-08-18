package models

type Item struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Qty    int     `json:"qty"`
	Weight float32 `json:"weight"`
}
