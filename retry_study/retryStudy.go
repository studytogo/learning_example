package retry_study

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-retry"
	"time"
)

func StudyRetry() {
	i := 0
	b := retry.NewFibonacci(1 * time.Second)
	b = retry.WithMaxRetries(5, b) // 设置最大重试次数为5

	err := retry.Do(context.Background(), b, func(ctx context.Context) error {
		i++
		// 这里是你需要重试的操作
		// 如果操作失败，返回错误；如果操作成功，返回nil
		fmt.Println("尝试操作...", i)
		//return retry.RetryableError(fmt.Errorf("操作失败"))
		return fmt.Errorf("操作失败")
	})

	if err != nil {
		fmt.Println("重试操作失败:", err)
	} else {
		fmt.Println("重试操作成功")
	}

}
