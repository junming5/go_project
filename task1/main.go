package main

import "sort"

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/description/
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		value, exist := m[v]
		if exist {
			value++
			m[v] = value
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0

	// zero := 0
	// for _, v := range nums {
	// 	zero ^= v
	// }
	// return zero
}

// 9. 回文数
// https://leetcode.cn/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	num := 0
	bitNum := x
	for bitNum > 0 {
		num *= 10
		num += bitNum % 10

		bitNum /= 10
	}
	return num == x
}

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/description/
func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	m := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stackSli := []rune{}
	for _, v := range s {
		value, _ := m[v]
		if len(stackSli) > 0 && stackSli[len(stackSli)-1] == value {
			stackSli = append(stackSli[:len(stackSli)-1])
		} else {
			stackSli = append(stackSli, v)
		}
	}
	return len(stackSli) == 0
}

// 14. 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	// cStr := strs[0]
	// strs = strs[1:]

	// for _, str := range strs {
	// 	tmpRunes := make([]rune, 0)
	// 	for i, r := range str {
	// 		if i < len(cStr) && str[i] == cStr[i] {
	// 			tmpRunes = append(tmpRunes, r)
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	cStr = string(tmpRunes)
	// }
	// return cStr

	cStr := strs[0]
	for i := 1; i < len(strs); i++ {
		index := 0
		str := strs[i]
		for index < len(cStr) && index < len(str) && str[index] == cStr[index] {
			index++
		}
		cStr = cStr[:index]
		if len(cStr) == 0 {
			return ""
		}
	}
	return cStr
}

// 66. 加一
// https://leetcode.cn/problems/plus-one/description/
func plusOne(digits []int) []int {
	count := len(digits)
	isAdd := false
	for i := count - 1; i >= 0; i-- {
		v := digits[i]
		v = v + 1
		if v <= 9 {
			digits[i] = v
			return digits
		} else {
			digits[i] = 0
			isAdd = i == 0
		}
	}
	if isAdd {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	count := len(nums)
	if count <= 1 {
		return 1
	}
	preNum := nums[0]
	k := 1
	for i := 1; i < count; i++ {
		num := nums[i]
		if num != preNum {
			if k != i {
				nums[k] = num
			}
			k++
		}
		preNum = num
	}
	return k
}

// 56. 合并区间
// https://leetcode.cn/problems/merge-intervals/description/
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	resSli := intervals[:1]
	for i := 1; i < len(intervals); i++ {
		num := intervals[i]
		res := resSli[len(resSli)-1]
		if res[1] < num[0] {
			resSli = append(resSli, num)
		} else {
			res[1] = max(res[1], num[1])
		}
	}
	return resSli
}

// 1. 两数之和
// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	// count := len(nums)
	// for i := 0; i < count-1; i++ {
	// 	for j := i+1; j < count; j++ {
	// 		if nums[i]+nums[j] == target {
	// 			return []int{i,j}
	// 		}
	// 	}
	// }
	// return nil

	m := make(map[int]int, len(nums))
	for i, v := range nums {
		mV, exist := m[target-v]
		if exist {
			return []int{i, mV}
		}
		m[v] = i
	}
	return nil
}

func main() {
	// isValid(")(){}")
	// longestCommonPrefix([]string{"cir"})
	// plusOne([]int{9, 9})
	// removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})

	// merge([][]int{[]int{1, 4}, []int{2, 3}})
	twoSum([]int{2, 7, 11, 15}, 9)
}
