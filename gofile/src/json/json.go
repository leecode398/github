package main
import(
    "fmt"
    "encoding/json"
)

// 定义一个结构体
type Monster struct{
    Name string  `json:"name"`  //指定反序列化后的名字
    Age int  `json:"age"`
    Birthday string
    Sal float64
    skill string
}
//结构体序列化
func testStruct() {
    monster := Monster{
        Name : "牛魔王",
        Age : 500,
        Birthday : "2020-2-29",
        Sal : 8000.0,
        skill : "牛魔拳",
    }

    //序列化
    data, err := json.Marshal(&monster)
    if err != nil {
        fmt.Printf("序列号错误 err=%v\n", err)
    }
    //输出序列化结果
    fmt.Printf("monster序列化后=%v\n", string(data))
}

//map序列化
func testMap() {
    //定义一个map
    var a map[string]interface{}
    //使用map前先make
    a = make(map[string]interface{})
    a["name"] = "红孩儿"
    a["age"] = 30
    a["adress"] = "火云洞"

    //将a序列化
    data, err := json.Marshal(&a)
    if err != nil {
        fmt.Printf("序列号错误 err=%v\n", err)
    }
    //输出序列化结果
    fmt.Printf("a序列化后=%v\n", string(data))
}

//切片序列化
func testSlice() {
    var slice []map[string]interface{}
    var m1 map[string]interface{}
    m1 = make (map[string]interface{})
    m1["name"] = "jack"
    m1["age"] = "7"
    m1["adress"] = "beijing"
    slice = append(slice, m1)

    var m2 map[string]interface{}
    m2 = make (map[string]interface{})
    m2["name"] = "tom"
    m2["age"] = "20"
    m2["adress"] = "nanjing"
    slice = append(slice, m2)

    //将a序列化
    data, err := json.Marshal(slice)
    if err != nil {
        fmt.Printf("序列号错误 err=%v\n", err)
    }
    //输出序列化结果
    fmt.Printf("slice列化后=%v\n", string(data))

}

//基本数据类型序列化
func testFloat64() {
    var num1 float64 = 23.45

    //将num1序列化
    data, err := json.Marshal(num1)
    if err != nil {
        fmt.Printf("序列号错误 err=%v\n", err)
    }
    //输出序列化结果
    fmt.Printf("num1列化后=%v\n", string(data))

}
func main() {
    testStruct()
    testMap()
    testSlice()
    testFloat64()
}