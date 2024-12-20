package processtask

type CopyRecordReq struct {
	// 任务ID
	TaskId uint64 `json:"task_id"`
	// 剧本ID
	RecordId uint64 `json:"record_id"`
}
type CopyRecordRes struct {
	SimRecordId string `json:"sim_record_id"`
	ScenId      string `json:"scen_id"`
	ScenVer     string `json:"scen_ver"`
	NewRecordId uint64 `json:"new_record_id"`
}

type GetRecordScenarioReq struct {
	// 任务ID
	TaskId uint64 `json:"task_id"`
	// 剧本ID
	RecordId uint64 `json:"record_id"`
}
type GetRecordScenarioRes struct {
	ScenId  string `json:"scen_id"`
	ScenVer string `json:"scen_ver"`
}

type GetTaskRecordIdsReq struct {
	TaskId uint64 `json:"task_id"`
}
type GetTaskRecordIdsRes struct {
	RecordIds []uint64 `json:"record_ids"`
}
