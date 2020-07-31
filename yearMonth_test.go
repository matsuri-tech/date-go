package mdate

import (
	"github.com/matsuri-tech/common-error-go"
	"reflect"
	"testing"
)

func TestYearMonth_NextMonth(t *testing.T) {

	tests := []struct {
		in   YearMonth
		want YearMonth
	}{
		{
			in: YearMonth{
				Year:  2020,
				Month: 1,
			},
			want: YearMonth{
				Year:  2020,
				Month: 2,
			},
		},
		{
			in: YearMonth{
				Year:  2019,
				Month: 12,
			},
			want: YearMonth{
				Year:  2020,
				Month: 1,
			},
		},
	}

	for _, tt := range tests {
		result := tt.in.NextMonth()
		if result != tt.want {
			t.Error(result, tt.want)
		}
	}
}

func TestYearMonthDiff(t *testing.T) {

	type in struct {
		okimochi string
		span     DateSpan
	}

	tests := []struct {
		in   in
		want YearMonths
	}{
		{
			in: in{
				"",
				MustDateSpan(NewDate(2020, 1, 1), NewDate(2020, 5, 25)),
			},
			want: YearMonths{
				YearMonth{
					Year:  2020,
					Month: 1,
				},
				YearMonth{
					Year:  2020,
					Month: 2,
				},
				YearMonth{
					Year:  2020,
					Month: 3,
				},
				YearMonth{
					Year:  2020,
					Month: 4,
				},
				YearMonth{
					Year:  2020,
					Month: 5,
				},
			},
		},
		{
			in: in{
				"",
				MustDateSpan(NewDate(2019, 11, 1), NewDate(2020, 3, 25)),
			},
			want: YearMonths{
				YearMonth{
					Year:  2019,
					Month: 11,
				},
				YearMonth{
					Year:  2019,
					Month: 12,
				},
				YearMonth{
					Year:  2020,
					Month: 1,
				},
				YearMonth{
					Year:  2020,
					Month: 2,
				},
				YearMonth{
					Year:  2020,
					Month: 3,
				},
			},
		},
	}

	for _, tt := range tests {
		result := YearMonthDiff(tt.in.span)
		if !reflect.DeepEqual(result, tt.want) {
			t.Error(result, tt.want)
		}
	}
}

func TestNewMonth(t *testing.T) {
	type want struct {
		error
		Month
	}

	tests := []struct {
		in   int
		want want
	}{
		{
			in: 8,
			want: want{
				nil,
				August,
			},
		},
		{
			in: 0,
			want: want{
				InvalidMonth(),
				0,
			},
		},
		{
			in: 13,
			want: want{
				InvalidMonth(),
				0,
			},
		},
	}

	for _, tt := range tests {
		result, err := NewMonth(tt.in)
		if err != nil {
			if err.(merrors.CommonError).ErrorType != tt.want.error.(merrors.CommonError).ErrorType {
				t.Error(result, tt.want)
			}
		} else {
			if err != tt.want.error {
				t.Error(result, tt.want)
			}
		}
		if result != tt.want.Month {
			t.Error(result, tt.want.Month)
		}
	}
}

func TestYearMonth_IsAfter(t *testing.T) {
	type in struct {
		ym1 YearMonth
		ym2 YearMonth
	}

	tests := []struct {
		in   in
		want bool
	}{
		{
			in: in{
				ym1: YearMonth{
					Year:  2020,
					Month: 1,
				},
				ym2: YearMonth{
					Year:  2020,
					Month: 2,
				}},
			want: false,
		},
		{
			in: in{
				ym1: YearMonth{
					Year:  2019,
					Month: 12,
				},
				ym2: YearMonth{
					Year:  2020,
					Month: 2,
				}},
			want: false,
		},
		{
			in: in{
				ym1: YearMonth{
					Year:  2020,
					Month: 12,
				},
				ym2: YearMonth{
					Year:  2019,
					Month: 2,
				}},
			want: true,
		},
		{
			in: in{
				ym1: YearMonth{
					Year:  2019,
					Month: 12,
				},
				ym2: YearMonth{
					Year:  2019,
					Month: 12,
				}},
			want: false,
		},
	}

	for _, tt := range tests {
		result := tt.in.ym1.IsAfter(tt.in.ym2)
		if result != tt.want {
			t.Error(result, tt.want)
		}
	}
}
