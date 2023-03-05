package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println("defer")
	fmt.Println("sum", sum(2, 3))
	deferValues()

	go showNumbers(100)
	runtime.Gosched()
	// the way to switch between goroutines

	go showNumbers(50)
	time.Sleep(time.Second)
	// in the code above, i set time sleep for first goroutine, therefore GO scheduler decides to execute another goroutine which prints numbers while the first is sleeping

	makePanic()

	fmt.Println("exit")
}

func showNumbers(num int) {
	fmt.Println("start =========== ")
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
	fmt.Println("done =========== ")
}

func deferValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("first deferred", i)
		//append to deferred stack and execute in reverse order
		//defer fmt.Println("first deferred", 0)
		//defer fmt.Println("first deferred", 1)
		//defer fmt.Println("first deferred", 2)
		//defer fmt.Println("first deferred", 3)
		//...
		//Output print:
		// first deferred 3
		// first deferred 2
		// first deferred 1
		// first deferred 0
	}

	//bad output
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("second deferred", i)
			//Output print:
			// second deferred 10
			// second deferred 10
			// second deferred 10
			// second deferred 10
			//...
		}()
	}

	// the ways to fix behavior as above

	// first one:
	for i := 0; i < 10; i++ {
		k := i
		//create new local variable and use it inside defer func

		defer func() {
			fmt.Println("third deferred", k)
			//Output print:
			// third deferred 3
			// third deferred 2
			// third deferred 1
			// third deferred 0
		}()
	}

	//second one:
	for i := 0; i < 10; i++ {
		defer func(i int) {
			// add argument to defer func

			fmt.Println("fourth deferred", i)
			//Output print:
			// fourth deferred 3
			// fourth deferred 2
			// fourth deferred 1
			// fourth deferred 0
		}(i)
	}
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y
	return
}

// handle panic
func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	panic("some panic")
}
