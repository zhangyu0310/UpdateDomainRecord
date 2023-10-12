package main

import (
	"UpdateDomainRecord/config"
	"UpdateDomainRecord/ipinfo"
	"UpdateDomainRecord/sdk"
	"flag"
	"fmt"
	"github.com/zhangyu0310/zlogger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configPath = flag.String("config", "./config.json", "config file path")
)

func main() {
	flag.Parse()
	err := config.InitializeConfigFromFile(*configPath)
	if err != nil {
		fmt.Println("Initialize config failed, err:", err)
		os.Exit(1)
	}
	cfg := config.GetGlobalConfig()
	_ = zlogger.New(cfg.LogPath, "UpdateDomainRecord", true, zlogger.LogLevelWarn)

	// TODO: watch配置更新并自动加载最新配置

	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	doWhatYouWant := func() {
		ip, err := ipinfo.GetMyIP()
		if err != nil {
			zlogger.Error("Get my ip failed, err:", err)
			return
		}
		err = sdk.RunOnce(ip)
		if err != nil {
			zlogger.Error("Run once failed, err:", err)
		}
	}
	doWhatYouWant()

	stop := false
	for !stop {
		select {
		case <-time.NewTimer(time.Duration(cfg.Frequency) * time.Minute).C:
			doWhatYouWant()
		case <-c:
			stop = true
			break
		}
	}
	zlogger.Info("Shutdown Server ...")
}
