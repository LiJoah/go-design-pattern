package FactoryMethod

import . "GO-Design-Pattern/创建型/工厂模式/Factories"

/***
 *	Factory Method
 */
type Factory interface {
	Create() Food
}

type MeatFactory struct {
}

func (mf MeatFactory) Create() Food {
	m := Meat{}
	return m
}

type HambergerFactory struct{

}

func (hf HambergerFactory) Create() Food {
	h := Hamberger{}
	return h
}