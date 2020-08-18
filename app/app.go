package app

import (
	"github.com/oceanho/gw"
	"github.com/oceanho/gwpro/app/dbModel"
	"github.com/oceanho/gwpro/app/gwImpl"
	"github.com/oceanho/gwpro/app/services"
)

type App struct {
}

func New() gw.App {
	return App{}
}

func (a App) Name() string {
	return "gw-pro"
}

func (a App) Router() string {
	return "gw-pro"
}

func (a App) Register(router *gw.RouterGroup) {
	services.RegisterServices(router)
}

func (a App) Migrate(ctx gw.MigrationContext) {
	dbModel.AutoMigrate(ctx.Store)
}

func (a App) Use(option *gw.ServerOption) {
	option.AuthManager = func(initCtx gw.ServerInitializationContext) gw.IAuthManager {
		return gwImpl.DefaultAuthManager(initCtx)
	}
}
