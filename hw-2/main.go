package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// type pair struct {
	// 	s [][]int
	// 	r []int
	// }
	// test := []pair{
	// 	{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
	// 	{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
	// 	{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
	// }
	// for _, t := range test {
	// 	s := t.s
	// 	r := t.r
	// 	r2 := Concat(s)
	// 	fmt.Printf("Test: %v\n", s)
	// 	fmt.Printf("Expected: %v\n", r)
	// 	fmt.Printf("Result: %v\n", r2)
	// 	if reflect.DeepEqual(r, r2) {
	// 		fmt.Printf("OK\n\n")
	// 	} else {
	// 		fmt.Printf("FAIL\n\n")
	// 	}
	// }
	c := isDivide(50)
	fmt.Printf("%d ", c)
}

func Concat(slices [][]int) []int {
	result := []int{}
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Pair struct {
	Key   string
	Value int
}

func FrequencyAnalysis(str string) []string {
	arrayOfStrings := strings.Split(str, " ")
	frequency := map[string]int{}
	for _, word := range arrayOfStrings {
		frequency[word] = frequency[word] + 1
	}

	keys := make([]Pair, len(frequency))
	i := 0
	for k, v := range frequency {
		keys[i] = Pair{k, v}
		i++
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Value > keys[j].Value
	})

	length := min(10, len(keys))
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = keys[i].Key
	}

	return result
}

func isDivide(amount int) int {
	count := 0
	for i := 1; i < amount+1; i++ {
		if i%6 == 0 || i%8 == 0 {
			count++
		}
	}
	return count
}
