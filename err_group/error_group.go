package err_group

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func StudyErrGroup() {
	var eg errgroup.Group
	eg.SetLimit(2)
	cancel, cancelFunc := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		cancelFunc()
	}()
	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			if cancel.Err() != nil {
				return nil
			}
			fmt.Println("---", i)
			time.Sleep(1 * time.Second)
			return errors.New("yyyy")
			//rand.Seed(time.Now().UnixNano())
			//intn := rand.Intn(3)
			//fmt.Println(intn)
			//if intn == 2 {
			//	return errors.New("hhhhhh")
			//} else if intn == 0 {
			//	for {
			//		fmt.Println("在循环")
			//		time.Sleep(1 * time.Second)
			//	}
			//}
			//return nil
		})
		fmt.Println("===", i)
	}

	errrr := eg.Wait()
	fmt.Println("hhh", errrr)
}
