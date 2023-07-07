package command

import "github.com/bref/outsider/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(initEnvCommand())

	//  app
	root.AddCommand(initAppCommand())

	//  cron命令
	root.AddCommand(initCronCommand())

	// config命令
	root.AddCommand(initConfigCommand())

	// provider命令
	root.AddCommand(initProviderCommand())

	// cmd命令
	root.AddCommand(initCmdCommand())

	// middleware命令
	root.AddCommand(initMiddlewareCommand())
}
