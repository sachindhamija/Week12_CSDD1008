package main

import "testing"

// TestIsEven tests the IsEven function.
func TestIsEven(t *testing.T) {
	if !IsEven(4) {
		t.Errorf("IsEven(4) should be true, got false")
	}
	if IsEven(5) {
		t.Errorf("IsEven(5) should be false, got true")
	}
}
