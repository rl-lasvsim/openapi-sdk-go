package qianxing

type SimulatorConfig struct {
	ScenID      string `json:"scen_id,omitempty"`
	ScenVer     string `json:"scen_ver,omitempty"`
	SimRecordID int    `json:"sim_record_id,omitempty"`
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

type Simulator struct {
	httpClient *HttpClient

	simulationId string
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

func (s *Simulator) Stop() error {
	var reply StopReq

	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/stop",
		&StopReq{
			SimulationId: s.simulationId,
		},
		&reply,
	)
	if err != nil {
		return err
	}

	return nil
}
