package console

import (
	"time"

	"github.com/bref/outsider/app/console/command/demo"
	"github.com/bref/outsider/framework"
	"github.com/bref/outsider/framework/cobra"
	"github.com/bref/outsider/framework/command"
)

func RunCommand(container framework.Container) error {
	// 根Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "outsider",
		// 简短介绍
		Short: "outsider 命令",
		// 根命令的详细介绍
		Long: "hade 框架提供的命令行工具,使用这个命令行工具能很方便执行框架自带命令,也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根Command设置服务容器
	rootCmd.SetContainer(container)
	//绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)

	return rootCmd.Execute()
}

// 绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddDistributedCronCommand("foo_func_for_test", "*/5 * * * * *", demo.FooCommand, 2*time.Second)
}
