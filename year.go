package mdate

// 日本の年度を表すデータ型
type JapanFiscalYear int

type Year int

func NewYear(y int) Year {
	return Year(y)
}
