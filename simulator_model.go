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

type StepReq struct {
	SimulationID string `json:"simulation_id"`
}

type StepRes struct {
	// Define fields according to the expected response structure
}

type ResetReq struct {
	SimulationID     string `json:"simulation_id"`
	ResetTrafficFlow bool   `json:"reset_traffic_flow"`
}

type ResetRes struct {
	// Define fields according to the expected response structure
}
