// Need go.mod to run test
package main

import "testing"

var tests = []struct {
	// Data I want to test
	name string
	dividend float32
	divisor float32
	expected float32
	isError bool
}{
	// every entry in this slice run once
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-20", 100.0, 5.0, 20.0, false},
}

// `t` and `tt` is conventional
func TestDivision(t *testing.T) {
	for _,tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Error("Got an error but did not get one")
			}
		}  else {
			if err != nil {
				t.Error("Did not expect an err but got one", err.Error())
			}
		}
		if got != tt.expected {
			t.Error("different got and expected")
		}
	}
}


// func TestDivide(t *testing.T) {
// 	// function is sharable with test code
// 	_, err := divide(10.0, 1.0)
// 	// err is anti case of test anticipation always
// 	if err != nil {
// 		t.Error("Got an error whe we should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0)
// 	if err == nil {
// 		t.Error("Did not get an err whe we should have")
// 	}
// }