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

//反序列化为结构体
func unmarshalStruct() {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2020-2-29\",\"Sal\":8000}"

	//定义一个Monstar实例
	var monster Monster

	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v\n", monster)
}

//反序列化为map
func unmarshalMap() {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2020-2-29\",\"Sal\":8000}"

	//定义一个map实例
	var a map[string]interface{}	//反序列化无需make,unmarshal底层会自动make

	//反序列化
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v\n", a)
}

//反序列化为map
func unmarshalSlice() {
	str := "[{\"adress\":\"beijing\",\"age\":\"7\",\"name\":\"jack\"}," + 
	"{\"adress\":\"nanjing\",\"age\":\"20\",\"name\":\"tom\"}]"

	//定义一个slice实例
	var slice []map[string]interface{}	//反序列化无需make,unmarshal底层会自动make

	//反序列化
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}