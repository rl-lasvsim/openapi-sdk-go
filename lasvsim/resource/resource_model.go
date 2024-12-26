package resource

type GetHdMapReq struct {
	ScenId  string `json:"scen_id"`
	ScenVer string `json:"scen_ver"`
}
type GetHdMapRes struct {
	Data *Qxmap `json:"data"`
}
