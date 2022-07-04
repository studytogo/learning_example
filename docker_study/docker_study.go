package docker_study

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"time"
)

func StudyDocker() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	//containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println("------", containers)
	//
	//for _, container := range containers {
	//	fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	//}
	//inDirPattern := "*-mxsche-in"
	//outDirPattern := "*-mxsche-out"
	//dockerInMountDir := "/mnt/in"
	//dockerOutMountDir := "/mnt/out"
	//inDir, err := os.MkdirTemp("", inDirPattern)
	//if err != nil {
	//	fmt.Println("oooooooo11111111", err)
	//}
	//
	//outDir, err := os.MkdirTemp("", outDirPattern)
	//if err != nil {
	//	fmt.Println("oooooooo2222222222", err)
	//}
	//mounts := []mount.Mount{
	//	{Type: mount.TypeVolume, Source: `E:\vol`, Target: `/data`},
	//	//	{Type: mount.TypeBind, Source: outDir, Target: dockerOutMountDir},
	//}

	create, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: "192.168.70.202:32373/registry/las_to_3dtiles",
		Cmd:   []string{"./las_to_3dtiles", " -bucket=test", " -objectName='500112/20220613/01/las/20211229_01_37_04_000-09-43-31-756 - Cloud - Cloud.las'", "-ss=4524"},
	}, nil, nil, nil, "")

	if err != nil {
		fmt.Println("22222222", err)
	}

	err = cli.ContainerStart(context.Background(), create.ID, types.ContainerStartOptions{})

	if err != nil {
		fmt.Println("333333", err)
	}
	fmt.Println("container started")

	//打印程序选择了本机哪个端口去映射到mongo中
	resultC, errC := cli.ContainerWait(context.Background(), create.ID, "")
	select {
	case err := <-errC:
		fmt.Println("errCerrCerrCerrCerrC", err)
	case result := <-resultC:
		fmt.Println("resultCresultCresultCresultCresultC", result)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reader, err := cli.ContainerLogs(ctx, create.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true})

	if err != nil {
		fmt.Println("llllllllllllllllllllllll", err)
	}
	defer reader.Close()

	buffer, err := io.ReadAll(reader)

	if err != nil && err != io.EOF {
		fmt.Println("rrrrrrrrrrrrrrrrrrrrrr", err)
	}
	fmt.Println("outputoutputoutputoutput", string(buffer))

	time.Sleep(100 * time.Second)
	fmt.Println("killing container")

	err = cli.ContainerRemove(context.Background(), create.ID, types.ContainerRemoveOptions{
		Force: true,
	})

	if err != nil {
		fmt.Println("555555", err)
	}

}
