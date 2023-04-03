package models

type Instrument struct {
	ID          int     `valid:"required"`
	Name        string  `valid:"required"`
	Brand       string  `valid:"required"`
	Model       string  `valid:"required"`
	Type        string  `valid:"required"`
	Price       float64 `valid:"required,numeric,min=0"`
	Color       string
	Description string `valid:"required"`
	Qty_Stock   int    `valid:"required,numeric,min=0"`
}
