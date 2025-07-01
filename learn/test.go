package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	// map1 := make(map[string]int)
	// map1["apple"] = 1
	// map1["banana"] = 2
	// fmt.Println(map1)
	// //当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略
	// _, ok := map1["apple2"]
	// if !ok {
	// 	fmt.Println("apple2 not found")
	// }else{
	// 	fmt.Println("apple2 found, value is ", map1["apple2"])
	// }
	// nums := []int{1, 2, 3, 4, 5}
	// sum := 0
	// //range 在数组和 slice 中迭代。
	// for _,num := range nums{
	// 	sum += num
	// }
	// fmt.Println(sum)
	// //range 在数组和 slice 中提供对每项的索引和值的访问。 上面我们不需要索引，所以我们使用 空白标识符 _ 将其忽略。 实际上，我们有时候是需要这个索引的。
	// for i,num := range nums{
	// 	if num == 3{
	// 		fmt.Println("nums[",i,"] = ",num)
	// 	}
	// }
	// kvs := map[string]int{
	// 	"apple": 1,
	// 	"banana": 2,
	// }
	// for k,v := range kvs{
	// 	fmt.Println("key = ",k,",value = ",v)
	// }
	// //range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置，然后第二个是字符本身
	// for i,c := range "go"{
	// 	fmt.Println(i,c)
	// }
}	