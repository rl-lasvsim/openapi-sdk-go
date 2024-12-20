package example

import (
	"fmt"
	"testing"

	qianxing "github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	httpclient "github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	simulation "github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

func TestSimulatorInit(t *testing.T) {
	cli := qianxing.NewClient(&httpclient.HttpConfig{
		// Endpoint: os.Getenv("QX_ENDPOINT"),
		// Token:    os.Getenv("QX_TOKEN"),
		Endpoint: "http://8.146.201.197:30080/dev",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjI3LCJvaWQiOjI1LCJuYW1lIjoi5byg5LiJIiwiaWRlbnRpdHkiOiJub3JtYWwiLCJwZXJtaXNzaW9ucyI6W10sImlzcyI6InVzZXIiLCJzdWIiOiJMYXNWU2ltIiwiZXhwIjoxNzM1MTk3NzkzLCJuYmYiOjE3MzQ1OTI5OTMsImlhdCI6MTczNDU5Mjk5MywianRpIjoiMjcifQ.NErDniRRuZFsvG7Dn623E4CHCIxjqs-7KvMPQ1E0WCI",
	})

	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:  "78533880676610",
		ScenVer: "0",
	})
	if err != nil {
		t.Fatal(err)
	}

	getTestVehicleIdListRes, err := simulator.GetTestVehicleIdList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("getTestVehicleIdListRes:", getTestVehicleIdListRes)

	GetVehicleIdListRes, err := simulator.GetVehicleIdList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetVehicleIdListRes)

	GetVehicleBaseInfoRes, err := simulator.GetVehicleBaseInfo(GetVehicleIdListRes.List)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(GetVehicleBaseInfoRes)
	t.Log(GetVehicleBaseInfoRes)

	GetVehicleMovingInfoRes, err := simulator.GetVehicleMovingInfo(GetVehicleIdListRes.List)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetVehicleMovingInfoRes)

	GetVehicleTargetSpeedRes, err := simulator.GetVehicleTargetSpeed(GetVehicleIdListRes.List[0])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetVehicleTargetSpeedRes)

	GetVehicleControlInfoRes, err := simulator.GetVehicleControlInfo(GetVehicleIdListRes.List)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetVehicleControlInfoRes)

	GetVehicleCollisionStatusRes, err := simulator.GetVehicleCollisionStatus(GetVehicleIdListRes.List[0])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetVehicleCollisionStatusRes)

	GetVehiclePerceptionInfoRes, err := simulator.GetVehiclePerceptionInfo(GetVehicleIdListRes.List[0])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetVehiclePerceptionInfoRes)

	GetVehicleNavigationInfo, err := simulator.GetVehicleNavigationInfo(GetVehicleIdListRes.List[0])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetVehicleNavigationInfo)

	for i := 0; i < 50; i++ {
		res, err := simulator.Step()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(res)
	}

	err = simulator.Stop()
	if err != nil {
		t.Fatal(err)
	}

	// resource 测试
	hdMapRes, err := cli.Resources.GetHdMap("78533880676610", "0")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*hdMapRes.Data)
}
