package test

import (
	"fmt"
	"math"
)

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
*/

type Shape interface {
	Area() float64 //面积

	Perimeter() float64 //周长
}

type Rectangle struct {
	width float64
	hight float64
}

func (r Rectangle) Area() float64 {
	return r.hight * r.width
}

func (r Rectangle) Perimeter() float64 {
	return (r.hight + r.width) * 2
}

type Circle struct {
	len float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.len * c.len
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.len * 2
}

func Test06() {
	r := Rectangle{
		width: 10,
		hight: 20,
	}
	fmt.Printf("长方体面积:%v\n", r.Area())
	fmt.Printf("长方体周长:%v\n", r.Perimeter())

	c := Circle{
		len: 4,
	}
	fmt.Printf("圆面积:%v\n", c.Area())
	fmt.Printf("圆面积:%v\n", c.Perimeter())

}
