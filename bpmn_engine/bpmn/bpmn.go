package bpmn

import (
	"sync"

	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine"
)

const engineVarKey = "__engine_vars_key__"

type Engine struct {
	id        string
	variables map[string]interface{}
	bpmn_engine.BpmnEngineState
}

func NewEngine(id string) *Engine {
	return &Engine{
		id:              id,
		BpmnEngineState: bpmn_engine.New(id),
		variables: map[string]interface{}{
			engineVarKey: new(sync.Map),
		},
	}
}
