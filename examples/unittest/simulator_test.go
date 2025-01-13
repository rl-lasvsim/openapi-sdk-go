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

func setupClient() *lasvsim.Client {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	return cli
}

func setupSimulator(t *testing.T) *simulation.Simulator {
	cli := setupClient()

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

func getSimTaskScen(t *testing.T, cli *lasvsim.Client) (string, string) {
	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	assert.NoError(t, err)
	recordId, err := strconv.ParseUint(os.Getenv("QX_RECORD_ID"), 10, 64)
	assert.NoError(t, err)
	res, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)
	assert.NoError(t, err)

	return res.ScenId, res.ScenVer
}

func TestSimulatorInitialization(t *testing.T) {
	// Test normal initialization
	simulator := setupSimulator(t)
	assert.NotNil(t, simulator, "simulator should be initialized")
	defer simulator.Stop()

	_, err := strconv.ParseUint(os.Getenv("QX_TASK_ID_NOT"), 10, 64)
	assert.Error(t, err, "should return error for invalid task ID")
}

func TestGetVehicleIdList(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test getting vehicle ID list
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")
}

func TestGetVehicleMovingInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test getting vehicle moving info
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	vehMovingInfos, err := simulator.GetVehicleMovingInfo([]string{res.List[0]})
	assert.NoError(t, err)
	vehMovingInfo := vehMovingInfos.MovingInfoDict[res.List[0]]
	assert.NotNil(t, vehMovingInfo, "not found vehicle moving info")
}

func TestSetVehicleMovingInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test setting vehicle moving info
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

func TestSimulatorStep(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test stepping the simulator
	stepRes, err := simulator.Step()
	assert.NoError(t, err)
	assert.NotNil(t, stepRes, "step result should not be nil")
}

func TestSimulatorStop(t *testing.T) {
	simulator := setupSimulator(t)

	// Test stopping the simulator
	err := simulator.Stop()
	assert.NoError(t, err)
}

func TestSimulatorReset(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test resetting the simulator
	resetRes, err := simulator.Reset(true)
	assert.NoError(t, err)
	assert.NotNil(t, resetRes, "reset result should not be nil")
}

func TestGetCurrentStage(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()
	cli := setupClient()
	scenId, scenVer := getSimTaskScen(t, cli)
	hdMap, err := cli.Resources.GetHdMap(scenId, scenVer)
	assert.NoError(t, err)
	assert.NotNil(t, hdMap, "hd map should not be nil")
	assert.Greater(t, len(hdMap.Data.Junctions), 0, "not found junction")

	// Test getting current stage
	stageRes, err := simulator.GetCurrentStage(hdMap.Data.Junctions[0].Id)
	if err != nil {
		return
	}
	assert.NotNil(t, stageRes, "current stage result should not be nil")
}

func TestGetMovementSignal(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()
	cli := setupClient()
	scenId, scenVer := getSimTaskScen(t, cli)
	hdMap, err := cli.Resources.GetHdMap(scenId, scenVer)
	assert.NoError(t, err)
	assert.NotNil(t, hdMap, "hd map should not be nil")
	assert.Greater(t, len(hdMap.Data.Junctions), 0, "not found junction")
	if len(hdMap.Data.Junctions[0].Movements) == 0 {
		return
	}

	// Test getting movement signal
	signalRes, err := simulator.GetMovementSignal(hdMap.Data.Junctions[0].Movements[0].Id)
	assert.NoError(t, err)
	assert.NotNil(t, signalRes, "movement signal result should not be nil")
}

func TestGetSignalPlan(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()
	cli := setupClient()

	scenId, scenVer := getSimTaskScen(t, cli)
	hdMap, err := cli.Resources.GetHdMap(scenId, scenVer)
	assert.NoError(t, err)
	assert.NotNil(t, hdMap, "hd map should not be nil")
	assert.Greater(t, len(hdMap.Data.Junctions), 0, "not found junction")

	// Test getting signal plan
	signalPlanRes, err := simulator.GetSignalPlan(hdMap.Data.Junctions[0].Id)
	if err != nil {
		return
	}
	assert.NotNil(t, signalPlanRes, "signal plan result should not be nil")
}

func TestGetMovementList(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()
	cli := setupClient()
	scenId, scenVer := getSimTaskScen(t, cli)
	hdMap, err := cli.Resources.GetHdMap(scenId, scenVer)
	assert.NoError(t, err)
	assert.NotNil(t, hdMap, "hd map should not be nil")
	assert.Greater(t, len(hdMap.Data.Junctions), 0, "not found junction")

	// Test getting movement list
	movementListRes, err := simulator.GetMovementList(hdMap.Data.Junctions[0].Id)
	assert.NoError(t, err)
	assert.NotNil(t, movementListRes, "movement list result should not be nil")
}

func TestGetVehicleBaseInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle base info
	baseInfoRes, err := simulator.GetVehicleBaseInfo([]string{res.List[0]})
	assert.NoError(t, err)
	assert.NotNil(t, baseInfoRes, "vehicle base info result should not be nil")
}

func TestGetVehiclePosition(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle position
	positionRes, err := simulator.GetVehiclePosition([]string{res.List[0]})
	assert.NoError(t, err)
	assert.NotNil(t, positionRes, "vehicle position result should not be nil")
}

func TestGetVehicleControlInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle control info
	controlInfoRes, err := simulator.GetVehicleControlInfo([]string{res.List[0]})
	assert.NoError(t, err)
	assert.NotNil(t, controlInfoRes, "vehicle control info result should not be nil")
}

func TestGetVehiclePerceptionInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle perception info
	perceptionInfoRes, err := simulator.GetVehiclePerceptionInfo(res.List[0])
	assert.NoError(t, err)
	assert.NotNil(t, perceptionInfoRes, "vehicle perception info result should not be nil")
}

func TestGetVehicleReferenceLines(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle reference lines
	referenceLinesRes, err := simulator.GetVehicleReferenceLines(res.List[0])
	assert.NoError(t, err)
	assert.NotNil(t, referenceLinesRes, "vehicle reference lines result should not be nil")
}

func TestGetVehiclePlanningInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle planning info
	planningInfoRes, err := simulator.GetVehiclePlanningInfo(res.List[0])
	assert.NoError(t, err)
	assert.NotNil(t, planningInfoRes, "vehicle planning info result should not be nil")
}

func TestGetVehicleNavigationInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle navigation info
	navigationInfoRes, err := simulator.GetVehicleNavigationInfo(res.List[0])
	assert.NoError(t, err)
	assert.NotNil(t, navigationInfoRes, "vehicle navigation info result should not be nil")
}

func TestGetVehicleCollisionStatus(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle collision status
	collisionStatusRes, err := simulator.GetVehicleCollisionStatus(res.List[0])
	assert.NoError(t, err)
	assert.NotNil(t, collisionStatusRes, "vehicle collision status result should not be nil")
}

func TestGetVehicleTargetSpeed(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test getting vehicle target speed
	targetSpeedRes, err := simulator.GetVehicleTargetSpeed(res.List[0])
	assert.NoError(t, err)
	assert.NotNil(t, targetSpeedRes, "vehicle target speed result should not be nil")
}

func TestSetVehiclePlanningInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test setting vehicle planning info
	planningInfoRes, err := simulator.SetVehiclePlanningInfo(res.List[0], []*simulation.Point{
		{X: 0.0, Y: 0.0},
	})
	assert.NoError(t, err)
	assert.NotNil(t, planningInfoRes, "set vehicle planning info result should not be nil")
}

func TestSetVehicleControlInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test setting vehicle control info
	controlInfoRes, err := simulator.SetVehicleControlInfo(res.List[0], utils.Ptr(1.0), utils.Ptr(1.0))
	assert.NoError(t, err)
	assert.NotNil(t, controlInfoRes, "set vehicle control info result should not be nil")
}

func TestSetVehiclePosition(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test setting vehicle position
	positionRes, err := simulator.SetVehiclePosition(res.List[0], &simulation.Point{}, utils.Ptr(1.0))
	assert.NoError(t, err)
	assert.NotNil(t, positionRes, "set vehicle position result should not be nil")
}

func TestSetVehicleBaseInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	// Test setting vehicle base info
	baseInfoRes, err := simulator.SetVehicleBaseInfo(res.List[0], &simulation.ObjBaseInfo{
		Width:  1.0,
		Length: 1.0,
		Height: 1.0,
	}, &simulation.DynamicInfo{
		FrontWheelStiffness: 1.0,
		FrontAxleToCenter:   1.0,
	})
	assert.NoError(t, err)
	assert.NotNil(t, baseInfoRes, "set vehicle base info result should not be nil")
}

