package Singleton

import "sync"

var (
	p *Pet
	once sync.Once
)

// 在初始化的时候构造了对象
func init()  {
	// 单例模式中最重要的
	once.Do(func() {
		p=&Pet{}
	})
}

func GetInstance() *Pet{
	return p
}

type Pet struct {
	name string
	age int
	mux sync.Mutex
}

func (p *Pet)SetName(n string)  {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.name=n
}

func (p *Pet) GetName()string{
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *Pet)IncrementAge(){
	p.mux.Lock()
	defer p.mux.Unlock()
	p.age++
}

func (p *Pet)GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}