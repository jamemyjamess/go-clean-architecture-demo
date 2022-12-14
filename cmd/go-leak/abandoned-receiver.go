package go_leak

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

var bufferChAmount = 1000000

func Test() {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	fmt.Println("m1.Mallocs:", m1.Mallocs)
	handler()
	runtime.ReadMemStats(&m2)
	fmt.Println("m2.Mallocs:", m2.Mallocs)
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
	for {
		time.Sleep(5 * time.Second)
		runtime.ReadMemStats(&m2)
		fmt.Println("m2.Mallocs:", m2.Mallocs)
		fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)

	}
}

func TestGoLeak() {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	fmt.Println("m1.Mallocs:", m1.Mallocs)
	handler()
	runtime.ReadMemStats(&m2)
	fmt.Println("m2.Mallocs:", m2.Mallocs)
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
	// fmt.Println("MaxRSS:", cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss)
	for {
		time.Sleep(5 * time.Second)
		runtime.ReadMemStats(&m2)
		fmt.Println("m2.Mallocs:", m2.Mallocs)
		fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)

	}
}

type Testing struct {
	No    int
	Value int
}

func handler() {
	ch := make(chan int, bufferChAmount)
	resultList := &[]Testing{}
	abandonedReceiver(ch, resultList)
	log.Println(len(*resultList))
}

func abandonedReceiver(ch chan int, resultList *[]Testing) {

	for i := 0; i < bufferChAmount; i++ {
		go recive(ch)
	}

	result := 0
	for data := range ch {
		result += data
		testItem := &Testing{}
		testItem.No = rand.Intn(9)
		testItem.Value = data
		*resultList = append(*resultList, *testItem)
		if result == bufferChAmount {
			close(ch)
		}
	}

	log.Println("success", result)

}

func recive(ch chan int) {
	ch <- 1
}
