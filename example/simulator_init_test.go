package example

import (
	"os"
	"testing"

	qianxing "github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
)

func TestSimulatorInit(t *testing.T) {
	cli := qianxing.NewClient(&qianxing.HttpConfig{
		Endpoint: os.Getenv("QX_ENDPOINT"),
		Token:    os.Getenv("QX_TOKEN"),
	})

	simulator, err := cli.InitSimulatorFromConfig(qianxing.SimulatorConfig{
		ScenID:  "78677711905027",
		ScenVer: "0",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = simulator.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
