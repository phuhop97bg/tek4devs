package user

import (
	"context"
	"tek4devs/common"
	"tek4devs/modules/task/entity"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertNewTask(ctx context.Context, data *entity.TaskCreationData) error
	GetTaskById(ctx context.Context, id string) (*entity.Task, error)
	FindTasks(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]entity.Task, error)
	UpdateTask(ctx context.Context, id string, data *entity.TaskPatchData) error
	DeleteTask(ctx context.Context, id string) error
}

type Business interface {
	CreateNewTask(ctx context.Context, data *entity.TaskCreationData) error
	ListTasks(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]entity.Task, error)
	GetTaskDetails(ctx context.Context, id string) (*entity.Task, error)
	UpdateTask(ctx context.Context, id string, data *entity.TaskPatchData) error
	DeleteTask(ctx context.Context, id string) error
}

type API interface {
	CreateTaskHandler() gin.HandlerFunc
	ListTaskHandler() gin.HandlerFunc
	GetTaskHandler() gin.HandlerFunc
	UpdateTaskHandler() gin.HandlerFunc
	DeleteTaskHandler() gin.HandlerFunc
}
