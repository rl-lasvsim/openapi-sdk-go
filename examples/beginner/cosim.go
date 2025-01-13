package beginner

import (
	"fmt"
	"os"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

// cosim case
// 使用千行联合仿真SDk大概流程：
// 1. 初始化客户端
// 2. 访问千行平台, 选择想要进行联合仿真的任务及剧本, 拷贝其场景信息
// 3. 通过拷贝的场景Id和Version初始化仿真器
// 4. 调用Step()方法执行仿真步骤
// 5. 调用Stop()方法停止仿真器(必有)
// 6. 到对应的任务下查看联合仿真的结果
func NewCosimCase() error {
	// 初始化客户端
	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: os.Getenv("QX_ENDPOINT"), // 设置终端的环境变量("https://8.146.201.197:30080/dev")
		Token:    os.Getenv("QX_TOKEN"),    // 设置Authorization的环境变量
	})

	// 访问千行平台, 选择想要进行联合仿真的任务及剧本, 拷贝其场景信息
	scene, err := cli.ProcessTask.CopyRecord(TaskId, RecordId)
	if err != nil {
		return err
	}

	// 通过场景ID和Version初始化仿真器
	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:      scene.ScenId,
		ScenVer:     scene.ScenVer,
		SimRecordID: scene.SimRecordId,
	})
	if err != nil {
		return err
	}

	// 关闭仿真器, 释放服务器资源
	defer simulator.Stop()

	for i := 0; i < 10; i++ {
		simulator.Step()
		GetStepSpawnIdListRes, err := simulator.GetStepSpawnIdList()
		if err != nil {
			return err
		}

		fmt.Println(GetStepSpawnIdListRes.IdList)

		vehicleIDList, err := simulator.GetVehicleIdList()
		if err != nil {
			return err
		}

		fmt.Println(vehicleIDList.List)

		vehicleBaseInfo, err := simulator.GetVehicleBaseInfo(vehicleIDList.List)
		if err != nil {
			return err
		}

		fmt.Println(vehicleBaseInfo.InfoDict)

		// 可在此处继续调用仿真器的其他方法，具体方法查看api文档
	}

	return nil
}
