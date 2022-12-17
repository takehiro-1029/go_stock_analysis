package model

import (
	"errors"
)

type Interval struct {
	ID   string
	Time string
}

func NewInterval(time string) (*Interval, error) {
	if time == "" {
		return nil, errors.New("timeを入力してください")
	}

	// TODO
	// time値の確認

	interval := &Interval{
		Time: time,
	}

	return interval, nil
}
