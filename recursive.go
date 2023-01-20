package recursive

func Recursive(number int) int {
	if number == 1 {
		return 1
	}
	return number * Recursive(number-1)
}
