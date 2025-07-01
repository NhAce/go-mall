package main

import "fmt"

func main() {
	res1 := add(1,2)
	fmt.Println("1+2 = ",res1)
	res2 := plusPlus(1,2,3)
	fmt.Println("1+2+3 = ",res2)

	swap1,swap2 := swap(1,2)
	fmt.Println("swap1 = ",swap1,",swap2 = ",swap2)
	//如果你仅仅需要返回值的一部分的话，你可以使用空白标识符 _。
	_,swap3 := swap(1,2)
	fmt.Println("swap3 = ",swap3)

	res3 := add2(1,2,3,4,5)
	fmt.Println("1+2+3+4+5 = ",res3)

	nums := []int{1,2,3,4,5}
	res4 := add2(nums...)
	fmt.Println("1+2+3+4+5 = ",res4)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())

	nextInt3 := intSeq()
	fmt.Println(nextInt3())
	fmt.Println(nextInt3())

	fmt.Println(fact(7))

	var fib func(n int) int
	fib = func(n int) int{
		if n < 2{
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
//这里是一个函数，接受两个 int 并且以 int 返回它们的和
func add(a int,b int) int{

	//Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值
	return a + b
}

//当多个连续的参数为同样类型时， 可以仅声明最后一个参数的类型，忽略之前相同类型参数的类型声明
func plusPlus(a,b,c int) int{
	return a + b + c
}


//多值返回
func swap(a int,b int) (int,int){
	return b,a
}

//变参函数
func add2(nums ...int) int{
	sum := 0
	for _,num := range nums{
		sum += num
	}
	return sum
}

//闭包 intSeq 函数返回一个在其函数体内定义的匿名函数。 返回的函数使用闭包的方式 隐藏 变量 i。 返回的函数 隐藏 变量 i 以形成闭包
func intSeq() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}

//递归
func fact(n int) int{
	if n == 0{
		return 1
	}
	return n * fact(n-1)
}
