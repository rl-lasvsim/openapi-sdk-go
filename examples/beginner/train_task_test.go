package beginner

// train task case
import (
	"os"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
	"github.com/stretchr/testify/assert"
)

/*
	TestCreateTrainTask 一个训练任务SDK的使用案例

注意: 使用联合仿真功能, 用户需要注册好千行仿真平台账号, 最好对平台有一定的使用经验。并且会设置电脑的环境变量
*/
func TestCreateTrainTask(t *testing.T) {
	var (
		// 接口地址
		endpoint string = os.Getenv("QX_ENDPOINT") // 线上环境地址: https://qianxing-api.risenlighten.com
		// 授权token
		token string = os.Getenv("QX_TOKEN") // 登录仿真平台后访问https://qianxing.risenlighten.com/#/usecenter/personalCenter, 点击最下面按钮复制token

		// 登录训练平台, 选择训练任务, 将训练任务ID赋值给下面的taskId变量
		taskId uint64 = 515
	)

	// 初始化客户端
	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: endpoint, // 接口地址
		Token:    token,    // 授权token
	})

	// 访问千行平台，通过训练任务ID获取场景信息列表
	sceneIdList, err := cli.TrainTask.GetSceneIdList(taskId)
	assert.NoError(t, err)

	// 通过场景ID和场景Version, 初始化仿真器(此案例默认使用场景列表中的第一个场景)
	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:  sceneIdList.SceneIdList[0],
		ScenVer: sceneIdList.SceneVersionList[0],
	})
	assert.NoError(t, err)

	// 关闭仿真器, 释放服务器资源
	_ = simulator.Stop()
}
