package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstancByDoubleCheck(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(gID int) {
			ins := GetInstancByDoubleCheck()
			fmt.Printf("groutine: %d, ins: %p\n", gID, ins)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
