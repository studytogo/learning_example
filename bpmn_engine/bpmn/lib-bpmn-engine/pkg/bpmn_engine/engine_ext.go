package bpmn_engine

import (
	"fmt"

	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20"
)

type FilterContext struct {
	processInstanceInfo *ProcessInstanceInfo
	Element             BPMN20.BaseElement
}

// GetVariable from the process instance's variable context
func (fctx FilterContext) GetVariable(key string) interface{} {
	return fctx.processInstanceInfo.GetVariable(key)
}

// SetVariable to the process instance's variable context
func (fctx FilterContext) SetVariable(key string, value interface{}) {
	fctx.processInstanceInfo.SetVariable(key, value)
}

func (state *BpmnEngineState) exclusivelyFilterByName(instance *ProcessInstanceInfo, element BPMN20.BaseElement, flows []BPMN20.TSequenceFlow) ([]BPMN20.TSequenceFlow, error) {
	var ret []BPMN20.TSequenceFlow

	filter, ok := state.filters[element.GetId()]
	if ok {
		for _, flow := range flows {
			if flow.Name != "" {
				branch, err := filter(FilterContext{processInstanceInfo: instance, Element: element})
				if err != nil {
					return nil, err
				}

				if flow.Name == branch {
					ret = append(ret, flow)
					break
				}
			}
		}
	}

	if len(ret) == 0 {
		for _, flow := range flows {
			if flow.Name == "" {
				ret = append(ret, flow)
				break
			}
		}
	}

	if len(ret) == 0 {
		return nil, fmt.Errorf("No flow found for gateway %s", element.GetId())
	}

	return ret, nil
}
