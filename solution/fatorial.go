package solution

func Fatorial(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return 1
	}

	return n * Fatorial(n-1)
}
