package main

import (
	"fmt"

	"github.com/yleoer/go-demo/productrepo"
)

func Foo() {
	fmt.Println("print 1")
	defer fmt.Println("print 2")
	fmt.Println("print 3")
}

func f(a ...int) {
	fmt.Printf("%#v\n", a)
	fmt.Println(a == nil)
}

func main() {
	env := "aliCloud"
	repo := productrepo.New(env)
	repo.StoreProduct("Hua", 104)
}
