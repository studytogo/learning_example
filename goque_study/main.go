package goque_study

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"github.com/beeker1121/goque"
	"time"
)

type Object struct {
	X int `json:"x"`
}

func GoqueNormal() {
	q := queue.New(128)
	q.Put("1")
	q.Put("2")
	i := q.Len()
	fmt.Println("-----len", i)
}

func GoqueStudy() {
	s, err := goque.OpenStack("./data_dir")
	if err != nil {
		fmt.Println("1111", err)
	}
	defer s.Close()
	//item, err := s.Push([]byte("item value"))
	//if err != nil {
	//	fmt.Println("2222", err)
	//}
	//// or
	//_, err = s.PushString("item value")
	//if err != nil {
	//	fmt.Println("3333", err)
	//}
	//// or
	//_, err = s.PushObject(Object{X: 1})
	//if err != nil {
	//	fmt.Println("4444", err)
	//}
	//// or
	//_, err = s.PushObjectAsJSON(Object{X: 1})
	//if err != nil {
	//	fmt.Println("4444", err)
	//}
	item, err := s.PeekByOffset(6)
	if err != nil {
		fmt.Println("-------------", err)
	}

	//item, err := s.Pop()
	fmt.Println(s.Length())
	fmt.Println(item.ID)    // 1
	fmt.Println(item.Key)   // [0 0 0 0 0 0 0 1]
	fmt.Println(item.Value) // [105 116 101 109 32 118 97 108 117 101]
	fmt.Println(item.ToString())
}

func StudyPriorityQueue() {
	pq, err := goque.OpenPriorityQueue("data_dir", goque.DESC)
	if err != nil {
		fmt.Println("11111", err)
		return
	}
	defer pq.Close()
	pq.Enqueue(1, []byte("1"))
	pq.Enqueue(3, []byte("3"))
	pq.Enqueue(2, []byte("2"))
	pq.Enqueue(4, []byte("4"))
	pq.Enqueue(3, []byte("3-1"))

	for {
		item, err := pq.Dequeue()
		if err != nil {
			fmt.Println("pppppp", err)
		}
		if item == nil {
			fmt.Println("------没有数据跳过")
			return
		}
		fmt.Println("item.ID", item.ID)               // 1
		fmt.Println("item.Priority", item.Priority)   // 0
		fmt.Println("item.Key", string(item.Key))     // [0 58 0 0 0 0 0 0 0 1]
		fmt.Println("item.Value", string(item.Value)) // [105 116 101 109 32 118 97 108 117 101]
		//fmt.Println(item.ToString()) // item value
		time.Sleep(1 * time.Second)
		fmt.Println("---------------------------------------")
	}
}

type MemoryQueue struct {
	Name string
	Sort int
}

func (mq MemoryQueue) Compare(item queue.Item) int {
	// DESC
	other, ok := item.(MemoryQueue)
	if !ok {
		return -1
	}

	p1, p2 := mq.Sort, other.Sort

	if p1 > p2 {
		return -1
	}

	if p1 < p2 {
		return 1
	}

	return 0
}

func StudyMemoryPriorityQueue() {
	q := queue.NewPriorityQueue(512, false)
	mq1 := MemoryQueue{Sort: 5, Name: "5a"}
	q.Put(mq1)
	mq2 := MemoryQueue{Sort: 3, Name: "3a"}
	q.Put(mq2)
	mq3 := MemoryQueue{Sort: 4, Name: "4a"}
	q.Put(mq3)
	mq4 := MemoryQueue{Sort: 2, Name: "2a"}
	q.Put(mq4)
	mq5 := MemoryQueue{Sort: 10, Name: "10a"}
	q.Put(mq5)
	mq6 := MemoryQueue{Sort: 10, Name: "10b"}
	q.Put(mq6)

	for {
		get, err := q.Get(1)
		if err != nil {
			fmt.Println("------------", err)
		}
		fmt.Println(fmt.Sprintf("%+v", get))
	}
}
