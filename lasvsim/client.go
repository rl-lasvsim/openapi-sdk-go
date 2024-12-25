package lasvsim

import (
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/processtask"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/resource"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simrecord"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/traintask"
)

// Client represents the main client for the API
type Client struct {
	config *httpclient.HttpConfig

	httpClient *httpclient.HttpClient

	TrainTask   *traintask.TrainTask
	Resources   *resource.Resource
	ProcessTask *processtask.ProcessTask
	Simulator   *simulation.Simulator
	SimRecord   *simrecord.SimRecord
}

// NewClient creates a new API client
func NewClient(config *httpclient.HttpConfig) *Client {
	client := &Client{
		config:     config,
		httpClient: httpclient.NewHttpClient(config),
	}
	client.initCommonClient()
	return client
}

// initCommonClient initializes the common client components
func (c *Client) initCommonClient() {
	c.TrainTask = traintask.NewTrainTask(c.httpClient)
	c.Resources = resource.NewResource(c.httpClient)
	c.ProcessTask = processtask.NewProcessTask(c.httpClient)
	c.SimRecord = simrecord.NewSimRecord(c.httpClient)
}

// InitSimulatorFromConfig initializes a simulator from the given configuration
func (c *Client) InitSimulatorFromConfig(simConfig simulation.SimulatorConfig) (*simulation.Simulator, error) {
	return simulation.NewSimulatorFromConfig(c.httpClient, simConfig)
}

// InitSimulatorFromSim initializes a simulator from existing simulation
func (c *Client) InitSimulatorFromSim(simId string, addr string) (*simulation.Simulator, error) {
	return simulation.NewSimulatorFromSim(c.httpClient, simId, addr)
}
