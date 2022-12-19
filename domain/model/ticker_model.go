package model

import (
	"sort"
	"time"
)

type Ticker struct {
	Open   float32   `json:"open"`
	High   float32   `json:"high"`
	Low    float32   `json:"low"`
	Close  float32   `json:"close"`
	Volume uint      `json:"volume"`
	Time   time.Time `json:"create_at"`
}

func NewTicker(open, high, low, close float32, volume uint, time time.Time) (*Ticker, error) {

	return &Ticker{
		Open:   open,
		High:   high,
		Low:    low,
		Close:  close,
		Volume: volume,
		Time:   time,
	}, nil
}

func TickerTimeSort(tickerSlice []Ticker) []Ticker {

	sort.Slice(tickerSlice, func(i, j int) bool {
		return tickerSlice[i].Time.Before(tickerSlice[j].Time)
	})

	return tickerSlice
}
