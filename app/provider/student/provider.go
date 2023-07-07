package student

import "github.com/bref/outsider/framework"

type StudentProvider struct {
	framework.ServiceProvider

	c framework.Container
}

func (sp *StudentProvider) Name() string {
	return StudentKey
}

func (sp *StudentProvider) Register(c framework.Container) framework.NewInstance {
	return NewStudentService
}

func (sp *StudentProvider) IsDefer() bool {
	return false
}

func (sp *StudentProvider) Params(c framework.Container) []interface{} {
	return []interface{}{sp.c}
}

func (sp *StudentProvider) Boot(c framework.Container) error {
	sp.c = c
	return nil
}
