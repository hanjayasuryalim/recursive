package recursive

import "fmt"

func Recursive(number int) int {
	if number < 1 {
		return 1
	}
	return number * Recursive(number-2)
}

func Hello(name string) string {
	return fmt.Sprintf("Hello %s, how are you?", name)
}
