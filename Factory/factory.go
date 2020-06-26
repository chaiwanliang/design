package Factory

import "fmt"

type Restaurant interface {
	GetFood()
}

type  DongLaiShun struct {

}

func (d *DongLaiShun)GetFood(){
	fmt.Println("DongLaiShun的食物已经生成完毕，就绪...")
}

type QuinFang struct {

}

func (q *QuinFang)GetFood()  {
	fmt.Println("QuinFang的食物已经生成完毕，就绪...")
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "d":
		return &DongLaiShun{}
	case "q":
		return &QuinFang{}
	}
	return nil
}
