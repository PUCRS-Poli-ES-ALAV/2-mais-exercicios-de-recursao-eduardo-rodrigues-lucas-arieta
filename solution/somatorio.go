package solution

func Somatorio(n int) int {
	if n == 0 {
		return 0
	}
	return n + Somatorio(n-1)
}
