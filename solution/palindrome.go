package solution

// Um palindromo é uma palavra que pode ser lida da mesma forma de trás para frente e de frente para trás.
func Palindrome(s string) bool {
	// e.g. se a palavra tiver somente um caractere ou nenhum, é um palíndromo
	if len(s) <= 1 {
		return true
	}

	length := len(s)
	first := s[0]
	last := s[length-1]
	// se o primeiro e o último caractere não forem iguais, não é um palíndromo
	if first != last {
		return false
	}

	// chama a função recursivamente deslocando o primeiro e o último caractere
	return Palindrome(s[1 : length-1])
}
