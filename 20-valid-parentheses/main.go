package main

func isValid(s string) bool {
	stck := []rune{}
	ps := map[rune]rune{'}': '{', ')': '(', ']': '['}

	for _, symbol := range s {
		openingBracket, isClosing := ps[symbol]
		if isClosing && len(stck) != 0 && stck[len(stck)-1] == openingBracket {
			stck = stck[:len(stck)-1]
			continue
		}

		stck = append(stck, symbol)
	}

	return len(stck) == 0
}
