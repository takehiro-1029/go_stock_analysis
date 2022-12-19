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
		return nil, errors.New("time入力してください")
	}

	if !CheckIntervalTime(time) {
		return nil, errors.New("time値が正しくありません。")
	}

	interval := &Interval{
		Time: time,
	}

	return interval, nil
}

func CheckIntervalTime(time string) bool {

	check := []string{"1min", "5min", "15min", "30min", "60min"}

	for i := range check {
		if check[i] == time {
			return true
		}
	}

	return false
}
