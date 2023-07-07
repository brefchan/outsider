package app

import (
	"github.com/bref/outsider/framework"
	"github.com/bref/outsider/framework/contract"
)

type OutsiderAppProvider struct {
	BaseFolder string
}

// Register 注册HadeApp方法
func (h *OutsiderAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeApp
}

// Boot 启动调用
func (h *OutsiderAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *OutsiderAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *OutsiderAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *OutsiderAppProvider) Name() string {
	return contract.AppKey
}
