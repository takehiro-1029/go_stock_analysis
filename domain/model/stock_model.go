package model

import (
	"errors"
)

type Stock struct {
	ID     string
	Symbol string
	Name   string
}

func NewStock(symbol, name string) (*Stock, error) {
	if symbol == "" || name == "" {
		return nil, errors.New("titleを入力してください")
	}

	stock := &Stock{
		Symbol: symbol,
		Name:   name,
	}

	return stock, nil
}
