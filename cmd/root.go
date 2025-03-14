package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/infraboard/mcenter/cmd/initial"
	"github.com/infraboard/mcenter/cmd/start"
	"github.com/infraboard/mcenter/conf"
	"github.com/infraboard/mcenter/tracer"
	"github.com/infraboard/mcenter/version"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/cache/redis"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mcenter",
	Short: "基于Etcd的微服务注册中心",
	Long:  "基于Etcd的微服务注册中心",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

func initail() {
	// 初始化全局变量
	err := loadGlobalConfig(confType)
	cobra.CheckErr(err)

	// 初始化全局日志配置
	err = loadGlobalLogger()
	cobra.CheckErr(err)

	// 加载缓存
	err = loadCache()
	cobra.CheckErr(err)

	// 初始化全局app
	err = app.InitAllApp()
	cobra.CheckErr(err)

	// 初始化Tracer
	err = tracer.InitTracer()
	cobra.CheckErr(err)
}

// config 为全局变量, 只需要load 即可全局可用户
func loadGlobalConfig(configType string) error {
	// 配置加载
	switch configType {
	case "file":
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
	case "env":
		err := conf.LoadConfigFromEnv()
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown config type")
	}

	return nil
}

// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	lc := conf.C().Log
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}
	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level
	switch lc.To {
	case conf.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

func loadCache() error {
	l := zap.L().Named("INIT")
	c := conf.C()
	// 设置全局缓存
	switch c.Cache.Type {
	case "memory", "":
		ins := memory.NewCache(c.Cache.Memory)
		cache.SetGlobal(ins)
		l.Info("use cache in local memory")
	case "redis":
		ins := redis.NewCache(c.Cache.Redis)
		cache.SetGlobal(ins)
		l.Info("use redis to cache")
	default:
		return fmt.Errorf("unknown cache type: %s", c.Cache.Type)
	}

	return nil
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// 初始化设置
	cobra.OnInitialize(initail)
	RootCmd.AddCommand(start.Cmd)
	RootCmd.AddCommand(initial.Cmd)
	err := RootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/config.toml", "the service config from file")
	RootCmd.PersistentFlags().StringVarP(&confETCD, "config-etcd", "e", "127.0.0.1:2379", "the service config from etcd")
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the mcenter version")
}
