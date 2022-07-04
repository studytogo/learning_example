package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello1",
	Short: "hello 子命令.",
	Long:  "这是一个Hello 子命令",
	Args:  cobra.MinimumNArgs(1),
	Run:   runHello,
}

func init() {
	rootCmd.AddCommand(helloCmd)
}

func runHello(cmd *cobra.Command, args []string) {
	// TODO 这里处理hello子命令

	fmt.Println("Hello ", args[0])
}
