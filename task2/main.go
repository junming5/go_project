package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func pointTest(a *int) {
	*a += 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func sliceTest(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func goroutineTest() {
	go func() {
		for i := 1; i <= 10; i++ {
			v := i % 2
			if v == 1 {
				fmt.Println("从1到10的奇数->", i)
			}
		}
	}()
	go func() {
		for i := 2; i <= 10; i++ {
			v := i % 2
			if v == 0 {
				fmt.Println("从2到10的偶数->")
			}
		}
	}()
	time.Sleep(time.Second)
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
func taskTest() map[string]time.Duration {
	scheduler := Scheduler{tasks: make([]func(taskName string), 0), taskMap: make(map[string]time.Duration)}
	scheduler.initTasks()
	scheduler.runTasks()
	return scheduler.taskMap
}

type Scheduler struct {
	tasks   []func(taskName string)
	taskMap map[string]time.Duration
}

func (s *Scheduler) initTasks() {
	for i := 0; i < 10; i++ {
		s.tasks = append(s.tasks, func(taskName string) {
			fmt.Println(taskName)
		})
	}
}

func (s *Scheduler) runTasks() {
	for i := 0; i < len(s.tasks); i++ {
		index := i
		go s.executeTask(index, s.tasks[i])
	}

	time.Sleep(5 * time.Second)
}
func (s *Scheduler) executeTask(index int, task func(taskName string)) {
	start := time.Now()
	str := fmt.Sprintf("task%d", index)

	task(str)
	duration := time.Since(start)
	s.taskMap[str] = duration
}

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
func interfaceTest() {
	r := Rectangle{}
	r.Area()
	r.Perimeter()
	c := Circle{}
	c.Area()
	c.Perimeter()
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r *Rectangle) Area() {
	fmt.Println("Rectangle Area()")
}

func (r *Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter()")
}

type Circle struct {
}

func (c *Circle) Area() {
	fmt.Println("Circle Area()")
}

func (c *Circle) Perimeter() {
	fmt.Println("Circle Perimeter()")
}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("EmployeeID=%s, name=%s, age=%d\n", e.EmployeeID, e.Name, e.Age)
}

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func channelTest() {
	ch := make(chan int)

	go func(ch chan<- int) {
		defer close(ch)

		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}(ch)

	go func(ch <-chan int) {
		for v := range ch {
			fmt.Printf("接收到整数为：%d\n", v)
		}
	}(ch)

	time.Sleep(time.Second)
}

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func channelTest2() {
	ch := make(chan int, 10)

	go func(ch chan<- int) {
		defer close(ch)

		for i := 1; i <= 100; i++ {
			ch <- i
		}
	}(ch)

	go func(ch <-chan int) {
		for v := range ch {
			fmt.Printf("接收到整数为：%d\n", v)
		}
	}(ch)

	time.Sleep(time.Second)
}

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
func lockTest() {
	s := SafeCounter{}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				s.Increment()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("Final count: %d\n", s.count)
}

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (s *SafeCounter) Increment() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
}

func (s *SafeCounter) GetCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func atomicTest() {
	var count int64 = 0
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("Final count: %d\n", count)
}

func main() {
	// p := 10
	// fmt.Println("old value = ", p)
	// pointTest(&p)
	// fmt.Println("new value = ", p)

	// nums := []int{1, 3, 5, 9}
	// fmt.Println("old slice = ", nums)
	// sliceTest(nums)
	// fmt.Println("new slice = ", nums)

	// goroutineTest()

	// fmt.Println(taskTest())

	// interfaceTest()

	// employee := Employee{
	// 	Person: Person{
	// 		Name: "zhangsan",
	// 		Age:  18,
	// 	},
	// 	EmployeeID: "10001",
	// }
	// employee.PrintInfo()

	// channelTest()

	// channelTest2()

	// lockTest()

	atomicTest()
}
