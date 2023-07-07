package rest

import (
	"net/http"
	"tek4devs/common"
	"tek4devs/modules/task"
	"tek4devs/modules/task/entity"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz task.Business
}

func NewAPI(biz task.Business) *api {
	return &api{biz: biz}
}

func (api *api) CreateTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.TaskCreationData

		if err := c.ShouldBind(&data); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := api.biz.CreateNewTask(c.Request.Context(), &data); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}

func (api *api) ListTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData struct {
			entity.Filter
			common.Paging
		}

		if err := c.ShouldBind(&requestData); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := api.biz.ListTasks(c.Request.Context(), &requestData.Filter, &requestData.Paging)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
			//"paging": requestData.Paging,
			//"extra":  requestData.Filter,
		})
	}
}

func (api *api) GetTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		data, err := api.biz.GetTaskDetails(c.Request.Context(), id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (api *api) UpdateTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data entity.TaskPatchData

		if err := c.ShouldBind(&data); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := api.biz.UpdateTask(c.Request.Context(), id, &data)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func (api *api) DeleteTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := api.biz.DeleteTask(c.Request.Context(), id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
