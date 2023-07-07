package demo

import (
	studentService "github.com/bref/outsider/app/provider/student"
	"github.com/bref/outsider/framework/contract"
	"github.com/bref/outsider/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()

	r.Bind(&studentService.StudentProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)
	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{
		service: service,
	}
}

func (api *DemoApi) Demo(c *gin.Context) {
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	password := configService.GetString("database.mysql.password")
	// 打印出来
	c.JSON(200, password)
}

func (api *DemoApi) Demo2(c *gin.Context) {
	studentProvider := c.MustMake(studentService.StudentKey).(studentService.IStudentService)
	students := studentProvider.GetAllStudent()
	usersDto := StudentToUserDTOs(students)
	c.JSON(200, usersDto)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}

	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}
