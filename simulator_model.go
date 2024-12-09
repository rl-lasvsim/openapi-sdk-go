package qianxing

type InitReq struct {
	ScenID      string `json:"scen_id,omitempty"`
	ScenVer     string `json:"scen_ver,omitempty"`
	SimRecordID int    `json:"sim_record_id,omitempty"`
}

type InitRes struct {
	SimulationId   string `json:"simulation_id"`
	SimulationAddr string `json:"simulation_addr"`
}

type StopReq struct {
	SimulationId string `json:"simulation_id"`
}

type StopRes struct {
}
