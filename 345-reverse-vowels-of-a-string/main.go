package main

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	chars := []byte(s)

	left, right := 0, len(s)-1

	for left < right {
		if vowels[chars[left]] && vowels[chars[right]] {
			chars[left], chars[right] = chars[right], chars[left]

			left++
			right--
		} else if !vowels[chars[left]] {
			left++
		} else {
			right--
		}
	}

	return string(chars)
}
