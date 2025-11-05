package main

import (
	"config"
	"fmt"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) log(level, format string, field ...string) {
	fmt.Printf("[LOG] [%s] %s\n", level, format)
}

func (l *Logger) Debug(format string, field ...string) {
	l.log("DEBUG", format)
}

func (l *Logger) Info(format string, field ...string) {
	l.log("INFO", format)
}

func (l *Logger) Warn(format string, field ...string) {
	l.log("WARN", format)
}

func (l *Logger) Error(format string, field ...string) {
	l.log("ERROR", format)
}

func main() {
	// 实例化日志
	log := NewLogger()

	// setting config hook
	// config.SetHook(config.Info, func(ctx config.HookContext) {
	// 	log.Info("正在初始化配置...")
	// }).SetHook(config.Debug, func(ctx config.HookContext) {
	// 	log.Debug(ctx.Message)
	// }).SetHook(config.Info, func(ctx config.HookContext) {
	// 	log.Info(ctx.Message)
	// }).SetHook(config.Warn, func(ctx config.HookContext) {
	// 	log.Warn(ctx.Message)
	// }).SetHook(config.Error, func(ctx config.HookContext) {
	// 	log.Error(ctx.Message)
	// })

	// 设置配置选项
	opts := config.NewOption()
	opts.Filename.Set("config")
	opts.FileType.Set("yaml")
	opts.Path.Set("./configs")
	opts.Env.Set("dev") // production | development
	opts.DebounceDur.Set(800 * config.OptionDateMillisecond)

	fmt.Println(opts)

	// 实例化配置管理器
	manager := config.Default()

	if err := manager.Init(func(ctx *config.Context) {
		log.Debug("配置文件更新了")
	}); err != nil {
		log.Error(fmt.Sprintf("初始化配置失败 Error: %s", err.Error()))
	}

	log.Info("初始化成功")

	select {}
}
