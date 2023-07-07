package command

import (
	"fmt"

	"github.com/bref/outsider/framework/cobra"
	"github.com/bref/outsider/framework/contract"
	"github.com/kr/pretty"
)

func initConfigCommand() *cobra.Command {
	configCommand.AddCommand(configGetCommand)
	return configCommand
}

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "配置控制命令",
	Long:  "应用配置相关命令,包括查询,更新,删除等功能",
	RunE: func(c *cobra.Command, args []string) error {
		c.Help()
		return nil
	},
}

var configGetCommand = &cobra.Command{
	Use:   "get",
	Short: "配置获取命令",
	Long:  "获取指定path路径的配置",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			fmt.Println("参数错误")
			return nil
		}

		container := cmd.GetContainer()
		configSevice := container.MustMake(contract.ConfigKey).(contract.Config)

		result := configSevice.Get(args[0])
		if result == nil {
			fmt.Println("配置路径 ", args[0], " 不存在")
			return nil
		}

		fmt.Printf("%# v\n", pretty.Formatter(result))

		return nil
	},
}
