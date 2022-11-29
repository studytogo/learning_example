package goque_study

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"github.com/beeker1121/goque"
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
	//_, err = s.Push([]byte("item value"))
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

	item, err := s.Pop()
	fmt.Println(item.ID)    // 1
	fmt.Println(item.Key)   // [0 0 0 0 0 0 0 1]
	fmt.Println(item.Value) // [105 116 101 109 32 118 97 108 117 101]
	fmt.Println(item.ToString())
}
