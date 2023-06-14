package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var print = fmt.Println
var printf = fmt.Printf

func main() {
}

// input: isValid("()[{}{}]")
// output: true
func isValid(s string) bool {
	sLength := len(s)

	if sLength%2 > 0 {
		return false
	}

	stack := []rune{}

	for i := 0; i < sLength; i++ {
		switch rune(s[i]) {
		case '}':
			if len(stack) <= 0 {
				return false
			}
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]

			} else {
				return false
			}
		case ']':
			if len(stack) <= 0 {
				return false
			}
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]

			} else {
				return false
			}
		case ')':
			if len(stack) <= 0 {
				return false
			}
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]

			} else {
				return false
			}
		default:
			stack = append(stack, rune(s[i]))
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

// input: lengthOfLongestSubstring("bbtablud")
// output: 6
// Because "tablud" is the biggest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	sLen := len([]rune(s))

	if sLen <= 0 {
		return 0
	}

	l := 0

	length := 0

	// Create slice with characters and unique substrings
	for r := 0; r < sLen; r++ {
		char := s[r]

		charIndexInSlice := strings.Index(string(s[l:r]), string(char))

		sliderWindowLength := r - l + 1

		if charIndexInSlice == -1 && sliderWindowLength > length {
			length = sliderWindowLength
		}

		if charIndexInSlice != -1 {
			l += charIndexInSlice + 1
		}
	}

	return length
}

// input: iSumCalcWindow([]int{1, 2, 3, 4, 5, 1, 4, 5, 8, 2}, 2)
// output: [3 7 6 9 10]
func iSumCalcWindow(arr []int, size int) []int {
	slice := []int{}
	arrLen := len(arr)
	for i := 0; i < arrLen; i += size {
		sum := 0
		j := i
		for j < i+size && j <= arrLen-1 {
			sum += arr[j]
			j++
		}
		slice = append(slice, sum)
	}
	return slice
}

func vals() {
	print("5 + 4 =", 5+4)
	print("5 - 4 =", 5-4)
	print("5 * 4 =", 5*4)
	print("5 / 4 =", 5/4)
	print("5 % 4 =", 5%4)
}

func date() {
	now := time.Now()
	print(now.Year(), now.Month(), now.Day())
	print(now.Hour(), now.Minute(), now.Second())
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	res := []int{}
	for i, n := range nums {
		if _, exists := dict[target-n]; exists {
			res = append(res, i)
			res = append(res, dict[target-n])
			return res
		}
		dict[n] = i
	}
	return []int{}
}

func sumTwoNumbers() {
	reader := bufio.NewReader(os.Stdin)
	aAsString, _ := reader.ReadString('\n')
	bAsString, _ := reader.ReadString('\n')
	aTrimmed := strings.TrimSpace(aAsString)
	bTrimmed := strings.TrimSpace(bAsString)
	a, _ := strconv.ParseFloat(aTrimmed, 64)
	b, _ := strconv.ParseFloat(bTrimmed, 64)
	sum := addf(a, b)
	fmt.Println(sum)
}

func addf(a float64, b float64) float64 {
	return a + b
}

func maps() {
	users := make(map[int]string)
	users[1] = "John"
	users[2] = "Carlos"
	users[3] = "Anna"
	print(users)
	delete(users, 1)
	val, exists := users[1]
	print(val, exists)
	print(users)
	for key, value := range users {
		print(key, value)
	}
}

func slicesAddAndRemove() {
	numbers := []int{}
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	print(numbers) // [ 1, 2, 3 ]
	indexToRemove := 1

	firstPart := numbers[:indexToRemove]
	secondPart := numbers[indexToRemove+1:]
	numbers = append(firstPart, secondPart...)

	print(numbers) // [ 1, 3 ]
}

func slices() {
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//[included:exclused]
	slice := numbers[1:9]
	print(numbers) // [1 2 3 4 5 6 7 8 9 10]
	print(slice)   // [2 3 4 5 6 7 8 9]
}

func doWhile() {
	i := 1
	for {
		print(i)
		if i >= i {
			break
		}
	}
}

func while() {
	i := 1
	for i <= 10 {
		print(i)
		i++
	}
}

func iterationOnArr() {
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range numbers {
		print(n)
	}
}

func simpleFor() {
	for i := 1; i <= 10; i++ {
		print(i)
	}
}

func askBudget() {
	reader := bufio.NewReader(os.Stdin)
	budgetAsString, _ := reader.ReadString('\n')
	budget, _ := strconv.ParseFloat(strings.TrimSpace(budgetAsString), 64)
	if budget > 10000 {
		print("Vá para Disney!")
	} else {
		print("Não vá para Disney!")
	}
}

func dates() {
	now := time.Now()
	print(now.Year(), now.Month(), now.Day())
	print(now.Hour(), now.Minute(), now.Second())
	print(now.YearDay())
}

func types() {
	// int32. float64, bool, string, rune
	// 0, 0.0, false, ""
	print(reflect.TypeOf((true)))
	print(reflect.TypeOf((25)))
	print(reflect.TypeOf(("Oi")))
	print(reflect.TypeOf((3.15)))
	print(reflect.TypeOf((true)))
}

func intParse() {
	cV1 := 1.5
	cV2 := int(cV1)
	print(cV2)
}

func strToInt() {
	cV3 := "41412414"
	cV4, err := strconv.Atoi(cV3)
	print(cV4, err, reflect.TypeOf(cV4))
}

func intToAscii() {
	cV5 := 54542452
	cV6 := strconv.Itoa(cV5)
	print(cV6)
}

func strToFloat() {
	cv7 := "3.14"
	cv8, err := strconv.ParseFloat(cv7, 4096)
	if err == nil {
		print(cv8)
	}
}

func askAge() {
	print("Hello! Type your age!")

	reader := bufio.NewReader(os.Stdin)

	ageAsString, _ := reader.ReadString('\n')

	ageAsString = strings.TrimSpace(ageAsString)

	age, _ := strconv.Atoi(ageAsString)

	if age >= 18 {
		print("You are over age!")
	} else {
		print("You are under age!")
	}
}

func stringsPlayground() {
	sv1 := "A word"

	replacer := strings.NewReplacer("A", "Another")

	sv2 := replacer.Replace(sv1)

	print(sv2)

	print("Length:", len(sv2))
	print("Contains 'Another':", strings.Contains(sv2, "Another"))
	print("First index 'o':", strings.Index(sv2, "o"))
	print("Replace:", strings.Replace(sv2, "o", "0", -1))

	sv3 := "\nSome Words\n"
	sv3 = strings.TrimSpace(sv3)
	print("Split :", strings.Split("a-b-c-d", "-"))
	print("Lower :", strings.ToLower(sv2))
	print("Lower :", strings.ToUpper(sv2))
	print("Prefix :", strings.HasPrefix("tacocat", "taco"))
	print("Suffix :", strings.HasSuffix("tacocat", "cat"))
}
