package pointers

import "testing"

func TestIncrement(t *testing.T) {
	originalValue := 5
	expectedValue := originalValue + 1
	Increment(&originalValue)
	if originalValue != expectedValue {
		t.Errorf("Increment did not increment the value correctly. Got %d; want %d.", originalValue, expectedValue)
	}
	t.Logf("TEST: OK TestIncrement %d = expectedValue %d", originalValue, expectedValue)
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("Swap did not work as expected. After swap, got a = %d, b = %d; want a = 2, b = 1.", a, b)
	}
	t.Logf("TEST: OK TestSwap a = %d, b = %d", a, b)
}

func TestSetAge(t *testing.T) {
	person := Person{Name: "John", Age: 30}
	newAge := 35
	SetAge(&person, newAge)
	if person.Age != newAge {
		t.Errorf("SetAge did not set the correct age. Got %d; want %d.", person.Age, newAge)
	}
	t.Logf("TEST: OK TestSetAge person.Age = %d, newAge = %d", person.Age, newAge)
}
