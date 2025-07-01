package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	//range 在数组和 slice 中迭代。
	for _,num := range nums{
		sum += num
	}
	fmt.Println(sum)
	//range 在数组和 slice 中提供对每项的索引和值的访问。 上面我们不需要索引，所以我们使用 空白标识符 _ 将其忽略。 实际上，我们有时候是需要这个索引的。
	for i,num := range nums{
		if num == 3{
			fmt.Println("nums[",i,"] = ",num)
		}
	}
	kvs := map[string]int{
		"apple": 1,
		"banana": 2,
	}
	for k,v := range kvs{
		fmt.Println("key = ",k,",value = ",v)
	}
	//range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置，然后第二个是字符本身
	for i,c := range "go"{
		fmt.Println(i,c)
	}
}