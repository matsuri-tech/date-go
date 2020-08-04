package mdate

import (
	"github.com/matsuri-tech/common-error-go"
	"time"
)

type Month int

type YearMonth struct {
	Year  Year
	Month Month
}

type YearMonths []YearMonth

const (
	January   Month = 1
	February  Month = 2
	March     Month = 3
	April     Month = 4
	May       Month = 5
	June      Month = 6
	July      Month = 7
	August    Month = 8
	September Month = 9
	October   Month = 10
	November  Month = 11
	December  Month = 12
)

const (
	ErrorInvalidMonth merrors.ErrorType = "invalidMonth"
)

func InvalidMonth() merrors.CommonError {
	return merrors.ErrorBadRequest("", ErrorInvalidMonth)
}

func NewMonth(m int) (Month, error) {
	if m < int(January) || m > int(December) {
		return Month(0), InvalidMonth()
	}
	return Month(m), nil
}

func NewYearMonth(year Year, month Month) YearMonth {
	return YearMonth{
		Year:  year,
		Month: month,
	}
}

func (ym YearMonth) NextMonth() YearMonth {
	if ym.Month == December {
		return YearMonth{
			Year:  ym.Year + 1,
			Month: January,
		}
	}
	return YearMonth{
		Year:  ym.Year,
		Month: ym.Month + 1,
	}
}

func (ym YearMonth) IsAfter(another YearMonth) bool {
	if ym.Year > another.Year {
		return true
	}
	if ym.Year < another.Year {
		return false
	}
	return ym.Month > another.Month
}

func (ym YearMonth) StartDate() Date {
	return NewDate(int(ym.Year), time.Month(ym.Month), 1)
}

func (ym YearMonth) EndDate() Date {
	nextYm := ym.NextMonth()
	return nextYm.StartDate().MinusNDay(1)
}