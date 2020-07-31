package mdate

import "github.com/matsuri-tech/common-error-go"

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

func YearMonthDiff(start Date, end Date) YearMonths {
	var result YearMonths
	var currentYearMonth = NewYearMonth(start.Year(), start.Month())
	for {
		result = append(result, currentYearMonth)
		if currentYearMonth == end.YearMonth() {
			break
		}
		currentYearMonth = currentYearMonth.NextMonth()
	}
	return result
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
