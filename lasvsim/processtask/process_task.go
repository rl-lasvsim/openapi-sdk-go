package processtask

import (
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
)

type ProcessTask struct {
	httpClient *httpclient.HttpClient
}

func NewProcessTask(hCli *httpclient.HttpClient) *ProcessTask {
	return &ProcessTask{httpClient: hCli.Clone()}
}

func (p *ProcessTask) CopyRecord(taskId uint64, recordId uint64) (*CopyRecordRes, error) {
	var reply CopyRecordRes
	err := p.httpClient.Post(
		"/openapi/process_task/v2/record/copy",
		&CopyRecordReq{TaskId: taskId, RecordId: recordId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (p *ProcessTask) GetRecordScenario(taskId uint64, recordId uint64) (*GetRecordScenarioRes, error) {
	var reply GetRecordScenarioRes
	err := p.httpClient.Post(
		"/openapi/process_task/v2/record/scenario/get",
		&GetRecordScenarioReq{TaskId: taskId, RecordId: recordId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (p *ProcessTask) GetTaskRecordIds(taskId uint64) (*GetTaskRecordIdsRes, error) {
	var reply GetTaskRecordIdsRes
	err := p.httpClient.Post(
		"/openapi/process_task/v2/record/id_list",
		&GetTaskRecordIdsReq{TaskId: taskId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
