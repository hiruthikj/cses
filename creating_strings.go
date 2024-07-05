package main

import (
	"fmt"
	"slices"
)

func nextPermutation(word []rune) bool {
    pivot := len(word) - 2

    for pivot >= 0 && word[pivot] >= word[pivot+1] {
        pivot--
    }

    if pivot < 0 {
        return false
    }

    nextNum := len(word) - 1

    for word[nextNum] <= word[pivot] {
        nextNum--
    }

    word[pivot], word[nextNum] = word[nextNum], word[pivot]

    i := pivot + 1
    j := len(word) - 1

    for i < j {
        word[i], word[j] = word[j], word[i]
        j--
        i++
    }
    return true
}

func creating_strings() {
	var inpStr string
	fmt.Scanf("%s\n", &inpStr)

	runeArr := []rune(inpStr)
	slices.Sort(runeArr)


	res := []string{string(runeArr)}
	totalPermutations := 1


	for nextPermutation(runeArr) {
		res = append(res, string(runeArr))
		totalPermutations++
	}

	fmt.Println(totalPermutations)
	for _, str := range res {
		fmt.Println(str)
	}
}