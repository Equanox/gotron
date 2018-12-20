package runner

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRunner(t *testing.T) {

	do := make(chan string, 100)
	for i := 0; i < cap(do)+1; i++ {
		go func(i int) {
			do <- strconv.Itoa(i)
		}(i)
	}

	task := Go(func(stop StopChan, finish Finish) {
		for {
			select {
			case t := <-do:
				fmt.Println(t)
			case _, ok := <-stop:
				if !ok {
					finish()
					return
				}
			}
		}
	})

	task.Stop()
	err := task.Wait()
	if err != nil {
		t.Error(err)
	}
}
