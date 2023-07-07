package demo

import (
	"log"

	"github.com/bref/outsider/framework/cobra"
)

func InitFoo() *cobra.Command {
	FooCommand.AddCommand(Foo1Command)
	return FooCommand
}

var FooCommand = &cobra.Command{
	Use:     "foo",
	Short:   "foo简要说明",
	Long:    "foo的长说明",
	Aliases: []string{"fo", "f"},
	Example: "foo命令的例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("excute command,every second")
		return nil
	},
}

var Foo1Command = &cobra.Command{
	Use:     "foo1",
	Short:   "foo1的简要说明",
	Long:    "foo1的长说明",
	Aliases: []string{"fo1", "f1"},
	Example: "foo命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		log.Println("Foo1Command,every 2 seconds")
		return nil
	},
}
