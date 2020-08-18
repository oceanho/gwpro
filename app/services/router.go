package services

import (
	"github.com/oceanho/gw"
	"github.com/oceanho/gwpro/app/services/restAPIs"
)

func RegisterServices(router *gw.RouterGroup) {
	router.RegisterRestAPIs(&restAPIs.UserRestAPI{})
}
