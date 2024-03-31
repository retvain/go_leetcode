package suites

type suite28 struct{}

func New28() Suit {
	return &suite28{}
}

func (s *suite28) Run() {
	haystack := "asdffreasdf"
	needle := "fre"
	r := strStr(haystack, needle)
	println(r)
}

func strStr(haystack string, needle string) int {
	result := -1

	shift := len(needle)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		piece := haystack[i:shift]
		if piece == needle {
			return i
		}

		shift++
	}

	return result
}
