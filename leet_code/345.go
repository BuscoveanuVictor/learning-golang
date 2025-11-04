package main
import (
	"strings"
)

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"

	sBytes := []byte(s)	

	i := 0
	j := len(s) - 1
	for i<j {
		if !strings.Contains(vowels, string(sBytes[i])) {
			i++
			continue
		}else if !strings.Contains(vowels, string(sBytes[j])) {
			j--
			continue
		}else{
			sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
			i++
			j--	
		}
	}
	return string(sBytes)
}