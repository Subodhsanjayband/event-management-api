package main

import (
	"github.com/Subodhsanjayband/event_manager/db"
	"github.com/Subodhsanjayband/event_manager/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegesterRoutes(server)
	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
