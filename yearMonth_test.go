package mdate

import (
	"github.com/matsuri-tech/common-error-go"
	"github.com/stretchr/testify/assert"
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

func TestYearMonth_StartDate(t *testing.T) {
	tests := []struct {
		in   YearMonth
		want Date
	}{
		{
			in: YearMonth{
				Year:  2020,
				Month: 1,
			},
			want: NewDate(2020, 1, 1),
		},
		{
			in: YearMonth{
				Year:  2019,
				Month: 12,
			},
			want: NewDate(2019, 12, 1),
		},
		{
			in: YearMonth{
				Year:  2020,
				Month: 2,
			},
			want: NewDate(2020, 2, 1),
		},
	}

	for _, tt := range tests {
		result := tt.in.StartDate()
		if result != tt.want {
			t.Error(result, tt.want)
		}
	}
}

func TestYearMonth_EndDate(t *testing.T) {
	tests := []struct {
		in   YearMonth
		want Date
	}{
		{
			in: YearMonth{
				Year:  2020,
				Month: 1,
			},
			want: NewDate(2020, 1, 31),
		},
		{
			in: YearMonth{
				Year:  2019,
				Month: 12,
			},
			want: NewDate(2019, 12, 31),
		},
		{
			in: YearMonth{
				Year:  2020,
				Month: 2,
			},
			want: NewDate(2020, 2, 29),
		},
		{
			in: YearMonth{
				Year:  2019,
				Month: 2,
			},
			want: NewDate(2019, 2, 28),
		},
		{
			in: YearMonth{
				Year:  2019,
				Month: 6,
			},
			want: NewDate(2019, 6, 30),
		},
	}

	for _, tt := range tests {
		result := tt.in.EndDate()
		if result != tt.want {
			t.Error(result, tt.want)
		}
	}
}

func TestMonth_String(t *testing.T) {

	tests := []struct {
		in   YearMonth
		want string
	}{
		{
			in: YearMonth{
				Year:  2021,
				Month: 2,
			},
			want: "2021-02",
		},
		{
			in: YearMonth{
				Year:  2021,
				Month: 12,
			},
			want: "2021-12",
		},
	}

	for _, tt := range tests {
		result := tt.in.String()
		if result != tt.want {
			t.Error(result, tt.want)
		}
	}
}

func TestYearMonth_NewFromString(t *testing.T) {
	tests := []struct {
		in      string
		want    YearMonth
		wantErr bool
	}{
		{
			in:      "2021-01",
			want:    YearMonth{2021, 1},
			wantErr: false,
		},
		{
			in:      "1980-12",
			want:    YearMonth{1980, 12},
			wantErr: false,
		},
		{
			in:      "2000-24",
			want:    YearMonth{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		r, err := NewYearMonthFromStr(tt.in)

		assert.Equal(t, tt.wantErr, err != nil, "%v, %v", tt.in, err)
		assert.Equal(t, tt.want, r, "%v", tt.in)
	}
}
