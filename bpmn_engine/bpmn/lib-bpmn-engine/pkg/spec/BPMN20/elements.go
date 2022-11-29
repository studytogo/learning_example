package BPMN20

type ElementType string

const (
	StartEvent             ElementType = "START_EVENT"
	EndEvent               ElementType = "END_EVENT"
	Task                   ElementType = "TASK"
	UserTask               ElementType = "USER_TASK"
	ManualTask             ElementType = "MANUAL_TASK"
	ServiceTask            ElementType = "SERVICE_TASK"
	ParallelGateway        ElementType = "PARALLEL_GATEWAY"
	ExclusiveGateway       ElementType = "EXCLUSIVE_GATEWAY"
	IntermediateCatchEvent ElementType = "INTERMEDIATE_CATCH_EVENT"
	EventBasedGateway      ElementType = "EVENT_BASED_GATEWAY"
	SubProcess             ElementType = "SUB_PROCESS"
)

type BaseElement interface {
	GetId() string
	GetName() string
	GetIncomingAssociation() []string
	GetOutgoingAssociation() []string
	GetType() ElementType
}

func (startEvent TStartEvent) GetId() string {
	return startEvent.Id
}

func (startEvent TStartEvent) GetName() string {
	return startEvent.Name
}

func (startEvent TStartEvent) GetIncomingAssociation() []string {
	return startEvent.IncomingAssociation
}

func (startEvent TStartEvent) GetOutgoingAssociation() []string {
	return startEvent.OutgoingAssociation
}

func (startEvent TStartEvent) GetType() ElementType {
	return StartEvent
}

func (endEvent TEndEvent) GetId() string {
	return endEvent.Id
}

func (endEvent TEndEvent) GetName() string {
	return endEvent.Name
}

func (endEvent TEndEvent) GetIncomingAssociation() []string {
	return endEvent.IncomingAssociation
}

func (endEvent TEndEvent) GetOutgoingAssociation() []string {
	return endEvent.OutgoingAssociation
}

func (endEvent TEndEvent) GetType() ElementType {
	return EndEvent
}

func (task TTask) GetId() string {
	return task.Id
}

func (task TTask) GetName() string {
	return task.Name
}

func (task TTask) GetIncomingAssociation() []string {
	return task.IncomingAssociation
}

func (task TTask) GetOutgoingAssociation() []string {
	return task.OutgoingAssociation
}

func (task TTask) GetType() ElementType {
	return Task
}

func (serviceTask TServiceTask) GetId() string {
	return serviceTask.Id
}

func (serviceTask TServiceTask) GetName() string {
	return serviceTask.Name
}

func (serviceTask TServiceTask) GetIncomingAssociation() []string {
	return serviceTask.IncomingAssociation
}

func (serviceTask TServiceTask) GetOutgoingAssociation() []string {
	return serviceTask.OutgoingAssociation
}

func (serviceTask TServiceTask) GetType() ElementType {
	return ServiceTask
}

func (userTask TUserTask) GetId() string {
	return userTask.Id
}

func (userTask TUserTask) GetName() string {
	return userTask.Name
}

func (userTask TUserTask) GetIncomingAssociation() []string {
	return userTask.IncomingAssociation
}

func (userTask TUserTask) GetOutgoingAssociation() []string {
	return userTask.OutgoingAssociation
}

func (userTask TUserTask) GetType() ElementType {
	return UserTask
}

func (manualTask TManualTask) GetId() string {
	return manualTask.Id
}

func (manualTask TManualTask) GetName() string {
	return manualTask.Name
}

func (manualTask TManualTask) GetIncomingAssociation() []string {
	return manualTask.IncomingAssociation
}

func (manualTask TManualTask) GetOutgoingAssociation() []string {
	return manualTask.OutgoingAssociation
}

func (manualTask TManualTask) GetType() ElementType {
	return ManualTask
}

func (parallelGateway TParallelGateway) GetId() string {
	return parallelGateway.Id
}

func (parallelGateway TParallelGateway) GetName() string {
	return parallelGateway.Name
}

func (parallelGateway TParallelGateway) GetIncomingAssociation() []string {
	return parallelGateway.IncomingAssociation
}

func (parallelGateway TParallelGateway) GetOutgoingAssociation() []string {
	return parallelGateway.OutgoingAssociation
}

func (parallelGateway TParallelGateway) GetType() ElementType {
	return ParallelGateway
}

func (exclusiveGateway TExclusiveGateway) GetId() string {
	return exclusiveGateway.Id
}

func (exclusiveGateway TExclusiveGateway) GetName() string {
	return exclusiveGateway.Name
}

func (exclusiveGateway TExclusiveGateway) GetIncomingAssociation() []string {
	return exclusiveGateway.IncomingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetOutgoingAssociation() []string {
	return exclusiveGateway.OutgoingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetType() ElementType {
	return ExclusiveGateway
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetId() string {
	return intermediateCatchEvent.Id
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetName() string {
	return intermediateCatchEvent.Name
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetIncomingAssociation() []string {
	return intermediateCatchEvent.IncomingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetOutgoingAssociation() []string {
	return intermediateCatchEvent.OutgoingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetType() ElementType {
	return IntermediateCatchEvent
}

func (eventBasedGateway TEventBasedGateway) GetId() string {
	return eventBasedGateway.Id
}

func (eventBasedGateway TEventBasedGateway) GetName() string {
	return eventBasedGateway.Name
}

func (eventBasedGateway TEventBasedGateway) GetIncomingAssociation() []string {
	return eventBasedGateway.IncomingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetOutgoingAssociation() []string {
	return eventBasedGateway.OutgoingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetType() ElementType {
	return EventBasedGateway
}

func (subProcess TSubProcess) GetId() string {
	return subProcess.Id
}

func (subProcess TSubProcess) GetName() string {
	return subProcess.Name
}

func (subProcess TSubProcess) GetIncomingAssociation() []string {
	return subProcess.IncomingAssociation
}

func (subProcess TSubProcess) GetOutgoingAssociation() []string {
	return subProcess.OutgoingAssociation
}

func (subProcess TSubProcess) GetType() ElementType {
	return SubProcess
}
