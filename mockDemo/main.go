package main

import "github.com/yleoer/go-demo/mockDemo/equipment"

func main() {
	phone := equipment.NewIphone6s()
	xiaoMing := NewPerson("xiaoMing", phone)
	xiaoMing.dayLife()
}
