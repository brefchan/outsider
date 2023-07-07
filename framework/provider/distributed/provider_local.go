package distributed

import (
	"github.com/bref/outsider/framework"
	"github.com/bref/outsider/framework/contract"
)

type LocalDistributedProvider struct{}

func (h *LocalDistributedProvider) Register(container framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

func (h *LocalDistributedProvider) Boot(container framework.Container) error {
	return nil
}

func (h *LocalDistributedProvider) IsDefer() bool {
	return false
}

func (h *LocalDistributedProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

func (h *LocalDistributedProvider) Name() string {
	return contract.DistributedKey
}
