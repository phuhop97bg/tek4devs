package main

import (
	"log"
	"net/http"
	"tek4devs/modules/task/business"
	"tek4devs/modules/task/repository/inmem"
	"tek4devs/modules/task/transport/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// Setup full service dependencies
	apiService := rest.NewAPI(business.NewBusiness(inmem.NewInMemStorage()))

	v1 := engine.Group("v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", apiService.CreateTaskHandler())
			tasks.GET("", apiService.ListTaskHandler())
			tasks.GET("/:id", apiService.GetTaskHandler())
			tasks.PATCH("/:id", apiService.UpdateTaskHandler())
			tasks.DELETE("/:id", apiService.DeleteTaskHandler())
		}
		ping := v1.Group("/ping")
		{
			ping.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, "pong")
			})
		}
	}

	if err := engine.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
