package postgres_db

import (
	"fmt"
	"github.com/pkg/errors"
	udb "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"k8s.io/apimachinery/pkg/util/json"
	"strings"
)

type TaskRecord struct {
	Id   uint64           `db:"id,omitempty"`
	Task postgresql.JSONB `db:"task,omitempty"`
}

const (
	Pending   = 1
	Running   = 2
	Successed = 3
	Faile     = 4
)

var enumMap = map[int]string{1: "Pending", 2: "Running", 3: "Successed", 4: "Faile"}
var enumStatusMap = map[string]int{"Failed": 1, "Terminated": 2, "Suspended": 3, "Unknown": 4, "Pending": 5, "Running": 6, "Successed": 7}

type TaskJson struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	State  int    `json:"state"`
	NodeId string `json:"nodeId"`
}

func PostgresDBStudy() {
	udb.LC().SetLevel(udb.LogLevelDebug)
	accessor, err := NewPgAccessor()
	if err != nil {
		fmt.Println("11111", err)
	}
	//collections, err := accessor.Collections()
	//if err != nil {
	//	fmt.Println("22222", err)
	//}
	//
	//for i := range collections {
	//	fmt.Printf("-> %q\n", collections[i].Name())
	//}
	m := make(map[string]interface{})
	m["workflowId"] = "3dd72571-ccf4-45e6-8951-dd015dec482f"
	results := []*TaskRecord{}
	err = accessor.Collection("tasks").
		Find(`(task->>'workflowId')::text = '3dd72571-ccf4-45e6-8951-dd015dec482f'`).
		All(&results)
	if err != nil {
		fmt.Println("333333", err)
	}

	viewResult := make(map[string]map[string]int)
	for _, v := range results {
		r := TaskJson{}
		j, err := v.Task.MarshalJSON()
		if err != nil {
			fmt.Println(">>>>>", err)
		}
		//fmt.Println("???????", string(j))
		err = json.Unmarshal(j, &r)
		if err != nil {
			fmt.Println("++++++", err)
		}

		if _, ok := viewResult[r.NodeId]; ok {
			if value, exist := viewResult[r.NodeId][enumMap[r.State]]; exist {
				viewResult[r.NodeId][enumMap[r.State]] = value + 1
			} else {
				viewResult[r.NodeId][enumMap[r.State]] = 1
			}

			if value, exist := viewResult[r.NodeId]["total"]; exist {
				viewResult[r.NodeId]["total"] = value + 1
			} else {
				viewResult[r.NodeId]["total"] = 1
			}
		} else {
			initTaskInfo(viewResult, r.NodeId)
			viewResult[r.NodeId][enumMap[r.State]] = 1
			viewResult[r.NodeId]["total"] = 1
			fmt.Println("_-----------", fmt.Sprintf("%+v", viewResult))
		}
	}

	//查询节点状态
	//for k, v := range viewResult {
	//	for k1, v1 := range v {
	//		if k1 !=
	//	}
	//}
	fmt.Println("PPPPPPPPP", fmt.Sprintf("%+v", viewResult))
}

func NewPgAccessor() (udb.Session, error) {
	settings := postgresql.ConnectionURL{
		Database: "mxsche",
		Host:     "192.168.70.202:30433",
		User:     "postgres",
		Password: "KtlV7PpWPc",
	}

	sess, err := postgresql.Open(settings)
	if err != nil {
		return nil, errors.Wrap(err, "postgresql.Open error")
	}

	return sess, nil
}

func makeQueryFormKV(target string, k string, v interface{}) string {
	return fmt.Sprintf("%s ->> '%s' = '%s'", target, k, v)
}

func makeQueryFormCond(target string, cond map[string]interface{}) string {
	list := []string{}
	for k, v := range cond {
		list = append(list, makeQueryFormKV(target, k, v))
	}
	return strings.Join(list, " and ")
}

func initTaskInfo(data map[string]map[string]int, nodeId string) {
	data[nodeId] = make(map[string]int)
	data[nodeId]["Successed"] = 0
	data[nodeId]["Running"] = 0
	data[nodeId]["Failed"] = 0
	data[nodeId]["Terminated"] = 0
	data[nodeId]["Suspended"] = 0
	data[nodeId]["Unknown"] = 0
	data[nodeId]["Pending"] = 0
	data[nodeId]["status"] = 0
}
