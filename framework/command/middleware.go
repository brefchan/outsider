package command

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bref/outsider/framework/cobra"
	"github.com/bref/outsider/framework/contract"
	"github.com/bref/outsider/framework/util"
	"github.com/go-git/go-git/v5"
	"github.com/jianfengye/collection"
	"github.com/pkg/errors"
)

func initMiddlewareCommand() *cobra.Command {
	middlewareCommand.AddCommand(middlewareAllCommand)
	middlewareCommand.AddCommand(middlewareCreateCommand)
	middlewareCommand.AddCommand(middlewareMigrateCommand)

	return middlewareCommand
}

var middlewareCommand = &cobra.Command{
	Use:   "middleware",
	Short: "中间件相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

// 显示所有安装的中间件
var middlewareAllCommand = &cobra.Command{
	Use:   "list",
	Short: "显示所有中间件",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		// 读取文件夹
		files, err := util.SubDir(appService.MiddlewareFolder())
		if err != nil {
			return err
		}

		for _, middlewareName := range files {
			fmt.Println(middlewareName)
		}
		return nil
	},
}

var middlewareMigrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "迁移gin-contrib中间件,迁移地址:https://github.com/gin-contrib/[middleware].git",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		fmt.Println("迁移一个Gin中间件")
		var repo string
		{
			prompt := &survey.Input{
				Message: "请输入中间件名称 ",
			}
			err := survey.AskOne(prompt, &repo)
			if err != nil {
				return err
			}
		}

		// 下载git到一个目录中
		appService := container.MustMake(contract.AppKey).(contract.App)

		middlewarePath := appService.MiddlewareFolder()
		url := "https://github.com/gin-contrib/" + repo + ".git"

		fmt.Println("下载中间件 gin-contrib:")
		fmt.Println(url)

		_, err := git.PlainClone(path.Join(middlewarePath, repo), false, &git.CloneOptions{
			URL:      url,
			Progress: os.Stdout,
		})

		if err != nil {
			return err
		}

		// step3:删除不必要的文件 go.mod,go.sum, .git
		repoFolder := path.Join(middlewarePath, repo)
		fmt.Println("remove " + path.Join(repoFolder, "go.mod"))
		os.Remove(path.Join(repoFolder, "go.mod"))
		fmt.Println("remove " + path.Join(repoFolder, "go.sum"))
		os.Remove(path.Join(repoFolder, "go.sum"))
		fmt.Println("remove " + path.Join(repoFolder, ".git"))
		os.RemoveAll(path.Join(repoFolder, ".git"))

		filepath.Walk(repoFolder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if filepath.Ext(path) != ".go" {
				return nil
			}

			c, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			isContain := bytes.Contains(c, []byte("github.com/gin-gonic/gin"))
			if isContain {
				fmt.Println("更新文件:" + path)
				c = bytes.ReplaceAll(c, []byte("github.com/gin-gonic/gin"), []byte("github.com/bref/outside/framework/gin"))
				err = ioutil.WriteFile(path, c, 0644)
				if err != nil {
					return err
				}
			}
			return nil
		})
		return nil
	},
}

var middlewareCreateCommand = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create", "init"},
	Short:   "创建一个中间件",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		fmt.Println("创建一个中间件")
		var name string
		var folder string
		{
			prompt := &survey.Input{
				Message: "请输入中间件名称: ",
			}
			err := survey.AskOne(prompt, &name)

			if err != nil {
				return err
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入中间件所在目录名称(默认:同中间件名称)",
			}

			err := survey.AskOne(prompt, &folder)
			if err != nil {
				return err
			}
		}

		if folder == "" {
			folder = name
		}

		app := container.MustMake(contract.AppKey).(contract.App)

		pFolder := app.MiddlewareFolder()
		subFolders, err := util.SubDir(pFolder)

		if err != nil {
			return nil
		}

		subColl := collection.NewStrCollection(subFolders)
		if subColl.Contains(folder) {
			fmt.Println("目录已经存在")
			return nil
		}

		// 开始创建文件
		if err := os.Mkdir(filepath.Join(pFolder, folder), 0700); err != nil {
			return err
		}

		funcs := template.FuncMap{"title": strings.Title}

		{
			// 创建
			file := filepath.Join(pFolder, folder, "middleware.go")
			f, err := os.Create(file)
			if err != nil {
				return errors.Cause(err)
			}

			t := template.Must(template.New("middleware").Funcs(funcs).Parse(middlewareTmp))
			if err := t.Execute(f, name); err != nil {
				return errors.Cause(err)
			}
		}
		fmt.Println("创建中间件成功,文件夹地址:", filepath.Join(pFolder, folder))
		return nil
	},
}

var middlewareTmp string = `
package {{.}}

import "github.com/bref/outsider/framework/gin"

// {{.|title}}Middleware 中间件函数
func {{.|title}}Middleware() gin.HandlerFunc{
	return func(context *gin.Context){
		context.Next()
	}
}
`
