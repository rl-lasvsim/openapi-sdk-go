package qianxing

type CopyRecordRes struct {
	SimRecordId string `json:"sim_record_id"`
	ScenId      string `json:"scen_id"`
	ScenVer     string `json:"scen_ver"`
	NewRecordId uint64 `json:"new_record_id"`
}

type GetRecordScenarioRes struct {
	ScenId  string `json:"scen_id"`
	ScenVer string `json:"scen_ver"`
}

type GetTaskRecordIdsRes struct {
	RecordIds []uint64 `json:"record_ids"`
}
