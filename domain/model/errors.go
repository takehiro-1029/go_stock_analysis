package model

type StockExistsError struct{}

func (e *StockExistsError) Error() string {
	return "存在する銘柄です。"
}

type IntervalCheckError struct{}

func (e *IntervalCheckError) Error() string {
	return "intervalが不正値です。再度ご確認ください。"
}
