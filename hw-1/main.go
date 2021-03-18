package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello")
}

func Addition(a int, b int) int {
	return a + b
}

func byteIsNumber(c byte) bool {
	var tmp int = byteToInt(c)
	if tmp > 0 && tmp < 10 {
		return true
	}
	return false
}

func byteToInt(b byte) int {
	return int(b - '0')
}

func Unpack(str string) (s string) {
	i := 0
	if byteIsNumber(str[0]) {
		return s
	}
	for i < len(str)-1 {
		current := str[i]
		length := 0
		count := 0
		var next byte = str[i+1]
		for byteIsNumber(next) {
			count = count + byteToInt(next)*int(math.Pow10(length))
			length++
			i++
			if i > len(str)-2 {
				break
			}
			next = str[i+1]
		}
		if count == 0 {
			count++
		}
		for j := 0; j < int(count); j++ {
			s = s + string(current)
		}
		i++
	}

	lastRune := str[len(str)-1]
	if !byteIsNumber(lastRune) {
		s = s + string(lastRune)
	}
	return s
}
