package homework

import "unicode/utf8"

// PairInput соответствует объекту nums/target из условия контеста.
type PairInput struct {
	Nums   []int
	Target int
}

// longestPrefix возвращает самый длинный общий префикс строк.
func longestPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}

	shortest := &words[0]
	for _, word := range words[1:] {
		if len(word) < len(*shortest) {
			shortest = &word
		}
	}

	for i, char := range *shortest {
		for _, word := range words {
			now, _ := utf8.DecodeRuneInString(word[i:])
			if now != char {
				return word[:i]
			}
		}
	}

	return *shortest
}

// pairIndices возвращает индексы двух разных элементов по возрастанию.
func pairIndices(input PairInput) [2]int {
	target := input.Target

	valueToIndex := make(map[int]int)

	for index, num := range input.Nums {
		index2, exists := valueToIndex[target-num]
		if exists {
			return [2]int{index2, index}
		}

		valueToIndex[num] = index
	}

	return [2]int{}
}

func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

// compactSorted создаёт независимый слайс уникальных значений.
func compactSorted(nums []int) []int {
	res := make([]int, 0)

	if len(nums) == 0 {
		return res
	}

	var last int = nums[0] + 1

	for _, num := range nums {
		if num == last {
			continue
		}
		last = num
		res = append(res, num)
	}

	return res
}

// balancedBrackets проверяет пары и порядок закрытия скобок.
func balancedBrackets(text string) bool {
	stack := make([]int8, 0)

	popAndCompareEquals := func(value int8) bool {
		if len(stack) == 0 {
			return false
		}

		last := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		return last == value
	}

	for _, char := range text {
		switch char {
		case '[':
			stack = append(stack, 0)
		case '{':
			stack = append(stack, 1)
		case '(':
			stack = append(stack, 2)
		case ']':
			if !popAndCompareEquals(0) {
				return false
			}
		case '}':
			if !popAndCompareEquals(1) {
				return false
			}
		case ')':
			if !popAndCompareEquals(2) {
				return false
			}
		default:
			return false
		}
	}

	return len(stack) == 0
}
