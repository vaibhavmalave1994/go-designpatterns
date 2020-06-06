package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Expected pointer to the singleton after calling GetInstance(), not nil")
	}

	expectedCounter := counter1

	currentCounter := counter1.AddOne()

	if currentCounter != 1 {
		t.Errorf("After calling for the first time for instance count must be 1, but it is  %d\n", currentCounter)
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but got different instance")
	}

	currentCounter2 := counter2.AddOne()

	if currentCounter2 != 2 {
		t.Errorf("After calling AddOne() using second counter the counter count must be 2, but it is  %d\n", currentCounter)
	}
}
