package student

import (
	"github.com/bref/outsider/framework"
)

type StudentService struct {
	IStudentService
	container framework.Container
}

func NewStudentService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &StudentService{
		container: container,
	}, nil
}

func (s *StudentService) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}
