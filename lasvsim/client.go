package lasvsim

// Client represents the main client for the API
type Client struct {
	config *HttpConfig

	httpClient *HttpClient

	// TrainTask   *TrainTask
	// Resources   *Resources
	ProcessTask *ProcessTask
	// Simulator   *Simulator
}

// NewClient creates a new API client
func NewClient(config *HttpConfig) *Client {
	client := &Client{
		config:     config,
		httpClient: NewHttpClient(config),
	}
	client.initCommonClient(client.httpClient)
	return client
}

// initCommonClient initializes the common client components
func (c *Client) initCommonClient(hCli *HttpClient) {
	// c.TrainTask = &TrainTask{client: NewHttpClient(config)}
	// c.Resources = &Resources{client: NewHttpClient(config)}
	c.ProcessTask = NewProcessTask(c.httpClient)
}

// InitSimulatorFromConfig initializes a simulator from the given configuration
func (c *Client) InitSimulatorFromConfig(simConfig SimulatorConfig) (*Simulator, error) {
	return NewSimulatorFromConfig(c.httpClient, simConfig)
}

// InitSimulatorFromSim initializes a simulator from existing simulation
func (c *Client) InitSimulatorFromSim(simId string, addr string) (*Simulator, error) {
	return NewSimulatorFromSim(c.httpClient, simId, addr)
}
