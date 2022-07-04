package semaphore_study

import (
	"fmt"
	"github.com/marusama/semaphore/v2"
)

func SemaphoreStudy() {
	//ctx := context.Background()
	sem := semaphore.New(5)
	for i := 0; i < 10; i++ {
		fmt.Println("+++++++++", i)
		sem.Acquire(nil, 1)
		sem.Release(1)
	}

	fmt.Println("-------结束")
}
