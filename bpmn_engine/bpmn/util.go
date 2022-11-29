package bpmn

import (
	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20"
)

func TSubProcessToTProcess(sub *BPMN20.TSubProcess) *BPMN20.TProcess {
	process := &BPMN20.TProcess{
		XMLName:                sub.XMLName,
		StartEvents:            sub.StartEvents,
		EndEvents:              sub.EndEvents,
		SequenceFlows:          sub.SequenceFlows,
		ServiceTasks:           sub.ServiceTasks,
		Tasks:                  sub.Tasks,
		UserTasks:              sub.UserTasks,
		ManualTasks:            sub.ManualTasks,
		ParallelGateway:        sub.ParallelGateway,
		ExclusiveGateway:       sub.ExclusiveGateway,
		IntermediateCatchEvent: sub.IntermediateCatchEvent,
		EventBasedGateway:      sub.EventBasedGateway,
	}
	return process
}

func createSubProcessBPMNDefinition(sub *BPMN20.TSubProcess, origin *BPMN20.TDefinitions) *BPMN20.TDefinitions {
	process := TSubProcessToTProcess(sub)
	return &BPMN20.TDefinitions{
		XMLName:            origin.XMLName,
		Id:                 origin.Id + "_" + sub.Id,
		Name:               origin.Name + "_" + sub.Name,
		TargetNamespace:    origin.TargetNamespace,
		ExpressionLanguage: origin.ExpressionLanguage,
		TypeLanguage:       origin.TypeLanguage,
		Exporter:           origin.Exporter,
		ExporterVersion:    origin.ExporterVersion,
		Process:            *process,
		Messages:           origin.Messages,
	}
}
