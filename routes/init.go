package routes

import (
	"github.com/gin-gonic/gin"
	api "repo.smartsheep.studio/smartsheep/neuralstack-api"
)

var plugin api.Plugin

func Init(p api.Plugin, router gin.IRouter) {
	plugin = p

	router.POST("/api/administration/configuration/reload", reloadConfiguration)
}
