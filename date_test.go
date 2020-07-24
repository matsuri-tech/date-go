package mdate

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewDateFromStr(t *testing.T) {
	str := "2019-12-12"
	date, err := NewDateFromStr(str)
	if err != nil {
		t.Error(err)
	}
	expectedDate := NewDate(2019, 12, 12)
	if !reflect.DeepEqual(date, expectedDate) {
		t.Error(date, expectedDate)
	}
}

func TestNewDateFromStr2(t *testing.T) {
	str := "12/12/2019"
	date, err := NewDateFromStr(str)
	if err == nil {
		t.Error(date)
	}
}

func TestDate_PlusNDay(t *testing.T) {
	date := NewDate(2020, 1, 1)
	expectedDate := NewDate(2020, 1, 2)
	if !reflect.DeepEqual(date.PlusNDay(1), expectedDate) {
		t.Error(date)
	}

	date = NewDate(2020, 2, 29)
	expectedDate = NewDate(2020, 3, 1)
	if !reflect.DeepEqual(date.PlusNDay(1), expectedDate) {
		t.Error(date)
	}

	date = NewDate(2019, 2, 28)
	expectedDate = NewDate(2019, 3, 1)
	if !reflect.DeepEqual(date.PlusNDay(1), expectedDate) {
		t.Error(date)
	}
}

func TestDate_MinusNDay(t *testing.T) {
	date := NewDate(2020, 1, 3)
	expectedDate := NewDate(2020, 1, 2)
	if !reflect.DeepEqual(date.MinusNDay(1), expectedDate) {
		t.Error(date)
	}

	date = NewDate(2020, 3, 1)
	expectedDate = NewDate(2020, 2, 29)
	if !reflect.DeepEqual(date.MinusNDay(1), expectedDate) {
		t.Error(date)
	}

	date = NewDate(2019, 3, 1)
	expectedDate = NewDate(2019, 2, 28)
	if !reflect.DeepEqual(date.MinusNDay(1), expectedDate) {
		t.Error(date)
	}
}

func TestNewDateFromStrWithFormat(t *testing.T) {
	str := "12/12/2019"
	date, err := NewDateFromStrWithFormat("01/02/2006", str)
	if err != nil {
		t.Error(err)
	}
	expectedDate := NewDate(2019, 12, 12)
	if !reflect.DeepEqual(date, expectedDate) {
		t.Error(date, expectedDate)
	}
}

func TestNewDateFromStrWithFormat2(t *testing.T) {
	str := "2019-12-12"
	date, err := NewDateFromStrWithFormat("2006-01-02", str)
	if err != nil {
		t.Error(err)
	}
	expectedDate := NewDate(2019, 12, 12)
	if !reflect.DeepEqual(date, expectedDate) {
		t.Error(date, expectedDate)
	}
}

func TestDateDiff(t *testing.T) {
	d1 := NewDate(2018, 10, 11)
	d2 := NewDate(2018, 10, 12)
	diff := d1.DateDiff(d2)
	if diff != 1 {
		t.Errorf("expected: 1 actual %d", diff)
	}
}

func TestDateDiff2(t *testing.T) {
	d1 := NewDate(2018, 10, 12)
	d2 := NewDate(2018, 10, 11)
	diff := d1.DateDiff(d2)
	if diff != -1 {
		t.Errorf("expected: 1 actual %d", diff)
	}
}

func TestAddDate(t *testing.T) {
	d1 := NewDate(2018, 10, 12)
	result := d1.PlusNDay(1)
	if !d1.PlusNDay(1).IsEqual(result) {
		t.Errorf("expcted equality actual d1 %s result %s", d1, result)
	}
}

func TestIsLater(t *testing.T) {
	d1 := NewDate(2018, 10, 12)
	d2 := NewDate(2018, 10, 11)
	if !d1.IsLater(d2) {
		t.Errorf("")
	}
}

type testCaseIsNextDay struct {
	prev Date
	next Date
}

func TestDate_IsNextDay_True(t *testing.T) {
	testCases := []testCaseIsNextDay{
		{
			prev: NewDate(2020, 1, 1),
			next: NewDate(2020, 1, 2),
		},
		{
			prev: NewDate(2019, 12, 31),
			next: NewDate(2020, 1, 1),
		},
		{
			prev: NewDate(2019, 2, 28),
			next: NewDate(2019, 3, 1),
		},
	}

	for _, tt := range testCases {
		if got := tt.next.IsNextDay(tt.prev); !got {
			t.Errorf("want true, got %v", got)
		}
	}
}

func TestDate_IsNextDay_False(t *testing.T) {
	testCases := []testCaseIsNextDay{
		{
			prev: NewDate(2020, 1, 1),
			next: NewDate(2020, 1, 3),
		},
		{
			prev: NewDate(2020, 2, 28),
			next: NewDate(2020, 3, 1),
		},
	}

	for _, tt := range testCases {
		if got := tt.next.IsNextDay(tt.prev); got {
			t.Errorf("want false, got %v", got)
		}
	}
}

func TestDate_JapanFiscalYear(t *testing.T) {
	tests := []struct {
		in   Date
		want JapanFiscalYear
	}{
		{
			in:   NewDate(2020, 1, 2),
			want: 2019,
		},
		{
			in:   NewDate(2020, 5, 6),
			want: 2020,
		},
	}

	for _, tt := range tests {
		if tt.in.JapanFiscalYear() != tt.want {
			t.Error(tt)
		}
	}
}

func TestDate_MarshalJSON(t *testing.T) {

	type Sample struct {
		D Date `json:"d"`
	}

	tests := []struct {
		in   Sample
		want string
	}{
		{
			in:   Sample{D: NewDate(2020, 2, 2)},
			want: `{"d":"2020-02-02"}`,
		},
	}
	for _, tt := range tests {
		b, err := json.Marshal(tt.in)
		if err != nil {
			t.Error(err)
		}
		if tt.want != string(b) {
			t.Error(string(b), tt.want)
		}
	}
}
