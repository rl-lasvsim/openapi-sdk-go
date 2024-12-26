package simrecord

import (
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
)

type SimRecord struct {
	httpClient *httpclient.HttpClient
}

func NewSimRecord(hCli *httpclient.HttpClient) *SimRecord {
	return &SimRecord{httpClient: hCli.Clone()}
}

func (p *SimRecord) GetRecordIds(scenId string, scenVer string) (*GetRecordIdsRes, error) {
	var reply GetRecordIdsRes
	err := p.httpClient.Post(
		"/openapi/sim_record/v1/ids/get",
		&GetRecordIdsReq{ScenId: scenId, ScenVer: scenVer},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (p *SimRecord) GetTrackResults(id string, objId string) (*GetTrackResultsRes, error) {
	var reply GetTrackResultsRes
	err := p.httpClient.Post(
		"/openapi/sim_record/v1/track_result/get",
		&GetTrackResultsReq{Id: id, ObjId: objId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (p *SimRecord) GetSensorResults(id string, objId string) (*GetSensorResultsRes, error) {
	var reply GetSensorResultsRes
	err := p.httpClient.Post(
		"/openapi/sim_record/v1/sensor_result/get",
		&GetSensorResultsReq{Id: id, ObjId: objId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (p *SimRecord) GetStepResults(id string, objId string) (*GetStepResultsRes, error) {
	var reply GetStepResultsRes
	err := p.httpClient.Post(
		"/openapi/sim_record/v1/step_result/get",
		&GetStepResultsReq{Id: id, ObjId: objId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (p *SimRecord) GetPathResults(id string, objId string) (*GetPathResultsRes, error) {
	var reply GetPathResultsRes
	err := p.httpClient.Post(
		"/openapi/sim_record/v1/path_result/get",
		&GetPathResultsReq{Id: id, ObjId: objId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (p *SimRecord) GetReferenceLineResults(id string, objId string) (*GetReferenceLineResultsRes, error) {
	var reply GetReferenceLineResultsRes
	err := p.httpClient.Post(
		"/openapi/sim_record/v1/reference_line_result/get",
		&GetReferenceLineResultsReq{Id: id, ObjId: objId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
