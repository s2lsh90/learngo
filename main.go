package main

import (
	"fmt"
	"learngo/something"
	"strings"
)

func main() {
	fmt.Printf("hello World")
	something.SayHello()
	//something.sayBye()

	const name string = " nico"
	//name = "sanghun"
	fmt.Println(name)
	//거의 사용 하지 않음

	//var names string = "nico"
	names := "nico"
	names = "hun"
	fmt.Println(names)

	totalLenght, upperName := lenAndUpper("hun")
	fmt.Println(totalLenght, upperName)

	totalLenght2, _ := lenAndUpper("hi")
	fmt.Println(totalLenght2, "a")

	fmt.Println("Test")

	repeatMe("nico", "lynn", "dal", "marl")
	totallenght3, uppercase := lendAndUpper("sanghun")
	fmt.Println(totallenght3, uppercase)

	total := superAdd(1, 2, 3, 4, 5)
	println(total)

	fmt.Println(canIDrink(16))
}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lendAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("i'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}
