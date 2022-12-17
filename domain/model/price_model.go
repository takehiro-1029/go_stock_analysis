package model

import (
	"errors"
	"time"
)

type Price struct {
	ID              string
	StockID         string
	IntervalID      string
	Open            float32
	High            float32
	Low             float32
	Close           float32
	Volume          uint
	AcquisitionTime time.Time
}

func NewPrice(stockID, intervalID string, open, high, low, close float32, volume uint, acquisitiontime time.Time) (*Price, error) {
	if stockID == "" {
		return nil, errors.New("stockIDを入力してください")
	}
	if intervalID == "" {
		return nil, errors.New("intervalIDを入力してください")
	}

	price := &Price{
		StockID:         stockID,
		IntervalID:      intervalID,
		Open:            open,
		High:            high,
		Low:             low,
		Close:           close,
		Volume:          volume,
		AcquisitionTime: acquisitiontime,
	}

	return price, nil
}
