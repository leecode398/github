# go note

### 匿名函数

匿名函数：顾名思义就是没有名字的函数。匿名函数最大的用途是来模拟块级作用域,避免数据污染的。

##### 实例:

1.

```go
package main

import (
   "fmt"
)
func main() {
   f:=func(){
      fmt.Println("hello world")
   }
   f()//hello world
   fmt.Printf("%T\n", f) //打印 func()
}
```

2.带参数

```go
package main

import (
   "fmt"
)
func main() {
   f:=func(args string){
      fmt.Println(args)
   }
   f("hello world")//hello world
   //或
   (func(args string){
        fmt.Println(args)
    })("hello world")//hello world
    //或
    func(args string) {
        fmt.Println(args)
```

3.带返回值

```go
package main

import "fmt"

func main() {
   f:=func()string{
      return "hello world"
   }
   a:=f()
   fmt.Println(a)//hello world
}
```

4.多个匿名函数

```go
package main

import "fmt"

func main() {
   f1,f2:=F(1,2)
   fmt.Println(f1(4))//6
   fmt.Println(f2())//6
}
func F(x, y int)(func(int)int,func()int) {
   f1 := func(z int) int {
      return (x + y) * z / 2
   }

   f2 := func() int {
      return 2 * (x + y)
   }
   return f1,f2
}
```

### 闭包

闭包：闭包是由函数和与其相关的引用环境组合而成的实体。

```go
func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

```

调用这个函数会返回一个函数变量。

i := incr()：通过把这个函数变量赋值给 i，i 就成为了一个闭包。

所以 i 保存着对 x 的引用，可以想象 i 中有着一个指针指向 x 或 i 中有 x 的地址。

由于 i 有着指向 x 的指针，所以可以修改 x，且保持着状态：

```go
println(i()) // 1
println(i()) // 2
println(i()) // 3
```

也就是说，x 逃逸了，它的生命周期没有随着它的作用域结束而结束。

但是这段代码却不会递增：

```go
println(incr()()) // 1
println(incr()()) // 1
println(incr()()) // 1
```

这是因为这里调用了三次incr()，返回了三个闭包，这三个闭包引用着三个不同的 x，它们的状态是各自独立的。

##### 示例

###### 1.

```go
package main

import "fmt"

func main() {
    a := Fun()
    b:=a("hello ")
    c:=a("hello ")
    fmt.Println(b)//worldhello
    fmt.Println(c)//worldhello hello 
}
func Fun() func(string) string {
    a := "world"
    return func(args string) string {
        a += args
        return  a
    }
}
```

###### 2.

```go
package main

import "fmt"

func main() {
   a := Fun()
   d := Fun()
   b:=a("hello ")
   c:=a("hello ")
   e:=d("hello ")
   f:=d("hello ")
   fmt.Println(b)//worldhellod
   fmt.Println(c)//worldhello hello
   fmt.Println(e)//worldhello
   fmt.Println(f)//worldhello hello
}
func Fun() func(string) string {
   a := "world"
   return func(args string) string {
      a += args
      return  a
   }
}
```

注意两次调用F()，维护的不是同一个a变量。

###### 3.

```go
package main

import "fmt"

func main() {
   a := F()
   a[0]()//0xc00004c080 3
   a[1]()//0xc00004c080 3
   a[2]()//0xc00004c080 3
}
func F() []func() {
   b := make([]func(), 3, 3)
   for i := 0; i < 3; i++ {
      b[i] = func() {
         fmt.Println(&i,i)
      }
   }
}
```

闭包通过引用的方式使用外部函数的变量。例中只调用了一次函数F,构成一个闭包，i 在外部函数B中定义，所以闭包维护该变量 i ，a[0]、a[1]、a[2]中的 i 都是闭包中 i 的引用。因此执行,i 的值已经变为3，故再调用a[0]()时的输出是3而不是0。

###### 4.

如何避免上面的BUG ，用下面的方法，注意和上面示例对比。

```go
package main

import "fmt"

func main() {
    a := F()
    a[0]() //0xc00000a0a8 0
    a[1]() //0xc00000a0c0 1
    a[2]() //0xc00000a0c8 2
}
func F() []func() {
    b := make([]func(), 3, 3)
    for i := 0; i < 3; i++ {
        b[i] = (func(j int) func() {
            return func() {
                fmt.Println(&j, j)
            }
        })(i)
    }
    return b
}
或者
package main

import "fmt"

func main() {
    a := F()
    a[0]() //0xc00004c080 0
    a[1]() //0xc00004c088 1
    a[2]() //0xc00004c090 2
}
func F() []func() {
    b := make([]func(), 3, 3)
    for i := 0; i < 3; i++ {
        j := i
        b[i] = func() {
            fmt.Println(&j, j)
        }
    }
    return b
}
```

