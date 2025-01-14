# Lasvsim OpenAPI SDK for Go

千行仿真平台（Lasvsim）的 Go SDK。提供了一种简单直观的方式来控制和获取自动驾驶场景的仿真。

## 安装

您可以直接使用 go get 安装该软件包：

```bash
go get github.com/rl-lasvsim/openapi-sdk-go/lasvsim
```

## 快速开始

以下是 SDK 使用的简单示例：

```go
package main

import (
	"fmt"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

func main() {
	var (
		// 接口地址
		endpoint string = os.Getenv("QX_ENDPOINT") // 线上环境地址: https://qianxing-api.risenlighten.com
		// 授权token
		token string = os.Getenv("QX_TOKEN") // 登录仿真平台后访问https://qianxing.risenlighten.com/#/usecenter/personalCenter, 点击最下面按钮复制token

		// 登录仿真平台, 选择想要进行联合仿真的任务及剧本，赋值给下面的taskId和recordId变量
		// 仿真任务ID
		taskId uint64 = 0
		// 剧本ID
		recordId uint64 = 0
	)

	// 1. 初始化客户端
	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: endpoint, // 接口地址
		Token:    token,    // 授权token
	})

	// 2. 拷贝剧本, 返回的结构中NewRecordId字段就是新创建的剧本ID, 仿真结束后可到该剧本下查看结果详情
	newRecord, err := cli.ProcessTask.CopyRecord(taskId, recordId)
	if err!=nil{
    panic(err)
  }
  fmt.Println("拷贝剧本成功")

	// 3. 通过拷贝的场景Id、Version和SimRecordId初始化仿真器
	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:      newRecord.ScenId,
		ScenVer:     newRecord.ScenVer,
		SimRecordID: newRecord.SimRecordId,
	})
	if err!=nil{
    panic(err)
  }
  fmt.Println("初始化仿真器成功")

	// 关闭仿真器, 释放服务器资源
	defer simulator.Stop()

	// 获取测试车辆列表
	testVehicleList, err := simulator.GetTestVehicleIdList()
	if err!=nil{
    panic(err)
  }
  fmt.Println("测试车辆ID列表:", testVehicleList)

	// 使测试车辆环形行驶
	for i := 0; i < 50; i++ {
		// 设置方向盘转角30度, 纵向加速度5
		var steWheel float64 = 10
		var lonAcc float64 = 0.05
		// 设置车辆的控制信息
		_, err := simulator.SetVehicleControlInfo(testVehicleList.List[0], &steWheel, &lonAcc)
		if err!=nil{
      panic(err)
    }

		// 执行仿真器步骤
		stepRes, err := simulator.Step()
		if err!=nil{
      panic(err)
    }
    fmt.Println("第 %d 步结果: %v\n", i, stepRes)
	}

	// 可在此处继续调用其他接口, 查看联合仿真文档: https://www.risenlighten.com/#/union

	// 仿真结束后, 到千行仿真平台对应的taskId/recordId下查看联合仿真结果详情
	fmt.Printf("https://qianxing.risenlighten.com/#/configuration/circleTask?id=%d\n", taskId)

	// 如想直接查看本次联合仿真的回放视频, 可访问下面网址：
	fmt.Printf("https://qianxing.risenlighten.com/#/sampleRoad/cartest/?id=%d&record_id=%d&sim_record_id=%s\n", taskId, newRecord.NewRecordId, newRecord.SimRecordId)
}
```

## API 参考

### Client

```go
type Client struct {
    config *httpclient.HttpConfig
    httpClient *httpclient.HttpClient
    TrainTask   *traintask.TrainTask
    Resources   *resource.Resource
    ProcessTask *processtask.ProcessTask
    Simulator   *simulation.Simulator
    SimRecord   *simrecord.SimRecord
}
```

#### 方法

- `NewClient(config *httpclient.HttpConfig) *Client`
  使用给定配置创建新的 API 客户端。

- `InitSimulatorFromConfig(simConfig simulation.SimulatorConfig) (*simulation.Simulator, error)`
  从给定配置初始化仿真器。

- `InitSimulatorFromSim(simId string, addr string) (*simulation.Simulator, error)`
  从现有仿真初始化仿真器。

### 子模块

#### TrainTask

管理训练任务。

- `CopyRecord(taskId uint64) (*GetSceneIdListRes, error)`
  复制训练任务记录

#### Resources

处理资源管理。

- `GetHdMap(scenId, scenVer string) (*GetHdMapRes, error)`
  获取高清地图

#### ProcessTask

管理处理任务。

- `CopyRecord(taskId uint64, recordId uint64) (*CopyRecordRes, error)`
  复制处理任务记录

- `GetRecordScenario(taskId uint64, recordId uint64) (*GetRecordScenarioRes, error)`
  获取记录场景

- `GetTaskRecordIds(taskId uint64) (*GetTaskRecordIdsRes, error)`
  获取任务记录 ID 列表

#### Simulator

提供仿真功能。

- `Step() (*StepRes, error)`
  执行仿真步进

- `Stop() error`
  停止仿真

- `Reset(resetTrafficFlow bool) (*ResetRes, error)`
  重置仿真

- `GetCurrentStage(junctionId string) (*GetCurrentStageRes, error)`
  获取当前交通灯阶段

- `GetMovementSignal(movementId string) (*GetMovementSignalRes, error)`
  获取移动信号

- `GetSignalPlan(junctionId string) (*GetSignalPlanRes, error)`
  获取信号计划

