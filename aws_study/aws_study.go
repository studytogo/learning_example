package aws_study

import (
	"bsdklib/store"
	"bsdklib/store/aws"
	"fmt"
	"strings"
)

func AwsGet() {
	S3ConfigAddress := "192.168.70.202:32709"
	S3ConfigAccessKey := "minio"
	S3ConfigSecretKey := "minio123"
	S3ConfigUseSSL := true
	bucket := "test"
	getPath := "/z_config/z_test/mapdeck调度前端yaml设计.txt"
	client := aws.NewClient(S3ConfigAddress, S3ConfigAccessKey,
		S3ConfigSecretKey, store.UseSSL(S3ConfigUseSSL))
	b, err := client.Get(getPath, store.GetBucket(bucket))
	if err != nil {
		fmt.Println("eeee", err)
	}
	split := strings.Split(strings.ReplaceAll(strings.Trim(string(b), "\r\n"), "\r\n", "\n"), "\n")

	for _, v := range split {
		fmt.Println("--------", v)
	}
}

func AwsGetList() {
	S3ConfigAddress := "192.168.70.202:32709"
	S3ConfigAccessKey := "minio"
	S3ConfigSecretKey := "minio123"
	S3ConfigUseSSL := true
	bucket := "test"
	getPath := "z_config/z_test/"
	fmt.Println("-----------", getPath)
	client := aws.NewClient(S3ConfigAddress, S3ConfigAccessKey,
		S3ConfigSecretKey, store.UseSSL(S3ConfigUseSSL))
	b, err := client.List(getPath, store.ListBucket(bucket))
	if err != nil {
		fmt.Println("eeee", err)
	}

	fmt.Println("------------", b)
}
