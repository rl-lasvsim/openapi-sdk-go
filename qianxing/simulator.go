package qianxing

type SimulatorConfig struct {
	ScenID      string `json:"scen_id,omitempty"`
	ScenVer     string `json:"scen_ver,omitempty"`
	SimRecordID int    `json:"sim_record_id,omitempty"`
}

type Simulator struct {
	httpClient   *HttpClient
	simulationId string
}

func NewSimulatorFromConfig(hCli *HttpClient, cfg SimulatorConfig) (*Simulator, error) {
	cloneCli := hCli.Clone()

	simtor := &Simulator{
		httpClient: cloneCli,
	}

	err := simtor.initFromConfig(cfg)
	if err != nil {
		return nil, err
	}

	return simtor, nil
}

func NewSimulatorFromSim(hCli *HttpClient, simId, simAddr string) (*Simulator, error) {
	cloneCli := hCli.Clone()
	simtor := &Simulator{
		httpClient: cloneCli,
	}

	err := simtor.initFromSim(simId, simAddr)
	if err != nil {
		return nil, err
	}

	return simtor, nil
}

func (s *Simulator) initFromConfig(simConfig SimulatorConfig) error {
	var reply struct {
		SimulationId   string `json:"simulation_id"`
		SimulationAddr string `json:"simulation_addr"`
	}

	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/init",
		&InitReq{
			ScenID:      simConfig.ScenID,
			ScenVer:     simConfig.ScenVer,
			SimRecordID: simConfig.SimRecordID,
		},
		&reply,
	)
	if err != nil {
		return err
	}

	return s.initFromSim(reply.SimulationId, reply.SimulationAddr)
}

func (s *Simulator) initFromSim(simId, simAddr string) error {
	s.httpClient.headers["simulation_id"] = simId
	s.httpClient.headers["x-md-rl-direct-addr"] = simAddr
	s.simulationId = simId
	return nil
}

func (s *Simulator) Step() (*StepRes, error) {
	var reply StepRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/step",
		&StepReq{SimulationID: s.simulationId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) Stop() error {
	var reply StopRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/stop",
		&StopReq{SimulationId: s.simulationId},
		&reply,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Simulator) Reset(resetTrafficFlow bool) (*ResetRes, error) {
	var reply ResetRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/reset",
		&ResetReq{SimulationID: s.simulationId, ResetTrafficFlow: resetTrafficFlow},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
