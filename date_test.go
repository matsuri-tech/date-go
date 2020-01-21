package mdate

import (
	"testing"
)

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
