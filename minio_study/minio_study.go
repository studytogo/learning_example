package minio_study

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func StudyMinio() {
	endpoint := "192.168.70.204:9002"
	accessKeyID := "mxsche"
	secretAccessKey := "mxsche123"
	//useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	//log.Printf("%#v\n", minioClient)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := minioClient.ListObjects(ctx, "sunmq", minio.ListObjectsOptions{})
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println("111111", object.Err)
			return
		}
		fmt.Println("3333333", object.Key)
		//	fmt.Println("2222222", fmt.Sprintf("%+v", object))
	}
}
