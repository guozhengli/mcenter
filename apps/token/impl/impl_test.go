package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl token.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(token.AppName).(token.Service)
}
