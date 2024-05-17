package entity

import (
	"encoding/json"
	"errors"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ErrNotFound error

func (p *Product) UnmarshalJSON(data []byte) error {
	type Alias Product
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if p.ID != 0 {
		return errors.New("product ID must not be provided")
	}
	if p.Name == "" {
		return errors.New("product name is required")
	}
	if p.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	if p.Quantity < 0 {
		return errors.New("product quantity must be zero or greater")
	}

	return nil
}
