package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
本次作业相对第一次还是比较简单的，请老师查阅
*/

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c *Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func useInterface(s Shape) {
	fmt.Println("面积是: ", s.area())
	fmt.Println("周长是: ", s.perimeter())
}

type Person struct {
	name string
	age  int
}

type Employee struct {
	EmployeeID int
	Person
}

func (e *Employee) PrintInfo() {
	fmt.Println("Name:", e.name, "Age:", e.age, "EmployeeID:", e.EmployeeID)
}

type SafeCounter struct {
	l   sync.Mutex
	num int
}

func (n *SafeCounter) addNum() {
	n.l.Lock()
	defer n.l.Unlock()
	n.num++
}
func (n *SafeCounter) getNum() int {
	n.l.Lock()
	defer n.l.Unlock()
	return n.num
}

type SafeCounter2 struct {
	num int64
}

func (n *SafeCounter2) addNum() {
	atomic.AddInt64(&n.num, 1)
}

func (n *SafeCounter2) getNum() int64 {
	return atomic.LoadInt64(&n.num)
}

func main() {
	/*
		1、题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
		在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
		考察点 ：指针的使用、值传递与引用传递的区别。
	*/
	num := 10
	addten(&num) //+10
	fmt.Println(num)

	/*
		2、实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
		考察点 ：指针运算、切片操作。
	*/
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubleNum(&arr)
	fmt.Println(arr)

	/*
		3、编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
		考察点 ： go 关键字的使用、协程的并发执行。
	*/
	go func() {
		for i := 1; i <= 10; i++ { //输出偶数
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ { //输出奇数
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	/*
		定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
		然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
		在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
		考察点 ：接口的定义与实现、面向对象编程风格。
	*/
	shape := Shape(&Rectangle{10, 20})
	fmt.Println("矩形:")
	useInterface(shape)

	shape = Shape(&Circle{10})
	fmt.Println("圆形:")
	useInterface(shape)

	/*
		使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
		考察点 ：组合的使用、方法接收者。
	*/

	e := &Employee{
		Person:     Person{"张三", 18},
		EmployeeID: 1001,
	}
	e.PrintInfo()

	/*
		编写一个程序，使用通道实现两个协程之间的通信。
		一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
		考察点 ：通道的基本使用、协程间通信。

		实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
		考察点 ：通道的缓冲机制。
	*/

	ch := make(chan int, 3)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			select {
			case x := <-ch:
				fmt.Println(x)
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	time.Sleep(time.Second * 2)

	/*
		共享计数器
	*/
	// counter := &SafeCounter{}//通过锁的方式
	counter := &SafeCounter2{} //通过原子性的方式
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.addNum()
			}
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(counter.getNum())

}

func addten(num *int) {
	*num += 10
}

func doubleNum(arr *[]int) {
	for i := range *arr {
		(*arr)[i] *= 2
	}
}
