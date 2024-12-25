package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
	"github.com/spf13/cobra"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Short: "A CLI tool that handles w/s/a/d input to control lasvsim simulation",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	cli := lasvsim.NewClient(&httpclient.HttpConfig{
		Endpoint: os.Getenv("QX_ENDPOINT"),
		Token:    os.Getenv("QX_TOKEN"),
	})
	taskId := uint64(0)
	recordId := uint64(0)

	newRecord, err := cli.ProcessTask.CopyRecord(taskId, recordId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter input (w/s/a/d, press Ctrl+C to exit):")

	simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
		ScenID:      newRecord.ScenId,
		ScenVer:     newRecord.ScenVer,
		SimRecordID: newRecord.SimRecordId,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 最后停止仿真
	defer simulator.Stop()

	for {
		input, _ := reader.ReadString('\n')
		// 去除换行符
		input = input[:len(input)-1]

		switch input {
		case "w":
			fmt.Println("Moving Up (w)")
			// 执行向上移动的逻辑
		case "s":
			fmt.Println("Moving Down (s)")
			// 执行向下移动的逻辑
		case "a":
			fmt.Println("Moving Left (a)")
			// 执行向左移动的逻辑
		case "d":
			fmt.Println("Moving Right (d)")
			// 执行向右移动的逻辑
		case "": // 处理只按回车的情况
			continue
		default:
			fmt.Println("Invalid input. Please use w, s, a, or d.")
		}
	}
}
