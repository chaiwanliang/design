package Adapter

import "fmt"

type Target interface {
	Execute()
}

type Adapted struct {

}

func (a *Adapted)SpecificExecute()  {
	fmt.Println("充电...")
}

type Adapter struct {
	*Adapted
}

func (a *Adapter)Execute()  {
	a.SpecificExecute()
}