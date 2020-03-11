package main
import(
    "fmt"
    "reflect"
)

//演示反射
func reflectTest01 (b interface{}) {
    //1.获取reflect.value
    rVal := reflect.ValueOf(b)
    // fmt.Printf("rVal kind=%v\n",rVal.Kind())
    //2.rVal
    rVal.Elem().SetInt(20)
}

func main() {
    // 1.定义一个int
    var num int = 10
    reflectTest01(&num)
    fmt.Println("num=", num)

}