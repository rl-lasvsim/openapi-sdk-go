package simulation

import (
	"os"
	"strconv"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/utils"
	"github.com/stretchr/testify/assert"
)

func setupSimulator(t *testing.T) *simulation.Simulator {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	assert.NoError(t, err)
	recordId, err := strconv.ParseUint(os.Getenv("QX_RECORD_ID"), 10, 64)
	assert.NoError(t, err)

	res, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)
	assert.NoError(t, err)

	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:  res.ScenId,
		ScenVer: res.ScenVer,
	})
	assert.NoError(t, err)
	return simulator
}

func TestMovingInfoSet(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	vehMovingInfos, err := simulator.GetVehicleMovingInfo([]string{res.List[0]})
	assert.NoError(t, err)
	vehMovingInfo := vehMovingInfos.MovingInfoDict[res.List[0]]
	assert.NotNil(t, vehMovingInfo, "not found vehicle moving info")

	_, err = simulator.SetVehicleMovingInfo(res.List[0], utils.Ptr(vehMovingInfo.U+1), nil, nil, nil, nil, nil)
	assert.NoError(t, err)
	modifedVehMovingInfos, err := simulator.GetVehicleMovingInfo([]string{res.List[0]})
	assert.NoError(t, err)
	modifedVehMovingInfo := modifedVehMovingInfos.MovingInfoDict[res.List[0]]
	assert.NotNil(t, modifedVehMovingInfo, "not found vehicle moving info")
	assert.Equal(t, vehMovingInfo.U+1, modifedVehMovingInfo.U, "not modified vehicle moving info")
	assert.Equal(t, vehMovingInfo.V, modifedVehMovingInfo.V, "not modified vehicle moving info")
}
