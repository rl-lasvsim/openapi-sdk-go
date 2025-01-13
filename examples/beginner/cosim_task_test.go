package beginner

import (
	"fmt"
	"os"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

/*
	TestCreateCosimTask 为帮助用户顺利使用千行仿真平台的联合仿真功能,编写一个联合仿真使用案例

该案例通过SDK创建一个联合仿真任务, 并控制仿真器运行, 最后在仿真任务结束后查看任务结果;
注意: 使用联合仿真功能, 用户需要注册好千行仿真平台账号, 最好对平台有一定的使用经验。并且会设置电脑的环境变量
*/
func TestCreateCosimTask(t *testing.T) {
	var (
		// "终端地址"的环境变量
		endpoint string = os.Getenv("QX_ENDPOINT") // 线上环境地址: https://qianxing-api.risenlighten.com
		// "token"的环境变量
		token string = os.Getenv("QX_TOKEN") // 登录仿真平台后访问https://qianxing.risenlighten.com/#/usecenter/personalCenter, 点击最下面按钮复制token

		// 登录仿真平台, 选择想要进行联合仿真的任务及剧本，赋值给下面的taskId和recordId变量
		// 任务ID
		taskId uint64 = 9988
		// 剧本ID
		recordId uint64 = 20776
	)

	// 1. 初始化客户端
	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: endpoint, // 设置的"终端地址"环境变量
		Token:    token,    // 设置的"token"环境变量
	})

	// 2. 拷贝剧本, 返回的结构中NewRecordId字段就是新创建的剧本ID, 仿真结束后可到该剧本下查看结果详情
	newRecord, err := cli.ProcessTask.CopyRecord(taskId, recordId)
	if err != nil {
		panic(err)
	}

	// 3. 通过拷贝的场景Id、Version和SimRecordId初始化仿真器
	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:      newRecord.ScenId,
		ScenVer:     newRecord.ScenVer,
		SimRecordID: newRecord.SimRecordId,
	})
	if err != nil {
		panic(err)
	}

	// 关闭仿真器, 释放服务器资源
	defer simulator.Stop()

	// 获取测试车辆列表
	testVehicleList, err := simulator.GetTestVehicleIdList()
	if err != nil {
		panic(err)
	}

	// 使测试车辆环形行驶
	for i := 0; i < 50; i++ {
		// 设置方向盘转角30度, 纵向加速度5
		var steWheel float64 = 10
		var lonAcc float64 = 0.05
		// 设置车辆的控制信息
		_, err := simulator.SetVehicleControlInfo(testVehicleList.List[0], &steWheel, &lonAcc)
		if err != nil {
			panic(err)
		}
		// 执行仿真器步骤
		simulator.Step()
	}

	// 可在此处继续调用其他接口, 查看联合仿真文档: https://www.risenlighten.com/#/union

	// 仿真结束后, 到千行仿真平台对应的taskId/recordId下查看联合仿真结果详情
	fmt.Printf("https://qianxing.risenlighten.com/#/configuration/circleTask?id=%d\n", taskId)

	// 如想直接查看本次联合仿真的回放视频, 可访问下面网址：
	fmt.Printf("https://qianxing.risenlighten.com/#/sampleRoad/cartest/?id=%d&record_id=%d&sim_record_id=%s\n", taskId, newRecord.NewRecordId, newRecord.SimRecordId)
}
