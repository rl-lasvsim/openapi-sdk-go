package traintask

import "github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"

type TrainTask struct {
	httpClient *httpclient.HttpClient
}

func NewTrainTask(hCli *httpclient.HttpClient) *TrainTask {
	return &TrainTask{httpClient: hCli.Clone()}
}

func (p *TrainTask) CopyRecord(taskId uint64) (*GetSceneIdListRes, error) {
	var reply GetSceneIdListRes
	err := p.httpClient.Post(
		"/openapi/train_task/{task_id}/scene_id_list",
		&GetSceneIdListReq{TaskId: taskId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
