package bpmn

import (
	"fmt"
	"log"
	"testing"

	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine"
)

var globalProcessInfo *bpmn_engine.ProcessInfo

func TestSub(t *testing.T) {
	// create a new named engine
	bpmnEngine := bpmn_engine.New("a name")
	// basic example loading a BPMN from file,
	process, err := bpmnEngine.LoadFromFile("sub.bpmn")
	if err != nil {
		panic("file \"simple_task.bpmn\" can't be read.")
	}
	globalProcessInfo = process
	// register a handler for a service task by defined task type
	bpmnEngine.AddTaskHandler("task_a", printContextHandler)
	bpmnEngine.AddTaskHandler("task_call_sub_1", subprocessCall)
	bpmnEngine.AddTaskHandler("task_b", printContextHandler)

	bpmnEngine.AddTaskHandler("task_sub_a", printContextHandler)

	// setup some variables
	variables := map[string]interface{}{}
	variables["foo"] = "bar"
	// and execute the process
	_, err = bpmnEngine.CreateAndRunInstance(process.ProcessKey, variables)
	if err != nil {
		log.Println("err", err)
	}
}

func printContextHandler(job bpmn_engine.ActivatedJob) {
	println("< Hello World >")
	println(fmt.Sprintf("ElementId                = %s", job.ElementId))
	println(fmt.Sprintf("BpmnProcessId            = %s", job.BpmnProcessId))
	println(fmt.Sprintf("ProcessDefinitionKey     = %d", job.ProcessDefinitionKey))
	println(fmt.Sprintf("ProcessDefinitionVersion = %d", job.ProcessDefinitionVersion))
	println(fmt.Sprintf("CreatedAt                = %s", job.CreatedAt))
	println(fmt.Sprintf("Variable 'foo'           = %s", job.GetVariable("foo")))
	job.Complete() // don't forget this one, or job.Fail("foobar")
}

func subprocessCall(job bpmn_engine.ActivatedJob) {
	println("< subprocess >")
	println(fmt.Sprintf("ElementId                = %s", job.ElementId))
	println(fmt.Sprintf("BpmnProcessId            = %s", job.BpmnProcessId))
	println(fmt.Sprintf("ProcessDefinitionKey     = %d", job.ProcessDefinitionKey))
	println(fmt.Sprintf("ProcessDefinitionVersion = %d", job.ProcessDefinitionVersion))
	println(fmt.Sprintf("CreatedAt                = %s", job.CreatedAt))
	println(fmt.Sprintf("Variable 'foo'           = %s", job.GetVariable("foo")))

	//job.Complete() // don't forget this one, or job.Fail("foobar")
	sub := globalProcessInfo.GetTSubProcess(job.ElementId)
	if sub != nil {
		//		dump.P(sub)
	}
	println("type", job.ElementType)

	def := createSubProcessBPMNDefinition(sub, globalProcessInfo.GetTDefinitions())

	// create a new named engine
	bpmnEngine := bpmn_engine.New("sub name")
	pi, err := bpmnEngine.LoadFromDefinitions(def, [16]byte{})
	if err != nil {
		println("1", err)
		job.Fail(err.Error())
		return
	}

	bpmnEngine.AddTaskHandler("task_sub_a", printContextHandler)
	bpmnEngine.AddTaskHandler("task_sub_b", printContextHandler)
	_, err = bpmnEngine.CreateAndRunInstance(pi.ProcessKey, map[string]interface{}{})
	if err != nil {
		println("2", err)
		job.Fail(err.Error())
		return
	}

	job.Complete()
}
