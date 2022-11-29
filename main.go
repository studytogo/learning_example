package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"learning_example/cmd"
	"learning_example/kafka_example"
	"learning_example/nats_study"
	"os"
	"time"
)

func main() {
	//TestKafka()
	//TestExecute()
	//StudyGoDatastructuresQueue()
	//kingpin_study.PcKingpin()
	//bpmn_engine.StudyBpmnEngine()
	//grpc_study.StudyGrpc()
	//WriteFileStudy()
	//fasttemplate.StudyFastTemplate()
	//docker_study.StudyDocker()
	//k8s_study.GetNodeResource()
	//minio_study.StudyMinio()
	//postgres_db.PostgresDBStudy()
	//err_group.StudyErrGroup()

	//semaphore_study.SemaphoreStudy()
	//goque_study.GoqueStudy()
	//gocron.StudyGocron()
	nats_study.NatsStudy()
	//kafka_example.KafkaReciveMessage()
	//bpmn_engine.StudyBpmnEngine()
	//goque_study.GoqueNormal()
}

////go:embed templates
//var tmpl embed.FS
//
////go:embed static
//var static embed.FS
//
//func StudyEmbed() {
//	app := fiber.New()
//	app.Use("/", filesystem.New(filesystem.Config{
//		Root: http.FS(static),
//	}))
//	app.Listen(":3000")
//}

func WriteFileStudy() {
	var data []byte
	err := os.WriteFile("test", data, 0666)
	if err != nil {
		fmt.Println("--------", err)
	}
	fmt.Println("++++", string(data))
}

func TestExecute() {
	cmd.Execute()
}

func TestKafka() {
	kafka_example.KafkaExample()
}

func StudyGoDatastructuresQueue() {
	queues := queue.New(100)
	queues.Put("1")
	queues.Put(5)

	for i := 0; i < 5; i++ {
		dispose := queues.Dispose()
		fmt.Println("OOOOOO", dispose)
		poll, err := queues.Poll(1, 5*time.Second)
		if err != nil {
			fmt.Println("------", err)
		}

		fmt.Println("+++", poll)
	}

}
