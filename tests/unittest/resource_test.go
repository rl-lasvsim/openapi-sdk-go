package simulation

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestGetHdMap(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	taskId, err := strconv.ParseUint(os.Getenv("QX_TASK_ID"), 10, 64)
	assert.NoError(t, err)
	recordId, err := strconv.ParseUint(os.Getenv("QX_RECORD_ID"), 10, 64)
	assert.NoError(t, err)
	fmt.Println(taskId, recordId)
	// Get scenario ID and version
	scenRes, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)
	assert.NoError(t, err)
	assert.NotEmpty(t, scenRes.ScenId, "scenario ID should not be empty")
	assert.NotEmpty(t, scenRes.ScenVer, "scenario version should not be empty")

	// Test getting HD map
	mapRes, err := cli.Resources.GetHdMap(scenRes.ScenId, scenRes.ScenVer)
	assert.NoError(t, err)
	assert.NotNil(t, mapRes, "HD map response should not be nil")
	assert.NotNil(t, mapRes.Data, "HD map data should not be nil")

	// Verify map components
	assert.NotNil(t, mapRes.Data.Junctions, "junctions should not be nil")
}

func TestGetHdMapWithInvalidInputs(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Token:    os.Getenv("QX_TOKEN"),
		Endpoint: os.Getenv("QX_ENDPOINT"),
	})

	// Test with empty scenario ID
	_, err := cli.Resources.GetHdMap("", "version")
	assert.Error(t, err, "should return error for empty scenario ID")

	// Test with empty version
	_, err = cli.Resources.GetHdMap("scenId", "")
	assert.Error(t, err, "should return error for empty version")

	// Test with invalid scenario ID
	_, err = cli.Resources.GetHdMap("invalid_scen_id", "version")
	assert.Error(t, err, "should return error for invalid scenario ID")
}

func TestGetHdMapDataValidation(t *testing.T) {
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

	// Get HD map
	mapRes, err := cli.Resources.GetHdMap(scenRes.ScenId, scenRes.ScenVer)
	assert.NoError(t, err)

	// 验证基本字段
	assert.NotNil(t, mapRes.Data.Header, "map header should not be nil")
	if mapRes.Data.Header != nil {
		assert.NotNil(t, mapRes.Data.Header.CenterPoint, "map center point should not be nil")
	}

	// 验证路口数据
	if len(mapRes.Data.Junctions) > 0 {
		junction := mapRes.Data.Junctions[0]
		assert.NotEmpty(t, junction.Id, "junction ID should not be empty")

		// 验证路口形状
		if junction.Shape != nil && len(junction.Shape.Points) > 0 {
			point := junction.Shape.Points[0]
			assert.NotNil(t, point.X, "junction shape point X should not be nil")
			assert.NotNil(t, point.Y, "junction shape point Y should not be nil")
		}

		// 验证路口内部道路
		if len(junction.Links) > 0 {
			link := junction.Links[0]
			assert.NotEmpty(t, link.Id, "junction link ID should not be empty")
		}
	}

	// 验证路段数据
	if len(mapRes.Data.Segments) > 0 {
		segment := mapRes.Data.Segments[0]
		assert.NotEmpty(t, segment.Id, "segment ID should not be empty")

		// 验证路段内的道路
		if len(segment.OrderedLinks) > 0 {
			link := segment.OrderedLinks[0]
			assert.NotEmpty(t, link.Id, "segment link ID should not be empty")

			// 验证车道
			if len(link.OrderedLanes) > 0 {
				lane := link.OrderedLanes[0]
				assert.NotEmpty(t, lane.Id, "lane ID should not be empty")
				assert.NotEmpty(t, lane.LinkId, "lane link ID should not be empty")
			}
		}
	}

	// 验证道路数据
	if len(mapRes.Data.Roads) > 0 {
		road := mapRes.Data.Roads[0]
		assert.NotEmpty(t, road.Id, "road ID should not be empty")
	}
}
