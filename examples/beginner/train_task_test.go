package beginner

// train task case
import (
	"os"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

/*
	TestCreateTrainTask 一个训练任务SDK的使用案例

注意: 使用联合仿真功能, 用户需要注册好千行仿真平台账号, 最好对平台有一定的使用经验。并且会设置电脑的环境变量
*/
func TestCreateTrainTask(t *testing.T) {
	var (
		// "终端地址"的环境变量
		endpoint string = os.Getenv("QX_ENDPOINT") // 线上环境地址: https://qianxing-api.risenlighten.com
		// "token"的环境变量
		token string = os.Getenv("QX_TOKEN") // 登录仿真平台后访问https://qianxing.risenlighten.com/#/usecenter/personalCenter, 点击最下面按钮复制token

		// 登录训练平台, 选择训练任务, 将ID赋值给下面的taskId变量
		taskId uint64 = 515
	)

	// 初始化客户端
	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: endpoint, // 设置的"终端地址"环境变量
		Token:    token,    // 设置的"token"环境变量
	})

	// 访问千行平台，通过训练任务ID获取场景信息列表
	sceneIdList, err := cli.TrainTask.GetSceneIdList(taskId)
	if err != nil {
		panic(err)
	}

	// 通过场景ID和场景Version, 初始化仿真器(此案例默认使用场景列表中的第一个场景)
	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:  sceneIdList.SceneIdList[0],
		ScenVer: sceneIdList.SceneVersionList[0],
	})
	if err != nil {
		panic(err)
	}

	// 关闭仿真器, 释放服务器资源
	_ = simulator.Stop()
}
