package controller

import (
	"marketplace/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var projects [5]domain.Project
	projects[0], _ = domain.NewProject("Project 1", 100000000)
	projects[1], _ = domain.NewProject("Project 2", 400000000)
	projects[2], _ = domain.NewProject("Project 3", 700000000)

	c.JSON(http.StatusOK, gin.H{"data": projects})
}
