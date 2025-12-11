package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
第三题、第六题和第七题借助了ai帮助，其余题目为自己思考实现，请老师帮忙检查
*/
func main() {
	// 第一题：只出现一次的数字
	nums := []int{2, 2, 3, 4, 3}
	fmt.Println(getOnce(nums))

	// 第二题：回文数
	//back := 121
	back := -121
	fmt.Println(isBackNum(back))

	// 第三题：有效括号
	//kuohao := "()[]{}"
	//kuohao := "([)]"
	kuohao3 := "{[]}"
	fmt.Println(isKuoHao(kuohao3))

	// 第四题：最长公共前缀
	arr4 := []string{"flower", "flow", "flight"}
	fmt.Println(getLongCommon(arr4))

	// 第五题：加一
	arr5 := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 9}
	fmt.Println(addOne(arr5))

	// 第六题：删除有序数组中的重复项
	arr6 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(arr6)
	fmt.Println(k, arr6[:k])
	//第七题：合并区间
	arr7 := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	mergeArr(arr7)
	fmt.Println(mergeArr(arr7))

	//第八题:两数之和
	arr8 := []int{2, 7, 1, 1, 5}
	targer := 6
	fmt.Println(twoSum(arr8, targer))

}

//1、生成只出现一次的数字
/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
*/
func getOnce(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		//用值作为map的键名，重复就给该键+1
		m[v]++
	}
	var ret int
	for k, v := range m {
		if v == 1 { //遍历map找到value是1的键
			ret = int(k)
		}
	}
	return ret
}

//2、回文数
/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func isBackNum(x int) bool {
	str1 := strconv.Itoa(x) //转换为字符串
	len := len(str1)
	var str2 string
	for i := len - 1; i >= 0; i-- {
		str2 += string(str1[i]) //逆序
	}
	if str1 == str2 {
		return true
	} else {
		return false
	}
}

//3、有效括号
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func isKuoHao(s string) bool {
	// 创建一个栈来存储左括号
	stack := []rune{}

	// 定义括号映射关系
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		// 如果是右括号
		if pair, ok := pairs[ch]; ok {
			// 栈为空或栈顶元素不匹配则无效
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 左括号入栈
			stack = append(stack, ch)
		}
	}

	// 栈为空则所有括号都正确匹配
	return len(stack) == 0
}

//4、最长公共前缀
/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/
func getLongCommon(arr []string) string {
	minLen := len(arr[0])
	for _, v := range arr { //找到最短的字符串
		if len(v) < minLen {
			minLen = len(v)
		}
	}
	for i := 0; i < minLen; i++ {
		char := arr[0][i]
		for j := 1; j < len(arr); j++ {
			if char != arr[j][i] {
				return arr[0][:i]
			}
		}
	}
	return arr[0][:minLen] //最短字符串即是最大公共前缀
}

//5、加一
/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。
*/

func addOne(digits []int) []int {
	str := ""
	back := []int{}
	for i := 0; i < len(digits); i++ {
		str += strconv.Itoa(digits[i])
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("strconv.Atoi error")
	}
	num += 1
	str = strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		n, _ := strconv.Atoi(string(str[i]))
		back = append(back, n)
	}
	return back
}

// 6、删除重复项
/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用双指针方法
	// slow 指针指向不重复元素的位置
	// fast 指针遍历整个数组
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// 当发现不同的元素时，移动slow指针，并将新元素放到slow位置
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		// fmt.Println(nums)
	}

	// 返回不重复元素的个数
	return slow + 1
}

//7、合并区间
/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
*/

func mergeArr(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// 首先按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// fmt.Println("排序后：", intervals)
	// 合并重叠区间
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// fmt.Println("第:", i, "次循环")
		last := result[len(result)-1]
		// fmt.Println("last:", last)
		current := intervals[i]
		// fmt.Println("current:", current)

		// 如果当前区间的起始位置小于等于上一个区间的结束位置，则有重叠
		if current[0] <= last[1] {
			// 合并区间，更新结束位置为两个区间结束位置的最大值
			last[1] = max(last[1], current[1])
		} else {
			// 没有重叠，直接添加新区间
			result = append(result, current)
			// fmt.Println(result)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 8、两数之和
/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。
*/

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
