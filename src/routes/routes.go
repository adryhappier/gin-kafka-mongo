package routes

import (
	handle_job "github.com/adryhappier/gin-kafka-mongo/src/handlers/job"
	// handle_user "github.com/adryhappier/gin-kafka-mongo/src/handlers/user"
	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		// api.GET("/users", handle_user.GetAllUser)
		// api.POST("/users", handle_user.CreateUser)
		// api.GET("/users/:id", handle_user.GetUser)
		// api.PUT("/users/:id", handle_user.UpdateUser)
		// api.DELETE("/users/:id", handle_user.DeleteUser)

		// jobs
		api.GET("/jobs", handle_job.GetAllJobs)
		api.POST("/sync/jobs", handle_job.SyncKafka)
	}
	r.Run(":8000")
}

/*
	File ini men-define endpoint dari macam-macam routes
	yang di buat pada folder src/routes
*/
