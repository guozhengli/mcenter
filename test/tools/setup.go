package tools

import (
	"encoding/json"
	"io"
	"os"

	"github.com/infraboard/mcenter/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/logger/zap"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps"
)

func DevelopmentSetup() {
	// 初始化日志实例
	zap.DevelopmentSetup()

	// 初始化配置, 提前配置好/etc/unit_test.env
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 初始化缓存
	ins := memory.NewCache(conf.C().Cache.Memory)
	cache.SetGlobal(ins)

	// 初始化全局app
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
}

func AccessToken() string {
	return os.Getenv("MCENTER_ACCESS_TOKEN")
}

func ReadJsonFile(filepath string, v any) error {
	fd, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fd.Close()

	payload, err := io.ReadAll(fd)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, v)
}
