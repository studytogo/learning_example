package nats_study

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"reflect"
	"strconv"
	"sync"
	"time"
)

func NatsStudy() {
	//pubMessage()
	//subMessage()
	//NatsQueueMessageReceive()
	select {}
}
func pubMessage() {
	testAddress := "nats://192.168.70.31:30744"
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

func DeleteJetStream() {
	testAddress := "nats://192.168.70.202:30744"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		fmt.Println("111", err)
	}
	js, err := nc.JetStream()
	if err != nil {
		fmt.Println("222", err)
	}
	name := "shenyang-prod_v0_1_4_mxsche_tasks"

	info, _ := js.StreamInfo(name)
	subjects := []string{"shenyang-prod_v0_1_4_mxsche_workflows.*"}

	fmt.Println("===================", info != nil)

	if info != nil {
		if reflect.DeepEqual(subjects, info.Config.Subjects) {
			fmt.Println("lllllllllllllllllll")
		} else {
			info, err = js.UpdateStream(&nats.StreamConfig{
				Name: name,
			})
		}
	} else {
		info, err = js.AddStream(&nats.StreamConfig{
			Name:     name,
			Subjects: subjects,
		})
	}

	//err = js.DeleteStream("shenyang-prod_v0_1_4_mxsche_tasks")

	if err != nil {
		fmt.Println("111", err)
	}

}

func SubMessage() {
	//js.Subscribe("ORDERS.scratch", func(msg *nats.Msg) {
	//	fmt.Printf("Received a JetStream message: %s\n", string(msg.Data))
	//	msg.Ack()
	//})
	testAddress := "nats://192.168.70.202:30744"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		fmt.Println("111", err)
	}
	js, err := nc.JetStream()
	if err != nil {
		fmt.Println("222", err)
	}

	sub, err := js.PullSubscribe("kkkktest", "MONITOR", nats.PullMaxWaiting(128))
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

func NatsPutKV() {
	testAddress := "nats://192.168.70.202:30744"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		fmt.Println("111", err)
	}
	js, err := nc.JetStream()
	if err != nil {
		fmt.Println("222", err)
	}
	kv, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: "zzg",
	})
	if err != nil {
		fmt.Println("333", err)
	}
	time.Sleep(20 * time.Second)
	kv.PutString("555", "abc")

}

func NatsAcceptKV() {
	testAddress := "nats://192.168.70.31:30744"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		fmt.Println("111", err)
	}
	//go func() {
	//	for {
	//		fmt.Println("nc.IsClosed()", nc.IsClosed())
	//		time.Sleep(10 * time.Minute)
	//	}
	//}()
	js, err := nc.JetStream()
	if err != nil {
		fmt.Println("222", err)
	}
	kv, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: "ggg",
	})
	if err != nil {
		fmt.Println("333", err)
	}
	c, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctxNats := nats.Context(c)
	var w = sync.Mutex{}
	for i := 0; i < 10000; i++ {
		go func() {
			w.Lock()
			watch, err := kv.Watch("aaa", ctxNats)
			if err != nil {
				fmt.Println("444", err)
			}
			w.Unlock()
			fmt.Println(i, "--------监听key")
			for {
				select {
				case e := <-watch.Updates():
					if e != nil {
						fmt.Println("ttttttt", e.Operation() == nats.KeyValuePut)
					}
					fmt.Println("hhhhhhhhhh", e)
				}
			}
		}()
	}
	select {}

}

func Request() {
	testAddress := "nats://192.168.70.202:30744"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		fmt.Println("---111", err)
	}
	defer nc.Drain()

	_, err2 := nc.Subscribe("greet", func(msg *nats.Msg) {
		fmt.Println("+//////", string(msg.Data))
		//name := msg.Subject[6:]
		msg.Respond([]byte("kkkkk, " + string(msg.Data)))
	})
	fmt.Println("---222", err2)
	select {}
}

func TestRequest_Reply() {
	testAddress := "nats://192.168.70.202:30744"
	nc, _ := nats.Connect(testAddress)
	defer nc.Drain()
	for {
		//设置随机数种子，由于种子数值，每次启动都不一样
		//所以每次随机数都是随机的
		//rand.Seed(time.Now().UnixNano())
		//随机生成100以内的正整数
		//intn := rand.Intn(10)
		topic := "greet"
		Reply(nc, topic)
		time.Sleep(1 * time.Second)
	}
}

func Reply(nc *nats.Conn, topic string) {
	rep, err := nc.Request("greet", []byte("999"), 5*time.Second)
	fmt.Println("+++1111", err)
	fmt.Println(string(rep.Data))
}

func NatsQueueMessageReceive() {
	testAddress := "nats://192.168.70.202:30744"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	subj := "mxsche-dev_v0_1_4_normal_mxsche_tasks.submitted"

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     subj,
		Subjects: []string{subj},
	})
	if err != nil {
		fmt.Println("aaaaa", err)
	}

	// Create the queue group
	//if _, err := js.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
	//	fmt.Printf("Received a message1: %s\n", string(msg.Data))
	//}); err != nil {
	//	fmt.Println("11")
	//	log.Fatal(err)
	//}
	sub, err := js.PullSubscribe(subj, "order-review", nats.MaxAckPending(-1))
	if err != nil {
		fmt.Println("2222", err)
		return
	}
	fmt.Println("等待接收")
	for {
		msgs, err := sub.Fetch(1)
		if err != nil {
			if err == nats.ErrTimeout {
				continue
			}
			fmt.Println("---ttt", err)
			return
		}

		for _, msg := range msgs {
			msg.Ack()

			fmt.Println("message....", string(msg.Data))
		}

		time.Sleep(1 * time.Second)
	}

	//fmt.Printf("Listening for messages on subject [%s], queue [%s]\n", subj, queue)
	//
	//fmt.Println("等待接收")
	//select {}
}

func NatsQueueMessagePublish() {
	testAddress := "nats://192.168.70.202:30744"
	nc, err := nats.Connect(testAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	subj := "mxsche-dev_v0_1_4_normal_mxsche_tasks.submitted"

	// Publish a message
	i := 0
	for {
		time.Sleep(3 * time.Second)
		msg := []byte("Hello World" + strconv.Itoa(i))
		if _, err := js.Publish(subj, msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Published message [%s] to subject [%s]\n", string(msg), subj)
		i++
	}

}
