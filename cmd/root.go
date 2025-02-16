package cmd

import (
	"fmt"

	"github.com/deviceinsight/kafkactl/internal/global"

	"github.com/deviceinsight/kafkactl/cmd/alter"
	"github.com/deviceinsight/kafkactl/cmd/attach"
	"github.com/deviceinsight/kafkactl/cmd/clone"
	"github.com/deviceinsight/kafkactl/cmd/config"
	"github.com/deviceinsight/kafkactl/cmd/consume"
	"github.com/deviceinsight/kafkactl/cmd/create"
	"github.com/deviceinsight/kafkactl/cmd/deletion"
	"github.com/deviceinsight/kafkactl/cmd/describe"
	"github.com/deviceinsight/kafkactl/cmd/get"
	"github.com/deviceinsight/kafkactl/cmd/produce"
	"github.com/deviceinsight/kafkactl/cmd/reset"
	"github.com/deviceinsight/kafkactl/internal/k8s"
	"github.com/deviceinsight/kafkactl/output"
	"github.com/spf13/cobra"
)

func NewKafkactlCommand(streams output.IOStreams) *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "kafkactl",
		Short: "command-line interface for Apache Kafka",
		Long:  `A command-line interface the simplifies interaction with Kafka.`,
	}

	globalConfig := global.NewConfig()

	cobra.OnInitialize(globalConfig.Init)

	rootCmd.AddCommand(config.NewConfigCmd())
	rootCmd.AddCommand(consume.NewConsumeCmd())
	rootCmd.AddCommand(create.NewCreateCmd())
	rootCmd.AddCommand(alter.NewAlterCmd())
	rootCmd.AddCommand(deletion.NewDeleteCmd())
	rootCmd.AddCommand(describe.NewDescribeCmd())
	rootCmd.AddCommand(get.NewGetCmd())
	rootCmd.AddCommand(produce.NewProduceCmd())
	rootCmd.AddCommand(reset.NewResetCmd())
	rootCmd.AddCommand(attach.NewAttachCmd())
	rootCmd.AddCommand(clone.NewCloneCmd())
	rootCmd.AddCommand(newCompletionCmd())
	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newDocsCmd())

	globalFlags := globalConfig.Flags()

	// use upper-case letters for shorthand params to avoid conflicts with local flags
	rootCmd.PersistentFlags().StringVarP(&globalFlags.ConfigFile, "config-file", "C", "",
		fmt.Sprintf("config file. default locations: %v", globalConfig.DefaultPaths()))
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Verbose, "verbose", "V", false, "verbose output")

	k8s.KafkaCtlVersion = Version

	output.IoStreams = streams
	rootCmd.SetOut(streams.Out)
	rootCmd.SetErr(streams.ErrOut)
	return rootCmd
}
