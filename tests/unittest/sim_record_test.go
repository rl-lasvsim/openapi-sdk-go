package simulation

import (
	"os"
	"strconv"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
	"github.com/stretchr/testify/assert"
)

func setupSimRecordTest(t *testing.T) (string, string, string, string) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	assert.NoError(t, err)
	recordId, err := strconv.ParseUint(os.Getenv("QX_RECORD_ID"), 10, 64)
	assert.NoError(t, err)

	// Get scenario ID and version
	scenRes, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)
	assert.NoError(t, err)

	// Get record IDs
	recordRes, err := cli.SimRecord.GetRecordIds(scenRes.ScenId, scenRes.ScenVer)
	assert.NoError(t, err)
	assert.Greater(t, len(recordRes.Ids), 0, "record list should not be empty")

	// Get vehicle ID
	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:  scenRes.ScenId,
		ScenVer: scenRes.ScenVer,
	})
	assert.NoError(t, err)
	defer simulator.Stop()

	vehRes, err := simulator.GetVehicleIdList()
	assert.NoError(t, err)
	assert.Greater(t, len(vehRes.List), 0, "vehicle list should not be empty")

	return recordRes.Ids[0], vehRes.List[0], scenRes.ScenId, scenRes.ScenVer
}

func TestGetRecordIds(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	assert.NoError(t, err)
	recordId, err := strconv.ParseUint(os.Getenv("QX_RECORD_ID"), 10, 64)
	assert.NoError(t, err)

	// Get scenario ID and version
	scenRes, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)
	assert.NoError(t, err)

	// Test getting record IDs
	res, err := cli.SimRecord.GetRecordIds(scenRes.ScenId, scenRes.ScenVer)
	assert.NoError(t, err)
	assert.NotNil(t, res.Ids, "record list should not be nil")
	assert.Greater(t, len(res.Ids), 0, "record list should not be empty")

	// Test with invalid inputs
	_, err = cli.SimRecord.GetRecordIds("", "")
	assert.Error(t, err, "should return error for empty inputs")
}

func TestGetTrackResults(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	recordId, objId, _, _ := setupSimRecordTest(t)

	// Test getting track results
	res, err := cli.SimRecord.GetTrackResults(recordId, objId)
	assert.NoError(t, err)
	assert.NotNil(t, res, "track results should not be nil")

	// Test with invalid inputs
	_, err = cli.SimRecord.GetTrackResults("", "")
	assert.Error(t, err, "should return error for empty inputs")
}

func TestGetSensorResults(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	recordId, objId, _, _ := setupSimRecordTest(t)

	// Test getting sensor results
	res, err := cli.SimRecord.GetSensorResults(recordId, objId)
	assert.NoError(t, err)
	assert.NotNil(t, res, "sensor results should not be nil")

	// Test with invalid inputs
	_, err = cli.SimRecord.GetSensorResults("", "")
	assert.Error(t, err, "should return error for empty inputs")
}

func TestGetStepResults(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	recordId, objId, _, _ := setupSimRecordTest(t)

	// Test getting step results
	res, err := cli.SimRecord.GetStepResults(recordId, objId)
	assert.NoError(t, err)
	assert.NotNil(t, res, "step results should not be nil")

	// Test with invalid inputs
	_, err = cli.SimRecord.GetStepResults("", "")
	assert.Error(t, err, "should return error for empty inputs")
}

func TestGetPathResults(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	recordId, objId, _, _ := setupSimRecordTest(t)

	// Test getting path results
	res, err := cli.SimRecord.GetPathResults(recordId, objId)
	assert.NoError(t, err)
	assert.NotNil(t, res, "path results should not be nil")

	// Test with invalid inputs
	_, err = cli.SimRecord.GetPathResults("", "")
	assert.Error(t, err, "should return error for empty inputs")
}

func TestGetReferenceLineResults(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	recordId, objId, _, _ := setupSimRecordTest(t)

	// Test getting reference line results
	res, err := cli.SimRecord.GetReferenceLineResults(recordId, objId)
	assert.NoError(t, err)
	assert.NotNil(t, res, "reference line results should not be nil")

	// Test with invalid inputs
	_, err = cli.SimRecord.GetReferenceLineResults("", "")
	assert.Error(t, err, "should return error for empty inputs")
}

func TestSimRecordResultsDataValidation(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	recordId, objId, _, _ := setupSimRecordTest(t)

	// Test track results data
	trackRes, err := cli.SimRecord.GetTrackResults(recordId, objId)
	if assert.NoError(t, err) && assert.NotNil(t, trackRes.Data) {
		if len(trackRes.Data) > 0 {
			track := trackRes.Data[0].Result
			assert.NotNil(t, track.X, "track X coordinate should not be nil")
			assert.NotNil(t, track.Y, "track Y coordinate should not be nil")
		}
	}

	// Test path results data
	pathRes, err := cli.SimRecord.GetPathResults(recordId, objId)
	if assert.NoError(t, err) && assert.NotNil(t, pathRes.Data) {
	}

	// Test reference line results data
	refLineRes, err := cli.SimRecord.GetReferenceLineResults(recordId, objId)
	if assert.NoError(t, err) && assert.NotNil(t, refLineRes.Data) {

	}
}
