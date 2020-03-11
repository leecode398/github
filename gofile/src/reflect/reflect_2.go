package main
import(
    "fmt"
    "reflect"
)

//演示反射
func reflectTest01 (b interface{}) {
    //通过反射获取传入变量的type,kind,值
    //1.先获取到reflect.Type
    rType := reflect.TypeOf(b)
    fmt.Println("rType=", rType)

    //2.获取reflect.value
    rVal := reflect.ValueOf(b)
    fmt.Println("rVal=", rVal)

    //将rVal转换成interface{}
    iV := rVal.Interface()
    //将interface{}通过断言转换成需要的类型
    num2 := iV.(int)
    fmt.Println("num2=", num2)
}

//演示反射
func reflectTest02 (b interface{}) {
    //通过反射获取传入变量的type,kind,值
    //1.先获取到reflect.Type
    rType := reflect.TypeOf(b)
    fmt.Println("rType=", rType)

    //2.获取reflect.value
    rVal := reflect.ValueOf(b)
    fmt.Println("rVal=", rVal)

    //将rVal转换成interface{}
    iV := rVal.Interface()

    fmt.Printf("iv=%v iv type=%T\n", iV, iV)
}

type Student struct {
    Name string
    Age int
}
func main() {
    //1.定义一个int
    // var num int = 100
    // reflectTest01(num)

    //2.定义一个student的实例
    stu := Student{
        Name : "tom",
        Age : 20,
    }

    reflectTest02(stu)
}