package main
import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["apple"] = 1
	map1["banana"] = 2
	fmt.Println(map1)
	//当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略
	_, ok := map1["apple2"]
	if !ok {
		fmt.Println("apple2 not found")
	}else{
		fmt.Println("apple2 found, value is ", map1["apple2"])
	}
}