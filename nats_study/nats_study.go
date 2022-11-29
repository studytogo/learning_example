package nats_study

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func NatsStudy() {
	//pubMessage()
	subMessage()
	select {}
}
func pubMessage() {
	testAddress := "nats://192.168.179.131:32172"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		fmt.Println("111", err)
	}
	js, err := nc.JetStream()
	//if err != nil {
	//	fmt.Println("222", err)
	//}
	//js.DeleteConsumer("ORDERS", "MONITOR")
	//js.DeleteStream("ORDERS")
	//
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS111",
		Subjects: []string{"ORDERS.scratch111"},
	})
	if err != nil {
		fmt.Println("aaaaa", err)
	}
	//
	//js.UpdateStream(&nats.StreamConfig{
	//	Name:     "ORDERS",
	//	MaxBytes: 8,
	//})
	//
	//js.AddConsumer("ORDERS", &nats.ConsumerConfig{
	//	Durable: "MONITOR",
	//})
	for i := 0; i < 500; i++ {
		p, err := js.Publish("ORDERS.scratch111", []byte("hello"))
		if err != nil {
			fmt.Println("ppppppp", err)
		}
		fmt.Println("iiiii", p.Stream)
		time.Sleep(1 * time.Second)
	}

}

func subMessage() {
	//js.Subscribe("ORDERS.scratch", func(msg *nats.Msg) {
	//	fmt.Printf("Received a JetStream message: %s\n", string(msg.Data))
	//	msg.Ack()
	//})
	testAddress := "nats://192.168.179.131:32172"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		fmt.Println("111", err)
	}
	js, err := nc.JetStream()
	if err != nil {
		fmt.Println("222", err)
	}
	//js.DeleteConsumer("ORDERS", "MONITOR")
	//js.DeleteStream("ORDERS")

	//js.AddStream(&nats.StreamConfig{
	//	Name:     "ORDERS",
	//	Subjects: []string{"ORDERS.scratch"},
	//})
	//
	//js.UpdateStream(&nats.StreamConfig{
	//	Name:     "ORDERS",
	//	MaxBytes: 8,
	//})
	//
	//js.AddConsumer("ORDERS", &nats.ConsumerConfig{
	//	Durable: "MONITOR",
	//})

	sub, err := js.PullSubscribe("ORDERS.scratch111", "MONITOR", nats.PullMaxWaiting(128))
	if err != nil {
		fmt.Println("sssss", err)
		time.Sleep(10 * time.Second)
	}

	for {
		f, err := sub.Fetch(1)
		if err != nil {
			fmt.Println("fffffff", err, "llll", sub.Subject)
		}
		for _, m := range f {
			fmt.Println("---------------", string(m.Data))
			m.Ack()
		}
	}

	//js.Subscribe("ORDERS.*", func(m *nats.Msg) {
	//	fmt.Println("-------------------------")
	//	//js.Subscribe("ORDERS.scratch", func(m *nats.Msg) {
	//	fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	//	m.Ack()
	//})
}

func SubMessage() {
	testAddress := "nats://192.168.179.131:32172"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		fmt.Println("33", err)
	}
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		fmt.Println("222", err)
	}
	js.AddConsumer("ORDERS", &nats.ConsumerConfig{
		Durable: "MONITOR",
	})
	fmt.Println("------------接收消息")
	//subMessage(js)

	select {}
}
