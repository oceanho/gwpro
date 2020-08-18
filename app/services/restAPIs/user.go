package restAPIs

import "github.com/oceanho/gw"

type UserRestAPI struct {
}

func (u UserRestAPI) Name() string {
	return "user"
}

func (u UserRestAPI) Get(ctx *gw.Context) {

}

func (u UserRestAPI) QueryList(ctx *gw.Context) {

}

func (u UserRestAPI) Post(ctx *gw.Context) {

}

func (u UserRestAPI) Delete(ctx *gw.Context) {

}
