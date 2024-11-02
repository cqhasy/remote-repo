# week2的go语言知识点：

## 1:比较字符串"hello，世界"的**长度** 和for range 该字符串的循环次数

- ```go
  a:="hello,世界"
  len(a)=12
  var b int=0
  for range a{
      b++
  }
  b==8
  for _, r := range a {
  		b = b + 1
  		fmt.Println(r)
  	}
  	fmt.Println(b)
  }
  b==8
  ```

- `len(a)` 返回的是**字符串的字节数**，而不是字符数。如果你想获得字符数（尤其是在包含非 ASCII 字符的情况下），可以使用 `utf8.RuneCountInString` 函数

- for range a 会按字符遍历而不是按字节遍历。

###  什么是 `rune`:

在 Go 中，`rune` 是一个类型别名，它代表一个 UTF-8 编码的 Unicode 字符。每个 `rune` 占用 4 个字节（32 位），可以表示大多数语言中的字符，包括汉字、表情符号等。

使用 `for _, r := range a` 可以确保你每次迭代获取的是一个完整的字符，而不是单个字节。这是因为 `range` 在遍历字符串时会自动处理 Unicode 编码，并返回每个字符的 `rune` 值。其中r代表rune.

## 2   x 的重新声明：

```go
  func main() {
      x := "hello!"
      for i := 0; i < len(x); i++ {
          x := x[i]
          if x != '!' {
              x := x + 'A' - 'a'
              fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
         }
     }
  }
```

在该代码中，x第一次声明是string型“hello”;之后又在for循环内部重新声明为uint32型，变为了对应的ASCLL码，随后又以字符形式打印。

## 3 匿名函数与闭包

- 匿名函数：即没有函数名的函数，通常用于实现回调函数和闭包。

  ```go
  func (参数)(返回值){
      函数体
  }
  例子：
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  	func() {
  		fmt.Println("hello world")
  	}()//两种执行方式：在后面直接执行。
  	a := func() {
  		fmt.Println("hello world")
  
  	}
  	a()//赋值给a,以a来执行。
  }
  ```

- 闭包：指的是函数和其引用的环境组合而成的实体，简单来说，闭包=函数＋引用环境。

  ```go
  例子：
  //定义一个函数，它的返回值是一个函数
  func a() func(){
      return func(){
          fmt.Println(Hello,October)
      }
  }
  func main () {
      r:=a()
      r()//相当与执行了a函数内部的匿名函数。这里只是调用了r()
  }
  若为下面这个例子：
  package main
  
  import "fmt"
  
  func a() func() string {
      name:="chenqihao"
      return func() string {
  		fmt.Println("Hello,October",name)
  		return "woaishiyue"
  	}
  }
  func main() {
  	r := a()
  
  	r()//这个只会输出Hello,October，因为在 Go 语言中，如果你调用一个返回值的函数但不保存或使用该返回值，Go 会忽略它。这 			//意味着即使函数 r() 返回一个值，如果你不将其存储在一个变量中或直接使用它，那么这个返回值就不会被处理。
  	b := r()//这里就会输出Hello,October
      				//woaishiyue
  	fmt.Println(r())//这里与b的输出相同。
  }//当然从这里我们也能看出来一个特点：闭包=函数+外层变量的引用，r()此时就是一个闭包。
  
  func a(name,string) func() string {
      return func() string {
  		fmt.Println("Hello,October",name)
  		return "woaishiyue"
  	}//效果与接受一个外层参数一致。
  ```

  ```go
  进阶：
  func a(l1,int)(func(int)int,func(int)int){
      fx:=func (l1,int)int{
          y1:=y1+l1
          return y1
      }
      gx:=func(l2.int)int{
          y2:=y2-l1
          return y2
      }
      return fx,gx
  }
  ```

  