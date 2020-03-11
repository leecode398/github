package main
import(
	"fmt"
)

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func pMax(a int, p []int) int {
	q := 0
	if a == 0 {
		return 0
	}
	for i := 0; i < a; i++ {
		q = Max(q, p[i] + pMax(a-i-1, p))
	}
	return q
}
func main() {
	p := []int{1,2,8,9,10,17,17,20,24,30}
	a := 10
	fmt.Println("最大利润:", pMax(a, p))
}