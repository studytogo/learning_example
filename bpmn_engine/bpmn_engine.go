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
	process, err := bpmnEngine.LoadFromFile(`E:/bpmns/diagram9.bpmn`)
	if err != nil {
		fmt.Println("================", err)
		panic("file \"simple_task.bpmn\" can't be read.")
	}
	//fmt.Println("++++++++++++++++", fmt.Sprintf("%+v", process))
	// register a handler for a service task by defined task type
	bpmnEngine.AddTaskHandler("Activity_0rtzuz3", printContextHandler1)
	//bpmnEngine.AddTaskHandler("Activity_01whl0g", printContextHandler2)
	bpmnEngine.AddTaskHandler("Activity_1s0fkkr", printContextHandler2)
	// setup some variables
	variables := map[string]string{}
	variables["foo"] = "bar"
	// and execute the process
	_, err = bpmnEngine.CreateAndRunInstance(process.ProcessKey, variables)
	fmt.Println("-------", err)
}

func printContextHandler1(job bpmn_engine.ActivatedJob) {
	println("< Hello World >11111111111")
	//println(fmt.Sprintf("ElementId                = %s", job.ElementId))
	//println(fmt.Sprintf("BpmnProcessId            = %s", job.BpmnProcessId))
	//println(fmt.Sprintf("ProcessDefinitionKey     = %d", job.ProcessDefinitionKey))
	//println(fmt.Sprintf("ProcessDefinitionVersion = %d", job.ProcessDefinitionVersion))
	//println(fmt.Sprintf("CreatedAt                = %s", job.CreatedAt))
	//println(fmt.Sprintf("Variable 'foo'           = %s", job.GetVariable("foo")))
	//job.Complete() // don't forget this one, or job.Fail("foobar")
	job.Fail("yyyyyy")
}

func printContextHandler2(job bpmn_engine.ActivatedJob) {
	println("< Hello World >222222222222")
	//println(fmt.Sprintf("ElementId                = %s", job.ElementId))
	//println(fmt.Sprintf("BpmnProcessId            = %s", job.BpmnProcessId))
	//println(fmt.Sprintf("ProcessDefinitionKey     = %d", job.ProcessDefinitionKey))
	//println(fmt.Sprintf("ProcessDefinitionVersion = %d", job.ProcessDefinitionVersion))
	//println(fmt.Sprintf("CreatedAt                = %s", job.CreatedAt))
	//println(fmt.Sprintf("Variable 'foo'           = %s", job.GetVariable("foo")))
	job.Complete() // don't forget this one, or job.Fail("foobar")
}
