package main

import (
	"context"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

//	go_leak "github.com/jamemyjamess/go-clean-architecture-demo/cmd/go-leak"

type Test struct {
	no    int
	value int
}

var m sync.Mutex

func main() {
	TestErrorGroup()
	// go_leak.TestGoLeak()
}

func TestErrorGroup() {
	// var m1, m2 runtime.MemStats
	// runtime.GC()
	// runtime.ReadMemStats(&m1)
	// fmt.Println("m1.Mallocs:", m1.Mallocs)
	errorGroupWithTimeOut()
	// runtime.ReadMemStats(&m2)
	// fmt.Println("m2.Mallocs:", m2.Mallocs)
	// fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	// fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)

	// for {
	// 	time.Sleep(5 * time.Second)
	// 	runtime.ReadMemStats(&m2)
	// 	fmt.Println("m2.Mallocs:", m2.Mallocs)
	// 	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	// 	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)

	// }

}

func errorGroupWithTimeOut() {
	rows := 10000000
	limit := 200
	section := int(math.Ceil(float64(rows) / float64(limit)))
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	// prevent memory leak
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	// set size of ch equal amount section (worker) is faster than size 1
	chTest := make(chan []Test, section)
	for i := 0; i < section; i++ {
		no := i
		offset := no * limit
		fn := func() error {
			return doWorkErrorGroupWithTimeOut(ctx, chTest, no, limit, offset, section)
		}
		g.Go(fn)
	}
	testListRes := []Test{}
	g.Go(func() error {
		for data := range chTest {
			testListRes = append(testListRes, data...)
			if len(testListRes) == rows {
				close(chTest)
			}
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Println("g.Wait() err:", err.Error())
	}
	log.Println("total test: ", len(testListRes))

}

func doWorkErrorGroupWithTimeOut(ctx context.Context, chTest chan<- []Test, no, limit, offset, section int) error {
	testListItem := &[]Test{}
	// example error
	// randInt := rand.Intn(2)
	// if randInt == 1 {
	// 	return fmt.Errorf("example error on work no: %v", no)
	// }
	for i := 0; i < limit; i++ {
		testItem := &Test{
			no:    i + no,
			value: rand.Intn(100),
		}
		*testListItem = append(*testListItem, *testItem)
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case chTest <- *testListItem:
		return nil
	}
}
