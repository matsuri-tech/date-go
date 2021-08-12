package mdate

import (
	"fmt"
	merrors "github.com/matsuri-tech/common-error-go"
	"strconv"
	"time"
)

type Month int

type YearMonth struct {
	Year  Year  `json:"year"`
	Month Month `json:"month"`
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

func (month Month) String() string {
	monthStr := strconv.Itoa(int(month))
	if month < 10 {
		return fmt.Sprintf("0%v", monthStr)
	}
	return monthStr
}

func (ym YearMonth) String() string {
	return fmt.Sprintf("%v-%v", ym.Year, ym.Month.String())
}

const yearMonthStrFormat = "2006-01"

func NewYearMonthFromStr(str string) (YearMonth, error) {
	t, err := time.Parse(yearMonthStrFormat, str)
	if err != nil {
		return YearMonth{}, err
	}

	y := NewYear(t.Year())
	m := Month(t.Month())

	return NewYearMonth(y, m), nil
}
