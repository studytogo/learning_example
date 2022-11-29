package bpmn_engine

import (
	"fmt"
	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine"
)

func StudyBpmnEngine() {
	//fmt.Println("----------------")
	// create a new named engine
	bpmnEngine := bpmn_engine.New("a name")
	// basic example loading a BPMN from file,\
	process, err := bpmnEngine.LoadFromFile(`E:\gowork\src\learning_example\bpmn_engine\bpmns\test_sunflow.bpmn`)
	if err != nil {
		fmt.Println("================", err)
		panic("file \"simple_task.bpmn\" can't be read.")
	}
	//fmt.Println("++++++++++++++++", fmt.Sprintf("%+v", process))
	// register a handler for a service task by defined task type
	bpmnEngine.AddTaskHandler("sub", printContextHandler1)
	bpmnEngine.AddTaskHandler("subTask", printContextHandler1)
	// setup some variables
	n := new(Node)
	n.NodeId = "task3"
	n.IsRun = false
	variables := map[string]interface{}{}
	variables["mxNodeIdSelectorRun"] = n
	// and execute the process
	instance, err := bpmnEngine.CreateInstance(process.ProcessKey, variables)
	if err != nil {
		fmt.Println("111111", err)
	}
	_, err = bpmnEngine.RunOrContinueInstance(instance.GetInstanceKey())
	fmt.Println("-------", err)
}

type Node struct {
	NodeId string
	IsRun  bool
}

func printContextHandler1(job bpmn_engine.ActivatedJob) {
	//fmt.Println("0000000000000000000000000000")
	println(job.ElementId)
	//println(fmt.Sprintf("ElementId                = %s", job.ElementId))
	//println(fmt.Sprintf("BpmnProcessId            = %s", job.BpmnProcessId))
	//println(fmt.Sprintf("ProcessDefinitionKey     = %d", job.ProcessDefinitionKey))
	//println(fmt.Sprintf("ProcessDefinitionVersion = %d", job.ProcessDefinitionVersion))
	//println(fmt.Sprintf("CreatedAt                = %s", job.CreatedAt))
	//println(fmt.Sprintf("Variable 'mxNodeIdSelectorRun' = %s", job.GetVariable("mxNodeIdSelectorRun")))
	//if job.GetVariable("mxNodeIdSelectorRun") != nil {
	//	nodeInfo := job.GetVariable("mxNodeIdSelectorRun").(*Node)
	//	fmt.Println("-----------------", fmt.Sprintf("%+v", nodeInfo))
	//	nodeInfo.NodeId = job.ElementId
	//	job.SetVariable("mxNodeIdSelectorRun", nodeInfo)
	//}
	job.Complete() // don't forget this one, or job.Fail("foobar")
	//job.Fail("yyyyyy")
}
