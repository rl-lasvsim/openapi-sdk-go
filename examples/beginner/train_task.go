package beginner

// train task case
import (
	"fmt"
	"os"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

var (
	TaskId   uint64 = 14362
	RecordId uint64 = 14362
)

// case
// 1. 初始化客户端
// 2. 根据训练任务获取场景id和version列表
// 3. 选用其中一个场景id和version初始化仿真器
// 4. 调用Step()方法执行仿真步骤
// 5. 调用Stop()方法停止仿真器(必有)
func NewTrainTaskCase() error {
	// 初始化客户端
	var cli = lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: os.Getenv("QX_ENDPOINT"), // 设置终端的环境变量("https://8.146.201.197:30080/dev")
		Token:    os.Getenv("QX_TOKEN"),    // 设置Authorization的环境变量
	})

	// 访问千行平台，通过训练任务ID获取场景信息列表
	sceneIdList, err := cli.TrainTask.GetSceneIdList(TaskId)
	if err != nil {
		return err
	}

	// 通过场景ID和场景Version，初始化仿真器
	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:  sceneIdList.SceneIdList[0],
		ScenVer: sceneIdList.SceneVersionList[0],
	})
	if err != nil {
		return err
	}

	// 关闭仿真器, 释放服务器资源
	defer simulator.Stop()

	// 获取车辆列表
	vechicleList, err := simulator.GetVehicleIdList()
	if err != nil {
		return err
	}
	fmt.Printf("获取仿真车辆ID列表结果:%+v\n", vechicleList)

	// 获取测试车辆列表
	testVechicleList, err := simulator.GetTestVehicleIdList()
	if err != nil {
		return err
	}
	fmt.Printf("获取测试仿真车辆ID列表结果:%+v\n", testVechicleList)

	// 获取车辆的基础信息, 传参是车辆id列表, []string类型
	vehicleBaseInfo, err := simulator.GetVehicleBaseInfo(vechicleList.List)
	if err != nil {
		return err
	}
	fmt.Printf("根据车辆列表获取车辆的基础信息结果:%+v\n", vehicleBaseInfo)

	// 获取训练过程中的仿真日志,需要先获取一个日志ID
	// 根据场景ID和版本获取日志ID列表，日志ID列表与剧本列表一一对应
	getRecordIds, err := cli.SimRecord.GetRecordIds(sceneIdList.SceneIdList[0], sceneIdList.SceneIdList[0])
	if err != nil {
		return err
	}
	fmt.Printf("根据场景ID和版本获取日志ID列表结果:%+v\n", getRecordIds)

	// 循环运行50步
	for i := 0; i < 500; i++ {
		stepRes, err := simulator.Step()
		if err != nil {
			return err
		}
		fmt.Printf("执行仿真步骤结果,(0:运行中;1001:正常结束;1002;未通过):%+v\n", stepRes)
		if stepRes.Code.IsStoped() || stepRes.Code.IsUnpassed() {
			break
		}

		// 获取车辆在仿真过程中的运动信息, 传参是车辆id列表, []string类型
		vehicleBaseInfo, err := simulator.GetVehicleMovingInfo(vechicleList.List)
		if err != nil {
			return err
		}
		fmt.Printf("根据车辆列表获取车辆的运动信息结果:%+v\n", vehicleBaseInfo)

		// 根据日志ID获取指定对象的轨迹结果
		getTrackResults, err := cli.SimRecord.GetTrackResults(getRecordIds.Ids[0], "测试车辆1")
		if err != nil {
			return err
		}
		fmt.Printf("根据日志ID获取指定对象的轨迹结果:%+v\n", getTrackResults)

		// 根据场景ID获取指定对象的传感器结果
		getSensorResults, err := cli.SimRecord.GetSensorResults(getRecordIds.Ids[0], "测试车辆1")
		if err != nil {
			return err
		}
		fmt.Printf("根据场景ID获取指定对象的传感器结果:%+v\n", getSensorResults)

		// 根据日志ID获取指定对象的瞬时结果
		getStepResultsRes, err := cli.SimRecord.GetStepResults(getRecordIds.Ids[0], "测试车辆1")
		if err != nil {
			return err
		}
		fmt.Printf("根据日志ID获取指定对象的瞬时结果:%+v\n", getStepResultsRes)

		// 根据日志ID获取指定对象的路径结果
		getPathResults, err := cli.SimRecord.GetPathResults(getRecordIds.Ids[0], "测试车辆1")
		if err != nil {
			return err
		}
		fmt.Printf("根据日志ID获取指定对象的路径结果:%+v\n", getPathResults)

		// 根据日志ID获取指定对象的参考线结果
		getReferenceLineResultsRes, err := cli.SimRecord.GetReferenceLineResults(getRecordIds.Ids[0], "测试车辆1")
		if err != nil {
			return err
		}
		fmt.Printf("根据日志ID获取指定对象的参考线结果:%+v\n", getReferenceLineResultsRes)
	}

	return nil
}
