package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl user.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(user.AppName).(user.Service)
}
