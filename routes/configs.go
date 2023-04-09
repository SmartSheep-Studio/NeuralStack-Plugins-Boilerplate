package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"net/http"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/configs"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/services"
)

func reloadConfiguration(c *gin.Context) {
	if !slices.Contains(configs.GetConfig[[]string](plugin, "whitelist-ip"), c.ClientIP()) {
		if _, err := services.GetAuthorization(c, true, "is:super-admin"); err != nil {
			return
		}
	}

	if err := configs.SafeLoadConfig(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, viper.AllSettings())
	}
}
