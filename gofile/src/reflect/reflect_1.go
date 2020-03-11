package main

import(
    "fmt"
    "reflect"
)

type X int

func exam1() {
    var a X=100
    t:=reflect.TypeOf(a)

    fmt.Println(t.Name(), t.Kind())
}

type user struct{
    name string
    age int
}

type manager struct{
    user
    title string
}

func exam2() {
    var m manager
    t:=reflect.TypeOf(&m) 
  
   if t.Kind() ==reflect.Ptr{               // 获取指针的基类型 
       t=t.Elem() 
    } 
  
   for i:=0;i<t.NumField();i++ {    //reflect.Value.NumField()获取结构体中字段的个数
       f:=t.Field(i) 
       fmt.Println("1",f.Name,f.Type,f.Offset) 
  
       if f.Anonymous{               // 输出匿名字段结构 
           for x:=0;x<f.Type.NumField();x++ { 
               af:=f.Type.Field(x) 
               fmt.Println("2","  ",af.Name,af.Type) 
            } 
        } 
    } 
}

func main() {
    // exam1()
    exam2()
}
