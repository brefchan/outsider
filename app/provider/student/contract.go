package student

const StudentKey = "student"

type IStudentService interface {
	GetAllStudent() []Student
}

type Student struct {
	ID   int
	Name string
}
