package demoMicroservices

import (
	"fmt"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int)
	atomic.AddInt32(&sum, int32(n))
	fmt.Println("run with ", n)
}

func demoFunc() {
	time.Sleep()
}
