package main

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var re = regexp.MustCompile(`[^a-zA-Z]+`)

func main() {
	// First task
	fmt.Println(WordCount("hello hello World"))
	// Second task
	fmt.Println(AreAnagrams("Hello", "World"))
	// Third task
	fmt.Println(FirstUnique("aabbcc"))
	//Forth task
	fmt.Println(RemoveDuplicates(([]int{4, 4, 4, 4})))
	//Fifth task
	fmt.Println(RemoveElement([]int{1, 2, 3}, 5))
	//Sixth task
	fmt.Println(IsPalindrome("hello"))
	//Seventh task
	DrawChessBoard()

}

func WordCount(s string) map[string]int {
	if len(s) == 0 {
		return nil
	}
	splitStr := strings.Split(s, " ")
	result := make(map[string]int)
	for _, word := range splitStr {
		result[word]++
	}
	return result
}

func AreAnagrams(s1, s2 string) bool {
	withOutSpaceS1 := strings.ReplaceAll(s1, " ", "")
	withOutSpaceS2 := strings.ReplaceAll(s2, " ", "")
	if len(withOutSpaceS1) != len(withOutSpaceS2) {
		return false
	}

	lowerS1 := strings.ToLower(withOutSpaceS1)
	lowerS2 := strings.ToLower(withOutSpaceS2)

	mapCountS1 := make(map[rune]int)
	mapCountS2 := make(map[rune]int)

	for _, char := range lowerS1 {
		mapCountS1[char]++
	}

	for _, char := range lowerS2 {
		mapCountS2[char]++
	}

	for r, _ := range mapCountS1 {
		if mapCountS1[r] != mapCountS2[r] {
			return false
		}
	}

	return true
}

func FirstUnique(s string) rune {
	countSymbols := make(map[rune]int)
	for _, char := range s {
		countSymbols[char]++
	}

	for _, ch := range s {
		if countSymbols[ch] == 1 {
			return ch
		}
	}
	return 0
}

func RemoveDuplicates(nums []int) []int {
	added := make(map[int]bool)
	result := []int{}
	for _, n := range nums {
		if _, ok := added[n]; !ok {
			result = append(result, n)
			added[n] = true
		}
	}
	return result
}

func RemoveElement(nums []int, index int) ([]int, error) {
	if index < 0 || len(nums) <= index {
		return nil, errors.New("index out of range")
	}

	var result []int
	result = append(result, nums[:index]...)
	result = append(result, nums[index+1:]...)
	return result, nil
}

func IsPalindrome(s string) bool {
	lowerStr := strings.ToLower(s)
	reStr := re.ReplaceAllString(lowerStr, "")
	runes := []rune(reStr)
	runesForReverse := []rune(reStr)
	slices.Reverse(runesForReverse)
	return slices.Equal(runes, runesForReverse)
}

func IsPalindromeV2(s string) bool {
	lowerStr := strings.ToLower(s)
	reStr := re.ReplaceAllString(lowerStr, "")
	runes := []rune(reStr)
	for i, j := 0, len(runes)-1; i < j; i++ {
		if runes[i] != runes[j] {
			return false
		}
		j--
	}
	return true
}

func DrawChessBoard() {
	boardSize := 8
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
