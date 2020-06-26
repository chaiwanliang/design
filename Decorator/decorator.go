package Decorator

import (
	"log"
	"math"
	"time"
)

// 定义函数的签名
type piFunc func(int) float64

// 装饰器的结构体
type Wrap struct {
	Id int
}

// 装饰器的实现
func (w *Wrap)WrapLogger(fun piFunc,logger *log.Logger) piFunc {
	return func(n int) float64 {
		fn := func(n int) (result float64){
			defer func(t time.Time) {
				logger.Printf("took=%v,n=%v,result=%v,id=%d",time.Since(t),n,result,w.Id)
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}
}

// 函数功能的具体实现
func Pi(n int) float64{
	ch :=make(chan float64)

	for k:=0;k<n;k++{
		go func(ch chan float64,k float64) {
			ch<-4*math.Pow(-1,k)/(2*k+1)
		}(ch,float64(k))
	}

	result :=0.0
	for i:=0;i<n ;i++  {
		result+=<-ch
	}
	return result
}

