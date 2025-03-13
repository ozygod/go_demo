package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var count uint32 = 0

func main() {
	for i := 0; i < 10; i++ {
		//go func(i uint32) {
		//	fn := func() {
		//		fmt.Println(i)
		//	}
		//	trigger(i, fn)
		//}(uint32(i))

		go trigger(uint32(i), func() {
			fmt.Println(i)
		})
	}

	trigger(10, func() {

	})
}

func trigger(i uint32, fn func()) {
	for {
		if n := atomic.LoadUint32(&count); n == i {
			fn()
			atomic.AddUint32(&count, 1)
			break
		}
		time.Sleep(time.Nanosecond)
	}
}
