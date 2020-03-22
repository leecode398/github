package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func func1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./test.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write([]byte(b))
}

func func2(w http.ResponseWriter, r *http.Request) {
	//Get请求,参数放在URL上(query param)
	queryParam := r.URL.Query() //自动识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) //打印客户端发来的请求body
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello/", func1)
	http.HandleFunc("/xxx/", func2)
	http.ListenAndServe("localhost:9090", nil)

}
