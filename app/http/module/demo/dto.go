package demo

import studentService "github.com/bref/outsider/app/provider/student"

type UserDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func userModelsToUserDTOs(models []UserModel) []UserDTO {
	ret := []UserDTO{}

	for _, model := range models {
		t := UserDTO{
			ID:   model.UserId,
			Name: model.Name,
		}
		ret = append(ret, t)
	}
	return ret
}

func StudentToUserDTOs(students []studentService.Student) []UserDTO {
	ret := []UserDTO{}

	for _, student := range students {
		t := UserDTO{
			ID:   student.ID,
			Name: student.Name,
		}
		ret = append(ret, t)
	}
	return ret
}
