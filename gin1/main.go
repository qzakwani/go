package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qzakwani/go/gin1/settings"
	"github.com/qzakwani/go/gin1/task"
)

func main() {
	r := gin.Default()
	task.Routes(r)
	r.Run(settings.Port)
}
