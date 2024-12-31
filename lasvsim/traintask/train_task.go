package traintask

import (
	"fmt"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
)

type TrainTask struct {
	httpClient *httpclient.HttpClient
}

func NewTrainTask(hCli *httpclient.HttpClient) *TrainTask {
	return &TrainTask{httpClient: hCli.Clone()}
}

func (p *TrainTask) GetSceneIdList(taskId uint64) (*GetSceneIdListRes, error) {
	var reply GetSceneIdListRes
	err := p.httpClient.Get(
		fmt.Sprintf("/openapi/train_task/%d/scene_id_list", taskId), map[string]string{},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
