package recursive

import "testing"

func TestRecursive(t *testing.T) {
	expected := 120
	result := Recursive(5)
	if expected != result {
		t.Errorf("expected %d, got %d",expected, result)
	}
}

func TestHello(t *testing.T) {
	expected := "Hello Jay, how are you?"
	result := Hello("Jay")
	if expected != result {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
