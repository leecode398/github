package monster
import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Monster struct {
	Name string
	Age int
	skill string
}

//给Monster绑定方法, 可以讲一个Monstar变量(对象), 序列化后保存在文件中
func (this *Monster) Store() bool {
	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}

	//保存到文件
	filePath := "/Users/lx/Documents/git_Projects/github/gofile/src/testing/monster.ser"
	ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return false
	}
	return true
}

//给Monster绑定方法Restore, 可以将一个序列化的Monster从文件中读取
func (this *Monster) ReStore() bool {
	//1. 先从文件读取序列化字符串
	filePath := "/Users/lx/Documents/git_Projects/github/gofile/src/testing/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}

	//2.使用读取到的data, 反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal err =", err)
		return false
	}
	return true
}