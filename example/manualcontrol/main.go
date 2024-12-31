package manualcontrol

// import (
// 	"fmt"
// 	"os"

// 	"github.com/gdamore/tcell"
// 	_ "github.com/rl-lasvsim/openapi-sdk-go/lasvsim"
// 	_ "github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
// 	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/simulation"
// 	"github.com/spf13/cobra"
// )

// func main() {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

// var rootCmd = &cobra.Command{
// 	Short: "A CLI tool that handles w/s/a/d input to control lasvsim simulation",
// 	Run:   run,
// }

// func run(cmd *cobra.Command, args []string) {
// 	// cli := lasvsim.NewClient(&httpclient.HttpConfig{
// 	// 	Endpoint: os.Getenv("QX_ENDPOINT"),
// 	// 	Token:    os.Getenv("QX_TOKEN"),
// 	// })
// 	// taskId := uint64(0)
// 	// recordId := uint64(0)

// 	// newRecord, err := cli.ProcessTask.CopyRecord(taskId, recordId)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	os.Exit(1)
// 	// }

// 	// simulator, err := cli.InitSimulatorFromConfig(simulation.SimulatorConfig{
// 	// 	ScenID:      newRecord.ScenId,
// 	// 	ScenVer:     newRecord.ScenVer,
// 	// 	SimRecordID: newRecord.SimRecordId,
// 	// })
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	os.Exit(1)
// 	// }
// 	// // 最后停止仿真
// 	// defer simulator.Stop()
// 	runInteractive(nil)
// }

// func runInteractive(simulator *simulation.Simulator) {
// 	s, err := tcell.NewScreen()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "%v\n", err)
// 		os.Exit(1)
// 	}
// 	if err := s.Init(); err != nil {
// 		fmt.Fprintf(os.Stderr, "%v\n", err)
// 		os.Exit(1)
// 	}
// 	defer s.Fini()

// 	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
// 	s.SetStyle(defStyle)
// 	s.Clear()

// 	printStr(s, 0, 0, defStyle, "Enter input (w/s/a/d, press Ctrl+C to exit):")
// 	s.Show()

// 	for {
// 		ev := s.PollEvent()
// 		switch ev := ev.(type) {
// 		case *tcell.EventKey:
// 			if ev.Key() == tcell.KeyCtrlC {
// 				return
// 			}
// 			switch ev.Rune() {
// 			case 'w':
// 				printStr(s, 0, 1, defStyle, "Moving Up (w)")
// 			case 's':
// 				printStr(s, 0, 1, defStyle, "Moving Down (s)")
// 			case 'a':
// 				printStr(s, 0, 1, defStyle, "Moving Left (a)")
// 			case 'd':
// 				printStr(s, 0, 1, defStyle, "Moving Right (d)")
// 			default:
// 				printStr(s, 0, 1, defStyle, "Invalid input.")
// 			}
// 			s.Show()
// 		}
// 	}
// }

// func printStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
// 	for _, r := range str {
// 		s.SetContent(x, y, r, nil, style)
// 		x++
// 	}
// }

// //
// // func HandleUp(simulator *simulation.Simulator,)