每次 操作仅将匿名函数放入到数组中，但并未执行，并且引用的变量都是 i，随着 i 的改变匿名函数中的 i 也在改变，所以当执行这些函数时，他们读取的都是环境变量 i 最后一次的值。解决的方法就是每次复制变量 i 然后传到匿名函数中，让闭包的环境变量不相同。

###### 5.

```go
package main

import "fmt"

func main() {
   fmt.Println(F())//2
}
func F() (r int) {
   defer func() {
      r++
   }()
   return 1
}
```

输出结果为2，即先执行r=1 ,再执行r++。

###### 6.递归函数

```go
package main

import "fmt"

func F(i int) int {
   if i <= 1 {
      return 1
   }
   return i * F(i-1)
}

func main() {
   var i int = 3
   fmt.Println(i, F(i))// 3 6
}
```

###### 7.斐波那契数列(Fibonacci)

```go
package main

import "fmt"

func fibonaci(i int) int {
    if i == 0 {
        return 0
    }
    if i == 1 {
        return 1
    }
    return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
    var i int
    for i = 0; i < 10; i++ {
        fmt.Printf("%d\n", fibonaci(i))
    }
}
```

### 切片

##### 1.创建切片

通过 make() 函数创建切片

```go
// 创建一个整型切片
// 其长度和容量都是 5 个元素
slice := make([]int, 5)
```

```go
// 创建一个整型切片
// 其长度为 3 个元素，容量为 5 个元素
slice := make([]int, 3, 5)
```

```go
Golang 不允许创建容量小于长度的切片，当创建的切片容量小于长度时会在编译时刻报错:

// 创建一个整型切片
// 使其长度大于容量
myNum := make([]int, 5, 3)
```

通过字面量创建切片

```go
// 创建字符串切片
// 其长度和容量都是 3 个元素
myStr := []string{"Jack", "Mark", "Nick"}
// 创建一个整型切片
// 其长度和容量都是 4 个元素
myNum := []int{10, 20, 30, 40}
```

##### 2.区别数组和切片的声明方式

当使用字面量来声明切片时，其语法与使用字面量声明数组非常相似。二者的区别是：如果在 [] 运算符里指定了一个值，那么创建的就是数组而不是切片。只有在 [] 中不指定值的时候，创建的才是切片。

```go
// 创建有 3 个元素的整型数组
myArray := [3]int{10, 20, 30}
// 创建长度和容量都是 3 的整型切片
mySlice := []int{10, 20, 30}
```

##### 3.nil和空切片

```go
nil切片
// 创建 nil 整型切片
var myNum []int
```

```go
空切片
// 使用 make 创建空的整型切片
myNum := make([]int, 0)
// 使用切片字面量创建空的整型切片
myNum := []int{}
```

不管是使用 nil 切片还是空切片，对其调用内置函数 append()、len() 和 cap() 的效果都是一样的。

##### 4.为切片中的元素赋值

对切片里某个索引指向的元素赋值和对数组里某个索引指向的元素赋值的方法完全一样。使
用 [] 操作符就可以改变某个元素的值，下面是使用切片字面量来声明切片：

```go
// 创建一个整型切片
// 其容量和长度都是 5 个元素
myNum := []int{10, 20, 30, 40, 50}
// 改变索引为 1 的元素的值
myNum [1] = 25
```

##### 5.通过切片创建新的切片

切片之所以被称为切片，是因为创建一个新的切片，也就是把底层数组切出一部分。通过切片创建新切片的语法如下：

```go
slice[i:j]
slice[i:j:k]
```

其中 i 表示从 slice 的第几个元素开始切，j 控制切片的长度(j-i)，k 控制切片的容量(k-i)，如果没有给定 k，则表示切到底层数组的最尾部。下面是几种常见的简写形式：

```go
slice[i:]  // 从 i 切到最尾部
slice[:j]  // 从最开头切到 j(不包含 j)
slice[:]   // 从头切到尾，等价于复制整个 slice
```

让我们通过下面的例子来理解通过切片创建新的切片的本质：

```go
// 创建一个整型切片
// 其长度和容量都是 5 个元素
myNum := []int{10, 20, 30, 40, 50}
// 创建一个新切片
// 其长度为 2 个元素，容量为 4 个元素
newNum := slice[1:3]
```

##### 6.共享底层数组的切片

由切片新建的切片共享底层数组,如果一个切片修改了该底层数组的共享部分,另一个切片也受影响

```go
// 修改 newNum 索引为 1 的元素
// 同时也修改了原切片 myNum 的索引为 2 的元素
newNum[1] = 35
```

把 35 赋值给 newNum 索引为 1 的元素的同时也是在修改 myNum 索引为 2 的元素：

切片只能访问到其长度内的元素，试图访问超出其长度的元素将会导致语言运行时异常。在使用这部分元素前，必须将其合并到切片的长度里。下面的代码试图为 newNum 中的元素赋值：

