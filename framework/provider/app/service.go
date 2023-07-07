package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/bref/outsider/framework"
	"github.com/bref/outsider/framework/contract"
	"github.com/bref/outsider/framework/util"
	"github.com/google/uuid"
)

type OutsiderAppService struct {
	contract.App
	container  framework.Container
	baseFolder string
	appId      string // 表示当前这个app唯一的id,可以用于分布式锁等

	configMap map[string]string // 配置加载

}

// Version 实现版本
func (h OutsiderAppService) Version() string {
	return "0.0.3"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (h OutsiderAppService) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}

	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (h OutsiderAppService) ConfigFolder() string {

	if val, ok := h.configMap["config_folder"]; ok {
		return val
	}
	return filepath.Join(h.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (h OutsiderAppService) LogFolder() string {
	if val, ok := h.configMap["log_folder"]; ok {
		return val
	}
	return filepath.Join(h.StorageFolder(), "log")
}

func (h OutsiderAppService) HttpFolder() string {
	if val, ok := h.configMap["http_folder"]; ok {
		return val
	}
	return filepath.Join(h.BaseFolder(), "app", "http")
}

func (h OutsiderAppService) ConsoleFolder() string {
	if val, ok := h.configMap["console_folder"]; ok {
		return val
	}
	return filepath.Join(h.BaseFolder(), "app", "console")
}

func (h OutsiderAppService) StorageFolder() string {
	if val, ok := h.configMap["storage_folder"]; ok {
		return val
	}
	return filepath.Join(h.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (h OutsiderAppService) ProviderFolder() string {
	if val, ok := h.configMap["provider_folder"]; ok {
		return val
	}
	return filepath.Join(h.BaseFolder(), "app", "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (h OutsiderAppService) MiddlewareFolder() string {
	if val, ok := h.configMap["middleware_folder"]; ok {
		return val
	}
	return filepath.Join(h.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (h OutsiderAppService) CommandFolder() string {
	if val, ok := h.configMap["command_folder"]; ok {
		return val
	}
	return filepath.Join(h.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (h OutsiderAppService) RuntimeFolder() string {
	if val, ok := h.configMap["runtime_folder"]; ok {
		return val
	}
	return filepath.Join(h.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (h OutsiderAppService) TestFolder() string {
	if val, ok := h.configMap["test_folder"]; ok {
		return val
	}
	return filepath.Join(h.BaseFolder(), "test")
}

func (h OutsiderAppService) AppID() string {
	return h.appId
}

// NewHadeApp 初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	// 如果没有设置，则使用参数
	if baseFolder == "" {
		flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
		flag.Parse()
	}

	appId := uuid.New().String()
	return &OutsiderAppService{baseFolder: baseFolder, container: container, appId: appId}, nil
}

func (app *OutsiderAppService) LoadAppConfig(kv map[string]string) {
	app.configMap = kv
}
