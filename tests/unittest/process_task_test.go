package simulation

import (
	"os"
	"strconv"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestProcessTaskGetRecordScenario(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	assert.NoError(t, err)
	recordId, err := strconv.ParseUint(os.Getenv("QX_RECORD_ID"), 10, 64)
	assert.NoError(t, err)

	// Test getting record scenario
	res, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ScenId, "scenario ID should not be empty")
	assert.NotEmpty(t, res.ScenVer, "scenario version should not be empty")
}

func TestProcessTaskGetTaskRecordIds(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	assert.NoError(t, err)

	// Test getting task record IDs
	res, err := cli.ProcessTask.GetTaskRecordIds(taskId)
	assert.NoError(t, err)
	assert.NotNil(t, res.RecordIds, "record ID list should not be nil")
}

func TestProcessTaskCopyRecord(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	assert.NoError(t, err)
	recordId, err := strconv.ParseUint(os.Getenv("QX_RECORD_ID"), 10, 64)
	assert.NoError(t, err)

	// Test copying record
	res, err := cli.ProcessTask.CopyRecord(taskId, recordId)
	assert.NoError(t, err)
	assert.NotZero(t, res.NewRecordId, "copied record ID should not be zero")

	// Verify the copied record exists
	records, err := cli.ProcessTask.GetTaskRecordIds(taskId)
	assert.NoError(t, err)
	found := false
	for _, id := range records.RecordIds {
		if id == res.NewRecordId {
			found = true
			break
		}
	}
	assert.True(t, found, "copied record should exist in task records")

	// Test copying with invalid task ID
	_, err = cli.ProcessTask.CopyRecord(0, recordId)
	assert.Error(t, err, "should return error for invalid task ID")
}

func TestProcessTaskInvalidInputs(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	// Test invalid task ID
	_, err := cli.ProcessTask.GetTaskRecordIds(0)
	assert.Error(t, err, "should return error for invalid task ID")

	// Test invalid record ID
	_, err = cli.ProcessTask.GetRecordScenario(0, 0)
	assert.Error(t, err, "should return error for invalid record ID")
}
