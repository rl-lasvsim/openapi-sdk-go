package simulation

import (
	"os"
	"strconv"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

func TestNewSimulatorFromConfig(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	recordId, err := strconv.ParseUint(os.Getenv("QX_RECORD_ID"), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	res, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)

	if err != nil {
		t.Fatal(err)
	}

	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:  res.ScenId,
		ScenVer: res.ScenVer,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer simulator.Stop()

	stepRes, err := simulator.Step()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(stepRes)

}
