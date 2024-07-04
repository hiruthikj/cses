package main

import (
	"fmt"
	"maps"
	"slices"
)

func creating_strings() {
	var inpStr string
	fmt.Scanf("%s\n", &inpStr)

	charCount := make(map[rune]int)
	
	for _, ch := range inpStr {
		charCount[ch] += 1
	}
	
	// fmt.Println(charCount)
	
	result := make([]string, 1)
	r := permuteString(charCount, "", len(inpStr), result)
	
	slices.Sort(r)
	fmt.Println(len(r))

	for _, str := range r {
		fmt.Println(str)
	}
}

func permuteString(charCountRem map[rune]int, currStr string, remCount int, result []string) []string {
	if remCount == 0 {
		// result = append(result, currStr)
		// fmt.Println(currStr)
		return []string{currStr}
	}

	r := []string{}
	for ch, _ := range charCountRem {
		charCountCp := make(map[rune]int)
		maps.Copy(charCountCp, charCountRem)
		charCountCp[ch] -= 1
		if charCountCp[ch] <= 0 {
			delete(charCountCp, ch)
		}
		// fmt.Println(charCountCp)
		r = append(r, permuteString(charCountCp, currStr+string(ch), remCount-1, result)...)
	}

	return r
}
