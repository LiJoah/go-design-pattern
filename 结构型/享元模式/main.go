package main

import (
	"fmt"
	"GO-Design-Pattern/结构型/享元模式/flyweight"
)

func main(){
	fly := flyweight.NewFlyweight()
	base := flyweight.NewPeopleBase()
	fly.SetElement("PeopleBase",base)

	people_1 := fly.GetElement("PeopleBase").(flyweight.IProperty)
	people_2 := fly.GetElement("PeopleBase").(flyweight.IProperty)
	people_1 = flyweight.NewHelmet(people_1,10,10)
	people_2 = flyweight.NewHelmet(people_2,100,100)

	hp_1 := people_1.GetHPLimit()
	mp_1 := people_1.GetMPLimit()

	hp_2 := people_2.GetHPLimit()
	mp_2 := people_2.GetMPLimit()

	fmt.Printf("People_1:\nHP:%d\nMP:%d\n",hp_1,mp_1)
	fmt.Printf("People_2:\nHP:%d\nMP:%d\n",hp_2,mp_2)
}

// 通过共享技术来实现大量细粒度对象的复用
