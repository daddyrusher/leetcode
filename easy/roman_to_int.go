package easy

func romanToInt(s string) int {
	dict := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && dict[runes[i]] < dict[runes[i+1]] {
			sum -= dict[runes[i]]
		} else {
			sum += dict[runes[i]]
		}
	}

	return sum
}
