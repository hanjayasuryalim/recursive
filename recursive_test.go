package recursive

import "testing"

func TestRecursive(t *testing.T) {
	expected := 120
	result := Recursive(5)
	if expected != result {
		t.Errorf("expected %d, got %d",expected, result)
	}
}
