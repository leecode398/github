package main

//被测试函数
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n-1; i++ {
		res += i
	}
	return res
}