package manualcontrol

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
)

// 控制配置结构
type ControlConfig struct {
	steWheelPerPress float64 //"a","d"
	lonAccPerPress   float64 //"s","w"
}

func NewControlConfig(steWheel, lonAcc float64) *ControlConfig {
	return &ControlConfig{steWheelPerPress: steWheel, lonAccPerPress: lonAcc}
}

type ManualControl struct {
	*lasvsim.Client
	*ControlConfig

	simulator        *simulation.Simulator
	controledVehicle string //控制的车辆
}

func NewManualControl(conf *ControlConfig) *ManualControl {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		// Endpoint: os.Getenv("QX_ENDPOINT"),
		// Token:    os.Getenv("QX_TOKEN"),
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjQ5LCJvaWQiOjI2LCJuYW1lIjoi5ZGo5YWr5Y-q6IO955yL5Zy65pmv5ZWK5ZWKIiwiaWRlbnRpdHkiOiJub3JtYWwiLCJwZXJtaXNzaW9ucyI6W10sImlzcyI6InVzZXIiLCJzdWIiOiJMYXNWU2ltIiwiZXhwIjoxNzM1NzAzOTA5LCJuYmYiOjE3MzUwOTkxMDksImlhdCI6MTczNTA5OTEwOSwianRpIjoiNDkifQ.A7MwADjN6DcBJ3QPLFTfO5ZU6a1fYyS-7AN0sEALun8",
		Endpoint: "http://8.146.201.197:30080/dev",
	})
	return &ManualControl{Client: cli, ControlConfig: conf}
}

func (m *ManualControl) Start(taskId, recordId uint64, controlVehicle string) {
	if controlVehicle == "" {
		controlVehicle = "测试车辆1"
	}
	m.controledVehicle = controlVehicle

	copyedRecord, err := m.Client.ProcessTask.CopyRecord(taskId, recordId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	simulator, err := m.Client.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:      copyedRecord.ScenId,
		ScenVer:     copyedRecord.ScenVer,
		SimRecordID: copyedRecord.SimRecordId,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 最后停止仿真
	defer simulator.Stop()
	m.runInteractive()
}
func (m *ManualControl) runInteractive() {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer s.Fini()

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)
	s.Clear()

	printStr(s, 0, 0, defStyle, "Enter input (w/s/a/d, press Ctrl+C to exit):")
	s.Show()

	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return
			}
			switch ev.Rune() {
			case 'w':
				m.simulator.SetVehicleControlInfo(m.controledVehicle, nil, &m.lonAccPerPress)
				printStr(s, 0, 1, defStyle, "Moving Down (s)")
			case 's':
				lonAcc := -m.lonAccPerPress
				m.simulator.SetVehicleControlInfo(m.controledVehicle, nil, &(lonAcc))
				printStr(s, 0, 1, defStyle, "Moving Down (s)")
			case 'a':
				steWheel := -m.steWheelPerPress
				m.simulator.SetVehicleControlInfo(m.controledVehicle, &(steWheel), nil)
				printStr(s, 0, 1, defStyle, "Moving Left (a)")
			case 'd':
				m.simulator.SetVehicleControlInfo(m.controledVehicle, &m.steWheelPerPress, nil)
				printStr(s, 0, 1, defStyle, "Moving Right (d)")
			default:
				printStr(s, 0, 1, defStyle, "Invalid input.")
			}
			s.Show()
		}
	}
}

func printStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, r := range str {
		s.SetContent(x, y, r, nil, style)
		x++
	}
}