- `GetMovementList(junctionId string) (*GetMovementListRes, error)`
  获取移动列表

- `GetVehicleIdList() (*GetVehicleIdListRes, error)`
  获取车辆 ID 列表

- `GetTestVehicleIdList() (*GetTestVehicleIdListRes, error)`
  获取测试车辆 ID 列表

- `GetVehicleBaseInfo(vehicleIdList []string) (*GetVehicleBaseInfoRes, error)`
  获取车辆基本信息

- `GetVehiclePosition(vehicleIdList []string) (*GetVehiclePositionRes, error)`
  获取车辆位置

- `GetVehicleMovingInfo(vehicleIdList []string) (*GetVehicleMovingInfoRes, error)`
  获取车辆移动信息

- `GetVehicleControlInfo(vehicleIdList []string) (*GetVehicleControlInfoRes, error)`
  获取车辆控制信息

- `GetVehiclePerceptionInfo(vehicleId string) (*GetVehiclePerceptionInfoRes, error)`
  获取车辆感知信息

- `GetVehicleReferenceLines(vehicleId string) (*GetVehicleReferenceLinesRes, error)`
  获取车辆参考线

- `GetVehiclePlanningInfo(vehicleId string) (*GetVehiclePlanningInfoRes, error)`
  获取车辆规划信息

- `GetVehicleNavigationInfo(vehicleId string) (*GetVehicleNavigationInfoRes, error)`
  获取车辆导航信息

- `GetVehicleCollisionStatus(vehicleId string) (*GetVehicleCollisionStatusRes, error)`
  获取车辆碰撞状态

- `GetVehicleTargetSpeed(vehicleId string) (*GetVehicleTargetSpeedRes, error)`
  获取车辆目标速度

- `SetVehiclePlanningInfo(vehicleId string, planningPath []*Point) (*SetVehiclePlanningInfoRes, error)`
  设置车辆规划信息

- `SetVehicleControlInfo(vehicleId string, steWheel *float64, lonAcc *float64) (*SetVehicleControlInfoRes, error)`
  设置车辆控制信息

- `SetVehiclePosition(vehicleId string, point *Point, phi *float64) (*SetVehiclePositionRes, error)`
  设置车辆位置

- `SetVehicleMovingInfo(vehicleId string, u, v, w, uAcc, vAcc, wAcc *float64) (*SetVehicleMovingInfoRes, error)`
  设置车辆移动信息

- `SetVehicleBaseInfo(vehicleId string, baseInfo *ObjBaseInfo, dynamicInfo *DynamicInfo) (*SetVehicleBaseInfoRes, error)`
  设置车辆基本信息

- `SetVehicleLinkNav(vehicleId string, linkNav []string) (*SetVehicleLinkNavRes, error)`
  设置车辆链接导航

- `SetVehicleDestination(vehicleId string, destination *Point) (*SetVehicleDestinationRes, error)`
  设置车辆目的地

- `GetPedIdList() (*GetPedIdListRes, error)`
  获取行人 ID 列表

- `GetPedBaseInfo(pedIdList []string) (*GetPedBaseInfoRes, error)`
  获取行人基本信息

- `SetPedPosition(pedId string, point *Point, phi *float64) (*SetPedPositionRes, error)`
  设置行人位置

- `GetNMVIdList() (*GetNMVIdListRes, error)`
  获取非机动车 ID 列表

- `GetNMVBaseInfo(nmvIdList []string) (*GetNMVBaseInfoRes, error)`
  获取非机动车基本信息

- `SetNMVPosition(nmvId string, point *Point, phi *float64) (*SetNMVPositionRes, error)`
  设置非机动车位置

#### SimRecord

管理仿真记录。

- `GetRecordIds(scenId string, scenVer string) (*GetRecordIdsRes, error)`
  获取记录 ID 列表

- `GetTrackResults(id string, objId string) (*GetTrackResultsRes, error)`
  获取轨迹结果

- `GetSensorResults(id string, objId string) (*GetSensorResultsRes, error)`
  获取传感器结果

- `GetStepResults(id string, objId string) (*GetStepResultsRes, error)`
  获取步进结果

- `GetPathResults(id string, objId string) (*GetPathResultsRes, error)`
  获取路径结果

- `GetReferenceLineResults(id string, objId string) (*GetReferenceLineResultsRes, error)`
  获取参考线结果

## 系统要求

- Go >= 1.19

## 支持

如需报告错误或请求新功能，请使用[GitHub Issues](https://github.com/rl-lasvsim/openapi-sdk-go/issues)页面。

## 包结构

- **processtask**: 管理处理任务
- **resource**: 处理资源管理
- **simrecord**: 管理仿真记录
- **simulation**: 提供仿真功能
- **traintask**: 管理训练任务

## 目录结构

```
.
├── lasvsim/           # SDK主包
│   ├── client.go      # 主客户端实现
│   ├── go.mod         # Go模块定义
│   ├── httpclient/    # HTTP客户端实现
│   ├── processtask/   # 处理任务相关实现
│   ├── resource/      # 资源管理相关实现
│   ├── simrecord/     # 仿真记录相关实现
│   ├── simulation/    # 仿真功能实现
│   └── traintask/     # 训练任务相关实现
└── examples/           # 示例代码
    ├── go.mod         # 示例模块定义
    ├── go.sum         # 模块依赖校验
    ├── simulator_init_test.go  # 仿真器初始化测试
    └── manual_control/         # 手动控制示例
        └── main.go             # 主程序
```
