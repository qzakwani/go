package task

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qzakwani/go/gin1/core"
)

var tasks = make(map[int]Task)

func CreateTask(ctx *gin.Context) {
	id := rand.Intn(100)
	println(id)
	tasks[id] = Task{
		Id:        id,
		Text:      ctx.PostForm("task"),
		CreatedAt: core.DtNow(),
	}

	ctx.JSON(http.StatusAccepted, tasks[id])
}

func GetTask(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 0)
	ctx.JSON(http.StatusOK, tasks[int(id)])
}

func EditTask(ctx *gin.Context) {
	_id, _ := strconv.ParseInt(ctx.Param("id"), 10, 0)
	id := int(_id)
	tasks[id] = Task{
		Id:        id,
		Text:      ctx.PostForm("task"),
		CreatedAt: tasks[id].CreatedAt,
	}
	ctx.JSON(http.StatusOK, tasks[id])
}

func DeleteTask(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 0)
	delete(tasks, int(id))
	ctx.Status(http.StatusOK)
}

func ListTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tasks)
}
