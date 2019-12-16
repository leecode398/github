/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
package main

import (
	"fmt"
	"regexp"
)

func myAtoi(str string) int {
	reg := regexp.MustCompile(`^[\+\-]?\d+`)
	fmt.Printf("%q\n", reg.FindAllString(str, -1))
	temp := reg.FindAllString(str, -1)
	n := trans(temp)
	return max(-2^31, min(2^31-1, temp))
}
func trans(str string) int {
	i := 1
	m := 0
	if str[0] == '-' {
		i = -1
	}else if str[0] == '+' {
		i = 1
	}else {
		m = str[0]
	}
	for k := m; k< len(str); k++ {
		m = m*10+str[k]
	}
}
func min(i, j int) int{
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int{
	if i > j {
		return i
	}
	return j
}
func main(){
	myAtoi("  asd+42")
}
// @lc code=end

