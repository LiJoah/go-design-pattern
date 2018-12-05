package 工厂模式

import "GO-Design-Pattern/创建型/工厂模式/Factories/AbstractFactory"

func main() {
	// Abstract Factory
	fa := AbstractFactory.FactoryA{}
	fa.CreateFood().Eat()
	fa.CreateDrink().Drink()

	fb := AbstractFactory.FactoryB{}
	fb.CreateFood().Eat()
	fb.CreateDrink().Drink()
}
