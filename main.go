package main

import (
	"github.com/bref/outsider/app/console"
	"github.com/bref/outsider/app/http"
	"github.com/bref/outsider/framework"
	"github.com/bref/outsider/framework/provider/app"
	"github.com/bref/outsider/framework/provider/config"
	"github.com/bref/outsider/framework/provider/distributed"
	"github.com/bref/outsider/framework/provider/env"
	"github.com/bref/outsider/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()

	// 绑定App服务提供者
	container.Bind(&app.OutsiderAppProvider{})
	// 绑定分布式锁服务提供者
	container.Bind(&distributed.LocalDistributedProvider{})
	// 绑定环境变量服务提供者
	container.Bind(&env.HadeEnvProvider{})
	// 绑定配置服务提供者
	container.Bind(&config.HadeConfigProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
		// 绑定完服务以后配置路由
		http.Routes(engine)
	}
	// 运行root命令
	console.RunCommand(container)

}
