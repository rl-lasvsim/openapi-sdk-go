package resource

import (
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
)

type Resource struct {
	httpClient *httpclient.HttpClient
}

func NewResource(hCli *httpclient.HttpClient) *Resource {
	return &Resource{httpClient: hCli.Clone()}
}

func (s *Resource) GetHdMap(scenId, scenVer string) (*GetHdMapRes, error) {
	var reply GetHdMapRes
	err := s.httpClient.Post(
		"/openapi/resource/v2/scenario/map/get",
		&GetHdMapReq{ScenId: scenId, ScenVer: scenVer},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
