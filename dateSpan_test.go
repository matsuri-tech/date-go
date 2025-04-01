package mdate

import (
	"fmt"
	"reflect"
	"testing"

	merrors "github.com/matsuri-tech/common-error-go"
)

func TestDateSpan_IsContinuous(t *testing.T) {
	tests := []struct {
		in  DateSpans
		out bool
	}{
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 13)),
				MustDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15)),
			},
			out: true,
		},
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 12)),
				MustDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15)),
			},
			out: false,
		},
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15)),
				MustDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 13)),
			},
			out: true,
		},
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 10), NewDate(2018, 11, 13)),
				MustDateSpan(NewDate(2018, 11, 9), NewDate(2018, 11, 15)),
			},
			out: true,
		},
	}

	for i, tt := range tests {
		if tt.in[0].IsContinuous(tt.in[1]) != tt.out {
			t.Error(i, tt)
		}
	}
}

func TestDateSpan_IsOverlapping(t *testing.T) {
	tests := []struct {
		in  DateSpans
		out bool
	}{
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 13)),
				MustDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15)),
			},
			out: false,
		},
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 12)),
				MustDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15)),
			},
			out: false,
		},
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 10), NewDate(2018, 11, 13)),
				MustDateSpan(NewDate(2018, 11, 9), NewDate(2018, 11, 15)),
			},
			out: true,
		},
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 10), NewDate(2018, 11, 13)),
				MustDateSpan(NewDate(2018, 11, 11), NewDate(2018, 11, 15)),
			},
			out: true,
		},
		{
			in: DateSpans{
				MustDateSpan(NewDate(2018, 11, 10), NewDate(2018, 11, 13)),
				MustDateSpan(NewDate(2018, 11, 9), NewDate(2018, 11, 11)),
			},
			out: true,
		},
	}
	for i, tt := range tests {
		if tt.in[0].IsOverlapping(tt.in[1]) != tt.out {
			t.Error(i, tt)
		}
	}
}

func TestDateSpan_IncludesDate(t *testing.T) {
	span := MustDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 15))

	tests := []struct {
		in  Date
		out bool
	}{
		{
			in:  NewDate(2018, 11, 13),
			out: true,
		},
		{
			in:  NewDate(2018, 11, 11),
			out: false,
		},
		{
			in:  NewDate(2018, 11, 16),
			out: false,
		},
	}

	for i, tt := range tests {
		if span.IncludesDate(tt.in) != tt.out {
			t.Error(i, tt)
		}
	}
}

func TestMerge(t *testing.T) {
	s1, _ := NewDateSpan(NewDate(2018, 11, 10), NewDate(2018, 11, 13))
	s2, _ := NewDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15))
	spans := DateSpans{s1, s2}.Merge()
	if spans.Len() != 1 {
		t.Error("")
	}
	if !spans[0].EndDate.IsEqual(NewDate(2018, 11, 15)) {
		fmt.Printf("%s", spans)
		t.Errorf("%s", spans[0].EndDate)
	}
}

func TestDateSpan_GetDateList(t *testing.T) {
	tests := []struct {
		in  DateSpan
		out Dates
	}{
		{
			in: MustDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 13)),
			out: Dates{
				NewDate(2018, 11, 12),
				NewDate(2018, 11, 13),
			},
		},
	}
	for i, tt := range tests {
		result := tt.in.GetDateList()
		if !reflect.DeepEqual(result, tt.out) {
			t.Error(i, result, tt.out)
		}
	}
}

func TestDateSpan_OverlappingYearMonth(t *testing.T) {
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
		{
			in: in{
				"",
				MustDateSpan(NewDate(2019, 11, 1), NewDate(2019, 11, 1)),
			},
			want: YearMonths{
				YearMonth{
					Year:  2019,
					Month: 11,
				},
			},
		},
	}

	for _, tt := range tests {
		result := tt.in.span.OverlappingYearMonth()
		if !reflect.DeepEqual(result, tt.want) {
			t.Error(result, tt.want)
		}
	}
}

