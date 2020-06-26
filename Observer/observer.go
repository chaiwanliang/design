package Observer

import (
	"fmt"
	"sync"
	"time"
)

// 数据
type Event struct {
	Data int
}

// 观察者接口（行为）
type Observer interface {
	NotifyCallback(event Event)
}

// 主题(行为)
type Subject interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

// 实体
type eventObserver struct {
	Id int
	Time time.Time
}
type eventSubject struct {
	Observer sync.Map
}

func (e eventObserver)NotifyCallback(event Event){
	fmt.Printf("Recieved:%d after %v\n",event.Data,time.Since(e.Time))
}

func (e *eventSubject)AddListener(obs Observer) {
	e.Observer.Store(obs,struct{}{})
}

func (e *eventSubject)RemoveListener(obs Observer) {
	e.Observer.Delete(obs)
}

func (e *eventSubject)Notify(event Event) {
	e.Observer.Range(func(key, value interface{}) bool {
		if key==nil{
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}

// 斐波那契
func Fib(n int) chan int  {
	out :=make(chan int)

	go func() {
		defer close(out)
		for i,j:=0,1;i<n;i,j=i+j,i {
			out<-i
		}
	}()
	return out
}