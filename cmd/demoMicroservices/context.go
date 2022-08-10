package demoMicroservices

import (
	"sync"
	"time"
)

func processRequest(wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		print("Processing request : %v\n", total)
		total++
		time.Sleep(time.Second)

	}
	print("Processed %v requests\n", total)
	wg.Done()
}
