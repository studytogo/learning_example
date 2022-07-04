package err_group

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func StudyErrGroup() {
	var eg errgroup.Group
	eg.SetLimit(0)
	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			fmt.Println("xxxxxxx")
			time.Sleep(5 * time.Second)
			return nil
		})
	}

	select {}
}
