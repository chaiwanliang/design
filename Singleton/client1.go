package Singleton

func IncrementAge(){
	p :=GetInstance()
	p.IncrementAge()
}