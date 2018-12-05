package singleton

import "sync"

//	定义一个对象类型，该类型只有唯一的对象
type Singleton map[string]string

var (
	// 只执行一次动作的对象
	once sync.Once //	保证一个函数只运行一次
	instance Singleton //	被确保唯一的对象
)

func GetInstance() Singleton {
	// 当且仅当第一次被调用时才执行函数 f,
	fn := func() {
		instance = make(Singleton)
	}
	once.Do(fn)
	return instance
}
