package main

import (
	"GO-Design-Pattern/创建型/对象池模式/objectpool"
	"sync"
	"fmt"
	"strconv"
)

func fn() {
	p := pool.NewPool(5)
	wait := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		index := i
		wait.Add(1)
		go func(pool pool.Pool, ind int) {
			select {
			case Obj := <-pool:
				Obj.Do(ind)
				pool <- Obj
			default:
				fmt.Println("No Object:" + strconv.Itoa(ind))
			}
			wait.Done()
		}(*p, index)
	}
	wait.Wait()

}

func test() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
func main() {
	//test()
	fn()
}
