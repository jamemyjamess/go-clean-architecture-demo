package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math"
// 	"math/rand"
// 	"runtime"
// 	"sync"
// 	"time"

// 	"golang.org/x/sync/errgroup"
// )

// //	go_leak "github.com/jamemyjamess/go-clean-architecture-demo/cmd/go-leak"

// type Test struct {
// 	no    int
// 	value int
// }

// var m sync.Mutex

// func main() {
// 	TestErrorGroup()
// 	// go_leak.TestGoLeak()
// }

// func TestErrorGroup() {
// 	var m1, m2 runtime.MemStats
// 	runtime.GC()
// 	runtime.ReadMemStats(&m1)
// 	fmt.Println("m1.Mallocs:", m1.Mallocs)
// 	errorGroupWithTimeOut()
// 	runtime.ReadMemStats(&m2)
// 	fmt.Println("m2.Mallocs:", m2.Mallocs)
// 	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
// 	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)

// 	for {
// 		time.Sleep(5 * time.Second)
// 		runtime.ReadMemStats(&m2)
// 		fmt.Println("m2.Mallocs:", m2.Mallocs)
// 		fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
// 		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)

// 	}

// }

// func errorGroup() {
// 	ctx := context.Background()
// 	testList := &[]Test{}
// 	rows := 100000
// 	limit := 2000
// 	section := int(math.Ceil(float64(rows) / float64(limit)))
// 	g, ctx := errgroup.WithContext(ctx)

// 	for i := 0; i < section; i++ {
// 		// log.Println("start section:", i)
// 		no := i
// 		offset := no * limit
// 		fun := func() error {
// 			return doWorkErrorGroup(ctx, testList, no, limit, offset)
// 		}
// 		g.Go(fun)
// 	}
// 	// log.Println("pending g.Wait()")
// 	if err := g.Wait(); err != nil {
// 		log.Println(err.Error())
// 	}
// 	//time.Sleep(10 * time.Second)
// 	fmt.Println("total test: ", len(*testList))
// }

// func doWorkErrorGroup(ctx context.Context, testList *[]Test, no, limit, offset int) error {
// 	err := error(nil)
// 	testListItem := &[]Test{}
// 	for i := 0; i < limit; i++ {
// 		testItem := &Test{}
// 		testItem.no = i + no
// 		testItem.value = rand.Intn(100)
// 		*testListItem = append(*testListItem, *testItem)
// 	}
// 	m.Lock()
// 	*testList = append(*testList, *testListItem...)
// 	m.Unlock()
// 	return err
// }

// func errorGroupWithTimeOut() {
// 	testList := &[]Test{}
// 	rows := 10000000
// 	limit := 2000
// 	section := int(math.Ceil(float64(rows) / float64(limit)))
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	// Even though ctx will be expired, it is good practice to call its
// 	// cancelation function in any case. Failure to do so may keep the
// 	// context and its parent alive longer than necessary.
// 	// prevent memory leak
// 	defer cancel()
// 	g, ctx := errgroup.WithContext(ctx)
// 	for i := 0; i < section; i++ {
// 		no := i
// 		offset := no * limit
// 		fn := func() error {
// 			return doWorkErrorGroupWithTimeOut(ctx, testList, no, limit, offset)
// 		}
// 		g.Go(fn)
// 	}
// 	if err := g.Wait(); err != nil {
// 		log.Println("g.Wait() err:", err.Error())
// 	}
// 	log.Println("total test: ", len(*testList))

// }

// func doWorkErrorGroupWithTimeOut(ctx context.Context, testList *[]Test, no, limit, offset int) error {
// 	testListItem := &[]Test{}
// 	done := make(chan struct{}, 1)
// 	defer close(done)
// 	// example error
// 	// randInt := rand.Intn(2)
// 	// if randInt == 1 {
// 	// 	return fmt.Errorf("example error on work no: %v", no)
// 	// }
// 	for i := 0; i < limit; i++ {
// 		testItem := &Test{
// 			no:    i + no,
// 			value: rand.Intn(100),
// 		}
// 		*testListItem = append(*testListItem, *testItem)
// 	}
// 	m.Lock()
// 	*testList = append(*testList, *testListItem...)
// 	m.Unlock()
// 	select {
// 	case <-ctx.Done():
// 		return ctx.Err()
// 	case done <- struct{}{}:
// 		// or close(done) on this
// 		return nil
// 	}
// }
