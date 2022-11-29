package semaphore_study

import (
	"fmt"
	"github.com/marusama/semaphore/v2"
	"time"
)

func SemaphoreStudy() {
	//ctx := context.Background()
	sem := semaphore.New(0)
	for i := 0; i < 10; i++ {
		sem.Acquire(nil, 1)
		go func() {
			fmt.Println("hahahahahahahaha")
			time.Sleep(10 * time.Second)
		}()
		sem.Release(1)
	}

	fmt.Println("-------结束")
}
