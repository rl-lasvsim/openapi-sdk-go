package beginner

import (
	"fmt"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
	"github.com/stretchr/testify/assert"
)

func TestCreateCosimTask1(t *testing.T) {
	var (
		endpoint        = "http://8.146.201.197:30080/dev"
		token           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjQ5LCJvaWQiOjI2LCJuYW1lIjoi5ZGo5YWr5Y-q6IO955yL5Zy65pmv5ZWK5ZWKIiwiaWRlbnRpdHkiOiJub3JtYWwiLCJwZXJtaXNzaW9ucyI6W10sImlzcyI6InVzZXIiLCJzdWIiOiJMYXNWU2ltIiwiZXhwIjoxNzQ1MzkwODUyLCJuYmYiOjE3NDQ3ODYwNTIsImlhdCI6MTc0NDc4NjA1MiwianRpIjoiNDkifQ.w9--Gsq2tLRwVv4xnD7H17uMr8e-GbOXnYw7ZUQt7ZY"
		taskId   uint64 = 15038
		recordId uint64 = 19513
	)

	// 1. 初始化客户端
	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: endpoint, // 接口地址
		Token:    token,    // 授权token
	})

	// 2. 拷贝剧本, 返回的结构中NewRecordId字段就是新创建的剧本ID, 仿真结束后可到该剧本下查看结果详情
	newRecord, err := cli.ProcessTask.CopyRecord(taskId, recordId)
	assert.NoError(t, err)

	// 3. 通过拷贝的场景Id、Version和SimRecordId初始化仿真器
	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:      newRecord.ScenId,
		ScenVer:     newRecord.ScenVer,
		SimRecordID: newRecord.SimRecordId,
	})

	// res, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)
	// assert.NoError(t, err)

	// simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
	// 	ScenID:  res.ScenId,
	// 	ScenVer: res.ScenVer,
	// })
	// assert.NoError(t, err)

	// 关闭仿真器, 释放服务器资源
	defer simulator.Stop()

	// 获取测试车辆列表
	testVehicleList, err := simulator.GetTestVehicleIdList()
	assert.NoError(t, err)

	stepRes, err := simulator.Reset(true,
		[]*simulation.ResetVehicleConfig{
			{
				VehicleId: testVehicleList.List[0],
				LinkPath:  []string{"sg4_lk0", "sg1_lk0"},
			},
		},
		&simulation.ResetEnvPtcs{
			VehicleConf: &simulation.VehicleDistribution{
				Density:   0.2,
				SizeRatio: []float64{0.5, 0.5, 0.0},
			},
			NmvConf: &simulation.NMVDistribution{
				Density:      0.2,
				SubtypeRatio: []float64{0.1, 0.9},
			},
		})
	assert.NoError(t, err)
	fmt.Println(stepRes)
	// 记录仿真器运行状态(true: 运行中; false: 运行结束), 任务运行过程中持续更新该状态
	var isRunning = true
	// 使测试车辆环形行驶
	for isRunning {
		// 执行仿真器步骤, 返回的结果中记录了当前任务的运行状态
		stepRes, err := simulator.Step()
		assert.NoError(t, err)
		fmt.Println(stepRes)

		isRunning = stepRes.Code.IsRuning()
	}
	fmt.Println("success")
}

// func TestCreateCosimTask1(t *testing.T) {
// 	var (
// 		endpoint        = "https://qianxing-api.risenlighten.com"
// 		token           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIwLCJvaWQiOjEwMiwibmFtZSI6IuWkj-aigeWHryIsImlkZW50aXR5Ijoibm9ybWFsIiwicGVybWlzc2lvbnMiOltdLCJpc3MiOiJ1c2VyIiwic3ViIjoiTGFzVlNpbSIsImV4cCI6MTc0NDc3MDE0OSwibmJmIjoxNzQ0MTY1MzQ5LCJpYXQiOjE3NDQxNjUzNDksImp0aSI6IjIwIn0.wCDRWxukQSyKNiA1nLaMFsUV9QXiu9OTCY1foYmMQlI"
// 		taskId   uint64 = 11477
// 		recordId uint64 = 29373
// 	)

// 	// 1. 初始化客户端
// 	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
// 		Endpoint: endpoint, // 接口地址
// 		Token:    token,    // 授权token
// 	})

// 	// 2. 拷贝剧本, 返回的结构中NewRecordId字段就是新创建的剧本ID, 仿真结束后可到该剧本下查看结果详情
// 	newRecord, err := cli.ProcessTask.CopyRecord(taskId, recordId)
// 	assert.NoError(t, err)

// 	// 3. 通过拷贝的场景Id、Version和SimRecordId初始化仿真器
// 	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
// 		ScenID:      newRecord.ScenId,
// 		ScenVer:     newRecord.ScenVer,
// 		SimRecordID: newRecord.SimRecordId,
// 	})

// 	// res, err := cli.ProcessTask.GetRecordScenario(taskId, recordId)
// 	// assert.NoError(t, err)

// 	// simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
// 	// 	ScenID:  res.ScenId,
// 	// 	ScenVer: res.ScenVer,
// 	// })
// 	// assert.NoError(t, err)

// 	// 关闭仿真器, 释放服务器资源
// 	defer simulator.Stop()

// 	// 获取测试车辆列表
// 	testVehicleList, err := simulator.GetTestVehicleIdList()
// 	assert.NoError(t, err)

// 	stepRes, err := simulator.Reset(true, []*simulation.ResetVehicleConfig{
// 		{
// 			VehicleId: testVehicleList.List[0],
// 			LinkPath:  []string{"sg1_lk0", "sg1_lk0"},
// 		},
// 	})
// 	assert.NoError(t, err)
// 	fmt.Println(stepRes)
// 	// 记录仿真器运行状态(true: 运行中; false: 运行结束), 任务运行过程中持续更新该状态
// 	var isRunning = true
// 	// 使测试车辆环形行驶
// 	for isRunning {
// 		// 执行仿真器步骤, 返回的结果中记录了当前任务的运行状态
// 		stepRes, err := simulator.Step()
// 		assert.NoError(t, err)
// 		fmt.Println(stepRes)

// 		isRunning = stepRes.Code.IsRuning()
// 	}
// 	fmt.Println("success")
// }
