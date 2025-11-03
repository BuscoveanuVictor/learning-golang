package main

func minLength(word1 string, word2 string) int {
	if len(word1) < len(word2) {
		return len(word1)
	}
	return len(word2)
}

func mergeStrings(word1 string, word2 string) string {
	result := ""
	for i := 0; i < minLength(word1, word2); i++ {
		result += string(word1[i]) + string(word2[i])
	}

	result += word1[minLength(word1, word2):] + word2[minLength(word1, word2):]
	return result
}


