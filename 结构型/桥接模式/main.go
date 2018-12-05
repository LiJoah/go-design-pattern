package main

import "GO-Design-Pattern/结构型/桥接模式/bridge"

func main(){
	// craete windows
	windows := new(bridge.Windows)
	windows.Boot()

	// create linux
	linux := new(bridge.Linux)
	linux.Boot()

	// Code util
	ide := bridge.CodeUtil{}

	// write programs
	sing := ide.GetSingProgram("Hello World")
	dog := ide.GetDogProgram()

	windows.Build(&sing)
	windows.Build(&dog)
	linux.Build(&sing)
	linux.Build(&dog)

	windows.ExecuteProgram("dog")
	windows.ExecuteProgram("sing")
	linux.ExecuteProgram("sing")
	linux.ExecuteProgram("dog")
}

// 桥接模式： 将继承关系转换为关联模式, 从而通过组合来达到目的