func TestClampDateSpan(t *testing.T) {
	type DateSpanList struct {
		a DateSpan
		b DateSpan
	}
	tests := []struct {
		caseName    string
		in          DateSpanList
		want        DateSpan
		expectedErr error
	}{
		{
			caseName: "simple",
			in: DateSpanList{
				a: MustDateSpan(NewDate(2021, 3, 10), NewDate(2021, 4, 20)),
				b: MustDateSpan(NewDate(2021, 4, 1), NewDate(2022, 3, 31)),
			},
			want: DateSpan{
				StartDate: NewDate(2021, 4, 1),
				EndDate:   NewDate(2021, 4, 20),
			},
			expectedErr: nil,
		},
		{
			caseName: "reverse of simple",
			in: DateSpanList{
				a: MustDateSpan(NewDate(2021, 4, 1), NewDate(2022, 3, 31)),
				b: MustDateSpan(NewDate(2021, 3, 10), NewDate(2021, 4, 20)),
			},
			want: DateSpan{
				StartDate: NewDate(2021, 4, 1),
				EndDate:   NewDate(2021, 4, 20),
			},
			expectedErr: nil,
		},
		{
			caseName: "no overlapped before",
			in: DateSpanList{
				a: MustDateSpan(NewDate(2021, 3, 10), NewDate(2021, 4, 20)),
				b: MustDateSpan(NewDate(2021, 5, 1), NewDate(2022, 3, 31)),
			},
			want:        DateSpan{},
			expectedErr: NoOverlapToClamp(),
		},
		{
			caseName: "no overlapped after",
			in: DateSpanList{
				a: MustDateSpan(NewDate(2021, 5, 1), NewDate(2022, 3, 31)),
				b: MustDateSpan(NewDate(2021, 3, 10), NewDate(2021, 4, 20)),
			},
			want:        DateSpan{},
			expectedErr: NoOverlapToClamp(),
		},
		{
			caseName: "matched perfectly",
			in: DateSpanList{
				a: MustDateSpan(NewDate(2021, 4, 1), NewDate(2022, 3, 31)),
				b: MustDateSpan(NewDate(2021, 4, 1), NewDate(2022, 3, 31)),
			},
			want: DateSpan{
				StartDate: NewDate(2021, 4, 1),
				EndDate:   NewDate(2022, 3, 31),
			},
			expectedErr: nil,
		},
		{
			caseName: "only last day match",
			in: DateSpanList{
				a: MustDateSpan(NewDate(2021, 4, 1), NewDate(2022, 3, 31)),
				b: MustDateSpan(NewDate(2022, 3, 31), NewDate(2022, 4, 1)),
			},
			want: DateSpan{
				StartDate: NewDate(2022, 3, 31),
				EndDate:   NewDate(2022, 3, 31),
			},
			expectedErr: nil,
		},
		{
			caseName: "included all",
			in: DateSpanList{
				a: MustDateSpan(NewDate(2021, 4, 1), NewDate(2022, 3, 31)),
				b: MustDateSpan(NewDate(2021, 3, 1), NewDate(2022, 4, 30)),
			},
			want: DateSpan{
				StartDate: NewDate(2021, 4, 1),
				EndDate:   NewDate(2022, 3, 31),
			},
			expectedErr: nil,
		},
		{
			caseName: "includes all",
			in: DateSpanList{
				a: MustDateSpan(NewDate(2021, 3, 1), NewDate(2022, 4, 30)),
				b: MustDateSpan(NewDate(2021, 4, 1), NewDate(2022, 3, 31)),
			},
			want: DateSpan{
				StartDate: NewDate(2021, 4, 1),
				EndDate:   NewDate(2022, 3, 31),
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		result, err := tt.in.a.ClampDateSpan(tt.in.b)
		if err != nil {
			if tt.expectedErr == nil || !merrors.ErrorTypeEqual(tt.expectedErr, err) {
				t.Errorf("Error occurs in case %v, error: %v", tt.caseName, err)
			}
		}

		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("caseName: %v, result: %v, expected: %v", tt.caseName, result, tt.want)
		}

	}
}
