package command

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jianfengye/collection"
	"github.com/pkg/errors"

	"github.com/bref/outsider/framework/cobra"
	"github.com/bref/outsider/framework/contract"
	"github.com/bref/outsider/framework/util"
)

func initProviderCommand() *cobra.Command {
	providerCommand.AddCommand(providerListCommand)
	providerCommand.AddCommand(providerCreateCommand)
	return providerCommand
}

var providerCommand = &cobra.Command{
	Use:   "provider",
	Short: "服务提供相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

var providerListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出容器内的所有服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()

		providerList := container.ProviderList()
		for _, provider := range providerList {
			println(provider.Name())
		}
		return nil
	},
}

var providerCreateCommand = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create", "init"},
	Short:   "创建一个服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		fmt.Println("创建一个服务")
		var name string
		var folder string

		{
			prompt := &survey.Input{
				Message: "请输入服务名称(服务凭证)",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				return err
			}
		}

		{
			prompt := &survey.Input{
				Message: "请输入服务所在目录名称(默认:同服务名称)",
			}
			err := survey.AskOne(prompt, &folder)
			if err != nil {
				return err
			}
		}

		// 检查服务是否存在
		providers := container.ProviderList()
		for _, provider := range providers {
			if name == provider.Name() {
				fmt.Println("服务名称已经存在")
				return nil
			}
		}

		if folder == "" {
			folder = name
		}

		app := container.MustMake(contract.AppKey).(contract.App)

		pFolder := app.ProviderFolder()
		subFolders, err := util.SubDir(pFolder)
		if err != nil {
			return err
		}

		subColl := collection.NewStrCollection(subFolders)
		if subColl.Contains(folder) {
			fmt.Println("目录名称已经存在")
			return nil
		}

		// 开始创建文件
		if err := os.Mkdir(filepath.Join(pFolder, folder), 0700); err != nil {
			return err
		}

		// 创建title这个模版方法
		funcs := template.FuncMap{"title": strings.Title}
		{
			// 创建contract.go
			file := filepath.Join(pFolder, folder, "contract.go")
			f, err := os.Create(file)
			if err != nil {
				return errors.Cause(err)
			}

			// 使用contractTmp模版来初始化template,并且让这个模版支持title方法,即支持{{.|title}}
			t := template.Must(template.New("contract").Funcs(funcs).Parse(contractTmp))
			// 将name传递到template中渲染,并且输出到contract.go中
			if err := t.Execute(f, name); err != nil {
				return errors.Cause(err)
			}
		}
		{
			// 创建provider.go
			file := filepath.Join(pFolder, folder, "provider.go")
			f, err := os.Create(file)
			if err != nil {
				return err
			}
			t := template.Must(template.New("provider").Funcs(funcs).Parse(providerTmp))
			if err := t.Execute(f, name); err != nil {
				return err
			}
		}
		{
			//  创建service.go
			file := filepath.Join(pFolder, folder, "service.go")
			f, err := os.Create(file)
			if err != nil {
				return err
			}
			t := template.Must(template.New("service").Funcs(funcs).Parse(serviceTmp))
			if err := t.Execute(f, name); err != nil {
				return err
			}
		}

		fmt.Println("创建服务成功, 文件夹地址:", filepath.Join(pFolder, folder))
		fmt.Println("请不要忘记挂载新创建的服务")

		return nil
	},
}

var contractTmp string = `
package {{.}}

const {{.|title}}Key = "{{.}}"

type Service interface {
	// 请在这里定义你的方法
	// Foo() string
}
`

var providerTmp string = `
package {{.}}

import "github.com/bref/outsider/framework"

type {{.|title}}Provider struct {
	framework.ServiceProvider

	c framework.Container
}

func (sp *{{.|title}}Provider) Name() string{
	return {{.|title}}Key
}

func (sp *{{.|title}}Provider) Register(c framework.Container) framework.NewInstance {
	return New{{.|title}}Service
}

func (sp *{{.|title}}Provider) IsDefer() bool {
	return false
}

func (sp *{{.|title}}Provider) Params(c framework.Container) []interface{}{
	return []interface{}{c}
}

func (sp *{{.|title}}Provider) Boot(c framework.Container) error{
	return nil
}
`

var serviceTmp string = `
package {{.}}

import "github.com/bref/outsider/framework"

type {{.|title}}Service struct {
	container framework.Container
}

func New{{.|title}}Service(params ...interface{}) (interface{},error){
	container := params[0].(framework.Container)
	return &{{.|title}}Service{container:container},nil
}


`
