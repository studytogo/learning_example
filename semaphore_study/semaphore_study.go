package semaphore_study

import (
	"fmt"
	"github.com/marusama/semaphore/v2"
	"time"
)

func SemaphoreStudy() {
	//ctx := context.Background()
	sem := semaphore.New(1)
	//go func() {
	//	for i := 1; i < 5; i++ {
	//		sem.SetLimit(i)
	//		fmt.Println("-----------------")
	//		time.Sleep(5 * time.Second)
	//	}
	//}()
	for i := 0; i < 15; i++ {
		go func(n int) {
			sem.Acquire(nil, 1)
			fmt.Println("hahahahahahahaha", n)
			time.Sleep(5 * time.Second)
			sem.Release(1)
		}(i)
	}

	fmt.Println("-------结束")
	select {}
}
