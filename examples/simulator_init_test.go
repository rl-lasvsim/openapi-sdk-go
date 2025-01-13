package main

import (
	"os"
	"testing"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	httpclient "github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	simulation "github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

func TestSimulatorInit(t *testing.T) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: os.Getenv("QX_ENDPOINT"),
		Token:    os.Getenv("QX_TOKEN"),
	})

	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:  "545339363676944979",
		ScenVer: "2",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err = simulator.Stop()
		if err != nil {
			t.Fatal(err)
		}
	}()

	getTestVehicleIdListRes, err := simulator.GetTestVehicleIdList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("getTestVehicleIdListRes:", getTestVehicleIdListRes)

	// GetVehicleIdListRes, err := simulator.GetVehicleIdList()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(GetVehicleIdListRes)

	// GetVehicleBaseInfoRes, err := simulator.GetVehicleBaseInfo(GetVehicleIdListRes.List)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(GetVehicleBaseInfoRes)
	// t.Log(GetVehicleBaseInfoRes)

	// GetVehicleMovingInfoRes, err := simulator.GetVehicleMovingInfo(GetVehicleIdListRes.List)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(GetVehicleMovingInfoRes)

	// GetVehicleTargetSpeedRes, err := simulator.GetVehicleTargetSpeed(GetVehicleIdListRes.List[0])
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(GetVehicleTargetSpeedRes)

	// GetVehicleControlInfoRes, err := simulator.GetVehicleControlInfo(GetVehicleIdListRes.List)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(GetVehicleControlInfoRes)

	// GetVehicleCollisionStatusRes, err := simulator.GetVehicleCollisionStatus(GetVehicleIdListRes.List[0])
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(GetVehicleCollisionStatusRes)

	// GetVehiclePerceptionInfoRes, err := simulator.GetVehiclePerceptionInfo(GetVehicleIdListRes.List[0])
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(GetVehiclePerceptionInfoRes)

	// GetVehicleNavigationInfo, err := simulator.GetVehicleNavigationInfo(GetVehicleIdListRes.List[0])
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(GetVehicleNavigationInfo)

	// for i := 0; i < 10; i++ {
	// 	res, err := simulator.Step()
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	t.Log(res)
	// }

	// SetVehicleDestinationRes, err := simulator.SetVehicleDestination("测试车辆1", &simulation.Point{0.1, 0.1, 0.1})
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(SetVehicleDestinationRes)

	// SetVehicleBaseInfoRes, err := simulator.SetVehicleBaseInfo("测试车辆1", &simulation.ObjBaseInfo{0.1, 0.1, 0.1, 0.1}, nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(SetVehicleBaseInfoRes)

	// var acc float64 = 0.11
	// SetVehicleControlInfoRes, err := simulator.SetVehicleControlInfo("测试车辆1", &acc, &acc)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(SetVehicleControlInfoRes)

	// step测试
	for i := 0; i < 10; i++ {
		GetStepSpawnIdListRes, err := simulator.GetStepSpawnIdList()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(GetStepSpawnIdListRes)
		GetParticipantBaseInfoRes, err := simulator.GetParticipantBaseInfo(GetStepSpawnIdListRes.IdList)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(*GetParticipantBaseInfoRes.BaseInfoDict["测试车辆1"])
		GetParticipantMovingInfoRes, err := simulator.GetParticipantMovingInfo(GetStepSpawnIdListRes.IdList)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(*GetParticipantMovingInfoRes.MovingInfoDict["环境车辆1"])
		GetParticipantPositionRes, err := simulator.GetParticipantPosition(GetStepSpawnIdListRes.IdList)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(*GetParticipantPositionRes.PositionDict["环境车辆1"])
	}
}
