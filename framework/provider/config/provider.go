package config

import (
	"github.com/bref/outsider/framework"
	"github.com/bref/outsider/framework/contract"
)

type HadeConfigProvider struct {
	c      framework.Container
	folder string
	env    string

	envMaps map[string]string
}

func (provider *HadeConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeConfig
}

func (provider *HadeConfigProvider) Boot(c framework.Container) error {
	provider.folder = c.MustMake(contract.AppKey).(contract.App).ConfigFolder()
	provider.envMaps = c.MustMake(contract.EnvKey).(contract.EnvInterface).All()
	provider.env = c.MustMake(contract.EnvKey).(contract.EnvInterface).AppEnv()
	provider.c = c
	return nil
}

func (provider *HadeConfigProvider) IsDefer() bool {
	return true
}

func (provider *HadeConfigProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.folder, provider.envMaps, provider.env, provider.c}
}

func (provider *HadeConfigProvider) Name() string {
	return contract.ConfigKey
}
