package singleton

import "sync"

// Singleton 单列模式
type Singleton struct{}

var singleton *Singleton
var once sync.Once

// GetInstance 获取单列模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}