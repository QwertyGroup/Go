package main

import "fmt"

func main() {
	orsen := orsen{Name: "Orsen"}
	orsen.sayName()
	do(&orsen)
	do2(orsen)
}

func do(h human) {
	h.sayName()
	h.sayLast()
}

func do2(sl sayLaster) {
	sl.sayLast()
}

type orsen struct {
	Name string
}

type nameSayer interface {
	sayName()
}

type sayLaster interface {
	sayLast()
}

type human interface {
	nameSayer
	sayLaster
}

func (o *orsen) sayName() {
	fmt.Println(o.Name)
}

func (o orsen) sayLast() {
	fmt.Println("last " + o.Name)
}

// // invalid receiver type nameSayer (nameSayer is an interface type)
// func (ns nameSayer) gogo() {
// 	ns.sayName()
// }
