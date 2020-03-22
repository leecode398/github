package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//公用一个client适用于请求频繁
var (
	client = http.Client{
		Transport: &http.Transport{
		DisableKeepAlives: false,
		},
	}
)
func main() {
	// resp, err := http.Get("http://47.100.185.222:9090/xxx/?name=dss&age=23")
	// if err != nil {
	// 	fmt.Println("get url failed, err:%v\n", err)
	// 	return
	// }
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "dss")
	data.Set("age", "23")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("Get", urlObj.String(), nil)
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("get url failed, err:%v\n", err)
	// 	return
	// }
	//请求不频繁,用完就关闭连接
	//禁用KeepAlive的client
	// tr := http.Transport{
	// 	DisableKeepAlives: true,
	// }
	// client := http.Client{
	// 	Transport: tr,
	// }
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get url failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	//从resp中把服务端返回的数据读出来
	// var data []byte
	// resp.Body.Read()
	// resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body) //读取服务器的请求body
	if err != nil {
		fmt.Printf("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}