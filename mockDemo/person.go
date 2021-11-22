package main

import (
	"fmt"

	"github.com/yleoer/go-demo/mockDemo/equipment"
)

type Person struct {
	name string
	phone equipment.Phone
}

func NewPerson(name string, phone equipment.Phone) *Person {
	return &Person{
		name: name,
		phone: phone,
	}
}

func (p *Person) goSleep() {
	fmt.Printf("%s go to sleep!", p.name)
}

func (p *Person) dayLife() bool {
	fmt.Printf("%s's daily life:\n", p.name)
	if p.phone.WeiXin() && p.phone.WangZhe() && p.phone.ZhiHu() {
		p.goSleep()
		return true
	}
	return false
}