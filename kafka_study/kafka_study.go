package kafka_study

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var c sarama.Client

func newKafka2(i int) sarama.SyncProducer {
	fmt.Println("ppppp", i)
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	address := []string{"192.168.70.202:30092"}

	if c == nil {
		client, err := sarama.NewClient(address, config)
		if err != nil {
			fmt.Println("111", err)
		}
		c = client
	}

	producer, err := sarama.NewSyncProducerFromClient(c)
	if err != nil {
		fmt.Println("222", err)
		return nil
	}

	return producer
}

func newKafka(i int) sarama.SyncProducer {
	fmt.Println("ppppp", i)
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	address := []string{"192.168.70.202:30092"}

	producer, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return producer
}

func StudyKafak() {
	for i := 0; i < 5; i++ {
		k := newKafka2(i)
		if k != nil {
			fmt.Println("iiiii")
			k.Close()
		}
	}

	fmt.Println("--------执行结束---------")
}

func SubMessage() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	address := []string{"192.168.60.216:31090"}

	consumer, err := sarama.NewConsumer(address, config)
	if err != nil {
		fmt.Println("111111", err)
		return
	}

	topic := "C.P.MD.Instance.Update4"

	partitions, err := consumer.Partitions("C.P.MD.Instance.Update4")
	if err != nil {
		fmt.Println("22222", err)
		return
	}

	for partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()

		// 处理消息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d, Offset: %d, Key: %s, Value: %s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}

	fmt.Println("-----------------------")

	select {}
}
