package AbstractFactory

import . "GO-Design-Pattern/创建型/工厂模式/Factories"

/***
 *	Abstract Factory
 */

type HJXFactory interface {
	CreateFood() Food
	CreateDrink() Drink
}

type FactoryA struct {
}

func (af FactoryA) CreateFood() Food {
	f := Meat{}
	return f
}

func (af FactoryA) CreateDrink() Drink {
	d := CoCo{}
	return d
}

type FactoryB struct {
}

func (bf FactoryB) CreateFood() Food {
	f := Hamberger{}
	return f
}

func (bf FactoryB) CreateDrink() Drink {
	d := Tea{}
	return d
}
