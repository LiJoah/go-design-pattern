package 创建者模式

import (
	"fmt"
	"errors"
)

type Color string
type LampStatus bool
type Brand string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

const (
	OppleBulb Brand = "OPPLE"
	Osram           = "OSRAM"
)

//	Lamp Builder define
type Builder interface {
	Color(Color) LampBuilder
	Brand(Brand) LampBuilder
	Build() LampOperation
}

type LampBuilder struct {
	Lamp
}

func (lb LampBuilder) Color(c Color) LampBuilder {
	lb.color = c
	return lb
}

func (lb LampBuilder) Brand(b Brand) LampBuilder {
	lb.brand = b
	return lb
}

func (lb LampBuilder) Build() LampOperation {
	lamp := Lamp{color: lb.color, brand: lb.brand, status: false}
	return lamp
}

func NewBuilder() Builder {
	return LampBuilder{}
}

type LampOperation interface {
	Open() error
	Close() error
	ProductionIllustrative()
}

type Lamp struct {
	color  Color
	brand  Brand
	status LampStatus
}

func (l Lamp) Open() error {
	if l.status {
		return errors.New("Lamp is opened")
	}
	fmt.Println("Open lamp.")
	l.status = true
	return nil
}

func (l Lamp) Close() error {
	if !l.status {
		return errors.New("Lamp is closed")
	}
	fmt.Println("Close lamp.")
	l.status = true
	return nil
}

func (l Lamp) ProductionIllustrative() {
	fmt.Println("I'm a lamp.")
	fmt.Println("Color:" + l.color)
	fmt.Println("Brand:" + l.brand)
}

func main() {
	b := NewBuilder()
	lamp_1 := b.Color(BlueColor).Brand(Osram).Build()
	lamp_1.Open()
	lamp_1.ProductionIllustrative()

	lamp_2 := b.Color(GreenColor).Brand(OppleBulb).Build()
	lamp_2.ProductionIllustrative()
}

// 将一个复杂的对象的构建与它的表示分离， 使得同样的构建过程可以创建不同的表示，
// 建造者模式是一步一步构建一个复杂的对象，
// 它允许用户通过指定复杂的对象的类型和内容就可以构建， 用户不需要知道内部的
// 具体的构建细节，建造者模式又称为生成器模式