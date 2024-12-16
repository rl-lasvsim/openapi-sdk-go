package qianxing

type Resource struct {
	httpClient *HttpClient
}

func NewResource(hCli *HttpClient) *Resource {
	return &Resource{httpClient: hCli.Clone()}
}

func (p *Resource) CopyRecord(taskId uint64, recordId uint64) CopyRecordRes {

}

func (p *Resource) GetRecordScenario(taskId uint64, recordId uint64) GetRecordScenarioRes {

}

func (p *Resource) GetTaskRecordIds(taskId uint64, recordId uint64) GetTaskRecordIdsRes {

}

func (p *Resource) GetHdMap() (*GetHdMapRes, error)       {}
func (p *Resource) GetScenario() (*GetScenarioRes, error) {}
