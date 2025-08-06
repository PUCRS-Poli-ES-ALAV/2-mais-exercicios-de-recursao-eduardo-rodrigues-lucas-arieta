package solution

func SomatorioEntreNums(k, j int) int {
	if k > j {
		return 0
	}
	if k == j {
		return k
	}
	return k + SomatorioEntreNums(k+1, j)
}
