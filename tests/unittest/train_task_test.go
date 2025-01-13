package simulation

import (
	"os"
	"strconv"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestGetSceneIdList(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TRAIN_TASK_ID"), 10, 64)
	assert.NoError(t, err)

	// Test getting scene ID list
	res, err := cli.TrainTask.GetSceneIdList(taskId)
	assert.NoError(t, err)
	assert.NotNil(t, res.SceneIdList, "scene ID list should not be nil")

	// Test with invalid task ID
	_, err = cli.TrainTask.GetSceneIdList(0)
	assert.Error(t, err, "should return error for invalid task ID")
}

func TestGetSceneIdListDataValidation(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TRAIN_TASK_ID"), 10, 64)
	assert.NoError(t, err)

	// Get scene ID list
	res, err := cli.TrainTask.GetSceneIdList(taskId)
	assert.NoError(t, err)

	// Validate scene IDs
	if len(res.SceneIdList) > 0 {
		for _, sceneId := range res.SceneIdList {
			assert.NotEmpty(t, sceneId, "scene ID should not be empty")
		}
	}
}

func TestGetSceneIdListWithInvalidConfig(t *testing.T) {
	// Test with invalid token
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    "invalid_token",
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TRAIN_TASK_ID"), 10, 64)
	assert.NoError(t, err)

	_, err = cli.TrainTask.GetSceneIdList(taskId)
	assert.Error(t, err, "should return error for invalid token")

	// Test with invalid endpoint
	cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: "invalid_endpoint",
	})

	_, err = cli.TrainTask.GetSceneIdList(taskId)
	assert.Error(t, err, "should return error for invalid endpoint")
}
