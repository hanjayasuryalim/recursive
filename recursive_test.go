package recursive

import "testing"

func TestRecursive(t *testing.T) {
	expected := 105
	result := Recursive(7)
	if expected != result {
		t.Errorf("expected %d, got %d",expected, result)
	}
}

func TestHello(t *testing.T) {
	expected := "Hello Jay"
	result := Hello("Jay")
	if expected != result {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
