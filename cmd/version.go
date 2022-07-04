package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version 子命令.",
	Long:  "这是一个version 子命令",
	Run:   runVersion,
}

var port *uint
var kafka *string

func init() {
	rootCmd.AddCommand(versionCmd)
	flags := pflag.NewFlagSet("flags", pflag.ExitOnError)
	port = flags.Uint("port", 0, "")
	kafka = flags.String("kafka_example", "xxx", "kafka地址")
	versionCmd.Flags().AddFlagSet(flags)
}

func runVersion(cmd *cobra.Command, args []string) {
	// TODO 这里处理version子命令

	fmt.Println("version is 1.0.0")
	fmt.Println("port is ", *port)
	fmt.Println("kafka_example address is", *kafka)
}
