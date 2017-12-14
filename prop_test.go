package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestAdd(t *testing.T) {
	type checkFunc func(int) error
	check := func(fns ...checkFunc) []checkFunc { return fns }

	equals := func(want int) checkFunc {
		return func(have int) error {
			if have != want {
				return fmt.Errorf("Expected %d, found %d", want, have)
			}
			return nil
		}
	}

	tests := []struct {
		name   string
		a      int
		b      int
		checks []checkFunc
	}{
		{
			"3 and 2 is 5",
			3,
			2,
			check(
				equals(5),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := Add(tt.a, tt.b)
			for _, check := range tt.checks {
				if err := check(have); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func TestAddProperties(t *testing.T) {
	type checkFunc func(int, int) bool

	tests := []struct {
		name     string
		property checkFunc
	}{
		{
			"commutative",
			func(a, b int) bool { return Add(a, b) == Add(b, a) },
		},
		{
			"first term identity",
			func(a, _ int) bool { return Add(a, 0) == a },
		},
		{
			"second term identity",
			func(_, b int) bool { return Add(0, b) == b },
		},
		{
			"first term nullabile",
			func(a, _ int) bool { return Add(a, -a) == 0 },
		},
		{
			"second term nullabile",
			func(_, b int) bool { return Add(-b, b) == 0 },
		},
		{
			"sum is positive if positive addends",
			func(a, b int) bool {
				if a > 0 && b > 0 {
					return Add(a, b) > 0
				}
				return true
			},
		},
		{
			"sum is negative if negative addends",
			func(a, b int) bool {
				if a < 0 && b < 0 {
					return Add(a, b) < 0
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := quick.Check(tt.property, nil); err != nil {
				t.Error(err)
			}
		})
	}
}
