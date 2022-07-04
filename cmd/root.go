package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zngw",
	Short: "这是 cobra 测试程序主入口",
	Long:  `这是 cobra 测试程序主入口， 无参数启动时进入`,
	Run:   runRoot,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func runRoot(cmd *cobra.Command, args []string) {
	fmt.Printf("execute %s args:%v \n", cmd.Name(), args)
	// TODO 这里处理无参数启动时程序处理
}
