package task

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	route := r.Group("task/")

	route.GET("", ListTasks)
	route.GET(":id", GetTask)
	route.POST("", CreateTask)
	route.PUT(":id", EditTask)
	route.DELETE(":id", DeleteTask)

}
