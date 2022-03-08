/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/inoth/inobot/components/config"
	"github.com/inoth/inobot/components/logger"
	"github.com/inoth/inobot/components/queue"
	"github.com/inoth/inobot/global"
	"github.com/inoth/inobot/src/consumer"
	"github.com/inoth/inobot/src/executor"
	"github.com/spf13/cobra"
)

var (
	replica int
	conf    string
	topic   string
	channel string
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "启动脚本调度",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start run script executor.")
		run()
	},
}

func init() {
	runCmd.Flags().IntVar(&replica, "replica", 1, "初始化队列监听、脚本调度个数")
	runCmd.Flags().StringVar(&conf, "conf", "config.yaml", "配置文件目录")
	runCmd.Flags().StringVar(&topic, "topic", "", "队列消息订阅key")
	runCmd.Flags().StringVar(&channel, "channel", "", "队列消息渠道，默认与[topic]相同")
	rootCmd.AddCommand(runCmd)
}

func run() {
	cfg := config.ViperConfig{}
	if conf != "" {
		cfg.Path = conf
	}
	if topic == "" {
		must(errors.New("必须订阅一个消息key"))
		return
	}
	if channel == "" {
		channel = topic
	}
	must(global.Register(
		&logger.LogrusConfig{},
		cfg.SetDefaultValue(map[string]interface{}{"replica": replica}),
		&executor.LuaPool{},
	).Init().Run(
		(&queue.NsqQueue{Host: config.Cfg.GetString("Nsq.Host")}).
			AddConsumer(
				&consumer.LuaScriptExec{Topic: topic, Channel: channel, Workers: &executor.LuaExecutor{}},
			)))
}

func must(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
