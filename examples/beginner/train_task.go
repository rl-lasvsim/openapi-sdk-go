package beginner

// 新建一个训练任务对象
// 对象下有仿真器相关的方法
import (
	"os"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

var (
	taskId   uint64 = 1
	recordId uint64 = 1
)

func TrainTaskCase() error {
	// 初始化客户端
	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: os.Getenv("QX_ENDPOINT"),
		Token:    os.Getenv("QX_TOKEN"),
	})

	// 通过任务ID获取场景ID列表
	_, err := cli.TrainTask.GetSceneIdList(taskId)
	if err != nil {
		return err
	}

	cli.InitSimulatorFromConfig(simulation.SimulatorConfig{})
	return nil
}
