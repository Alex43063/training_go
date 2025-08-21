package main

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

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

	for _, ch := range lowerS1 {
		if !strings.Contains(lowerS2, string(ch)) {
			return false
		}
	}
	return true
}

func FirstUnique(s string) rune {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		count := 0
		for j := 0; j < len(runes); j++ {
			if runes[i] == runes[j] {
				count++
			}
		}
		if count == 1 {
			return runes[i]
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
	re := regexp.MustCompile(`[^a-zA-Z]+`)
	reStr := re.ReplaceAllString(lowerStr, "")
	runes := []rune(reStr)
	runesForReverse := []rune(reStr)
	slices.Reverse(runesForReverse)
	return slices.Equal(runes, runesForReverse)
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