func TestSetVehicleDestination(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch vehicle ID
	res, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(res.List), 0, "not found vehicle id list")

	positions, err := simulator.GetVehiclePosition([]string{res.List[0]})
	assert.NoError(t, err)
	positon := positions.PositionDict[res.List[0]]
	// Test setting vehicle destination
	destinationRes, err := simulator.SetVehicleDestination(res.List[0], &simulation.Point{
		X: positon.Point.X + 1,
		Y: positon.Point.Y,
	})
	assert.NoError(t, err)
	if destinationRes == nil {
		t.Skip("destination is nil, skipping test")
	}
}

func TestGetPedIdList(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test getting pedestrian ID list
	pedIdListRes, err := simulator.GetPedIdList()
	assert.NoError(t, err)
	assert.NotNil(t, pedIdListRes, "pedestrian ID list result should not be nil")
}

func TestGetPedBaseInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test getting pedestrian base info
	pedBaseInfoRes, err := simulator.GetPedBaseInfo([]string{"pedId"})
	assert.NoError(t, err)
	assert.NotNil(t, pedBaseInfoRes, "pedestrian base info result should not be nil")
}

func TestSetPedPosition(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	pedListRes, err := simulator.GetPedIdList()
	assert.NoError(t, err)
	if len(pedListRes.List) == 0 {
		t.Skip("no pedestrian found, skipping test")
	}

	pedPositionRes, err := simulator.SetPedPosition(pedListRes.List[0], &simulation.Point{}, utils.Ptr(1.0))
	assert.NoError(t, err)
	assert.NotNil(t, pedPositionRes, "set pedestrian position result should not be nil")
}

func TestGetNMVIdList(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test getting non-motorized vehicle ID list
	nmvIdListRes, err := simulator.GetNMVIdList()
	assert.NoError(t, err)
	assert.NotNil(t, nmvIdListRes, "non-motorized vehicle ID list result should not be nil")
}

func TestGetNMVBaseInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test getting non-motorized vehicle base info
	nmvBaseInfoRes, err := simulator.GetNMVBaseInfo([]string{"nmvId"})
	assert.NoError(t, err)
	assert.NotNil(t, nmvBaseInfoRes, "non-motorized vehicle base info result should not be nil")
}

func TestSetNMVPosition(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test setting non-motorized vehicle position
	nmvIdListRes, err := simulator.GetNMVIdList()
	assert.NoError(t, err)
	if nmvIdListRes == nil || len(nmvIdListRes.List) == 0 {
		return
	}
	nmvPositionRes, err := simulator.SetNMVPosition(nmvIdListRes.List[0], &simulation.Point{}, utils.Ptr(1.0))
	assert.NoError(t, err)
	assert.NotNil(t, nmvPositionRes, "set non-motorized vehicle position result should not be nil")
}

func TestGetStepSpawnIdList(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Test getting step spawn ID list
	stepSpawnIdListRes, err := simulator.GetStepSpawnIdList()
	assert.NoError(t, err)
	assert.NotNil(t, stepSpawnIdListRes, "step spawn ID list result should not be nil")
}

func TestGetParticipantBaseInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch participant ID
	res, err := simulator.GetStepSpawnIdList()
	assert.NoError(t, err)
	if res == nil || len(res.IdList) == 0 {
		return
	}

	assert.Greater(t, len(res.IdList), 0, "not found participant id list")

	// Test getting participant base info
	participantBaseInfoRes, err := simulator.GetParticipantBaseInfo([]string{res.IdList[0]})
	assert.NoError(t, err)
	assert.NotNil(t, participantBaseInfoRes, "participant base info result should not be nil")
}

func TestGetParticipantMovingInfo(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch participant ID
	res, err := simulator.GetStepSpawnIdList()
	assert.NoError(t, err)
	if res == nil || len(res.IdList) == 0 {
		return
	}

	// Test getting participant moving info
	participantMovingInfoRes, err := simulator.GetParticipantMovingInfo([]string{res.IdList[0]})
	assert.NoError(t, err)
	assert.NotNil(t, participantMovingInfoRes, "participant moving info result should not be nil")
}

func TestGetParticipantPosition(t *testing.T) {
	simulator := setupSimulator(t)
	defer simulator.Stop()

	// Fetch participant ID
	res, err := simulator.GetStepSpawnIdList()
	assert.NoError(t, err)
	if res == nil || len(res.IdList) == 0 {
		return
	}

	// Test getting participant position
	participantPositionRes, err := simulator.GetParticipantPosition([]string{res.IdList[0]})
	assert.NoError(t, err)
	assert.NotNil(t, participantPositionRes, "participant position result should not be nil")
}