```go
// 修改 newNum 索引为 3 的元素
// 这个元素对于 newNum 来说并不存在
newNum[3] = 45
```

##### 7.切片扩容

相对于数组而言，使用切片的一个好处是：可以按需增加切片的容量。Golang 内置的 append() 函数会处理增加长度时的所有操作细节。要使用 append() 函数，需要一个被操作的切片和一个要追加的值，当 append() 函数返回时，会返回一个包含修改结果的新切片。函数 append() 总是会增加新切片的长度，而容量有可能会改变，也可能不会改变，这取决于被操作的切片的可用容量。

```go
myNum := []int{10, 20, 30, 40, 50}
// 创建新的切片，其长度为 2 个元素，容量为 4 个元素
newNum := myNum[1:3]
// 使用原有的容量来分配一个新元素
// 将新元素赋值为 60
newNum = append(newNum, 60)
```

由于共享底层数组,所以此时切片myNum值为{10, 20, 30, 60, 50}

```go
// 创建一个长度和容量都是 4 的整型切片
myNum := []int{10, 20, 30, 40}
// 向切片追加一个新元素
// 将新元素赋值为 50
newNum := append(myNum, 50)
```

当这个 append 操作完成后，newSlice 拥有一个全新的底层数组，这个数组的容量是原来的两倍.函数append()会智能地出路底层数组的容量曾长.当切片容量小于1000个元素时,总是会成倍地增加容量.一旦超过1000,容量的增长因子会设为1.25,也就是每次增加25%的容量.

##### 8.限制切片的容量

在创建切片时，使用第三个索引选项引可以用来控制新切片的容量。其目的并不是要增加容量，而是要限制容量。允许限制新切片的容量为底层数组提供了一定的保护，可以更好地控制追加操作。

```go
// 创建长度和容量都是 5 的字符串切片
fruit := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
```

```go
// 将第三个元素切片，并限制容量
// 其长度为 1 个元素，容量为 2 个元素
myFruit := fruit[2:3:4]
```

如果设置的容量比可用的容量还大，就会得到一个运行时错误：

```go
myFruit := fruit[2:3:6]
```

##### 9.遍历切片

切片是一个集合，可以迭代其中的元素。Golang 有个特殊的关键字 range，它可以配合关键字 for 来迭代切片里的元素

```go
myNum := []int{10, 20, 30, 40, 50}
// 迭代每一个元素，并显示其值
for index, value := range myNum {
    fmt.Printf("index: %d value: %d\n", index, value)
}
```

输出:

```go
index: 0 value: 10
index: 1 value: 20
index: 2 value: 30
index: 3 value: 40
index: 4 value: 50
```

当迭代切片时，关键字 range 会返回两个值。第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本。需要强调的是，range 创建了每个元素的副本，而不是直接返回对该元素的引用。要想获取每个元素的地址，可以使用切片变量和索引值：

```go
myNum := []int{10, 20, 30, 40, 50}
// 修改切片元素的值
// 使用空白标识符(下划线)来忽略原始值
for index, _ := range myNum {
    myNum[index] += 1
}
for index, value := range myNum {
    fmt.Printf("index: %d value: %d\n", index, value)
}
```

输出:

```go
index: 0 value: 11
index: 1 value: 21
index: 2 value: 31
index: 3 value: 41
index: 4 value: 51
```

关键字 range 总是会从切片头部开始遍历。如果想对遍历做更多的控制，可以使用传统的 for 循环配合 len() 函数实现：

```go
myNum := []int{10, 20, 30, 40, 50}
// 从第三个元素开始迭代每个元素
for index := 2; index < len(myNum); index++ {
    ...
}
```

##### 10.切片的拷贝操作

Golang 内置的 copy() 函数可以将一个切片中的元素拷贝到另一个切片中，其函数声明为：

```go
func copy(dst, src []Type) int
```

它表示把切片 src 中的元素拷贝到切片 dst 中，返回值为拷贝成功的元素个数。如果 src 比 dst 长，就截断；如果 src 比 dst 短，则只拷贝 src 那部分：

```go
num1 := []int{10, 20, 30}
num2 := make([]int, 5)
count := copy(num2, num1)
fmt.Println(count)
fmt.Println(num2)
```

输出:

```go
3
[10 20 30 0 0]
```

3表示拷贝成功的元素个数.bre

##### 11.把切片传递给函数

函数间传递切片就是要在函数间以值的方式传递切片。由于切片的尺寸很小，在函数间复制和传递切片成本也很低。
让我们创建一个包含 100 万个整数的切片，并将这个切片以值的方式传递给函数 foo()：

```go
myNum := make([]int, 1e6)
// 将 myNum 传递到函数 foo()
slice = foo(myNum)
// 函数 foo() 接收一个整型切片，并返回这个切片
func foo(slice []int) []int {
...
return slice
} 
```

