package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/silvagpmiguel/bored/bored-backend/model"
	"github.com/silvagpmiguel/bored/bored-backend/service"
)

var activityService *service.Service

func redirectToFrontend(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/home")
}

func getActivity(c *gin.Context) {
	id := c.Query("id")
	actType := c.Query("type")

	result, created := activityService.GetActivity(id, actType)

	if created {
		c.JSON(http.StatusOK, result)
	} else {
		c.String(http.StatusInternalServerError, "No activity ", id)
	}
}

func fetchActivityFromBoredAPI(c *gin.Context) {
	result, err := service.GetBoredFromBoredAPI()
	if err == nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.String(http.StatusInternalServerError, "%v", err)
	}
}

func updateActivity(c *gin.Context) {
	var newActivity model.Activity
	id := c.Query("id")
	_type := c.Query("type")
	err := c.BindJSON(&newActivity)
	if err == nil {
		activityService.UpdateActivity(id, _type, newActivity)
		c.String(http.StatusOK, "Updated Activity : %v", id)
	} else {
		c.String(http.StatusInternalServerError, "couldn't save activity : %v", err)
	}
}

func saveActivity(c *gin.Context) {
	var activity model.Activity
	err := c.BindJSON(&activity)
	if err == nil {
		activityService.CreateActivity(&activity)
		c.String(http.StatusOK, "Created Activity : %v", activity.ID)
	} else {
		c.String(http.StatusInternalServerError, "couldn't save activity : %v", err)
	}
}

func Start() error {
	var err error
	activityService, err = service.New()

	if err != nil {
		return fmt.Errorf("failed to start controller : %v", err)
	}

	router := gin.Default()

	if _, err := os.Stat("../bored-frontend/dist"); err == nil {
		router.GET("/", redirectToFrontend)
		router.Static("/home", "../bored-frontend/dist")
	}

	apiRoutes := router.Group("/api")
	apiRoutes.GET("/bored/get", getActivity)
	apiRoutes.GET("/bored/new", fetchActivityFromBoredAPI)
	apiRoutes.POST("/bored/save", saveActivity)
	apiRoutes.POST("/bored/update", updateActivity)
	router.Run(":8080")
	return nil
}
