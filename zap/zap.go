package zap

import (
	"go.uber.org/zap"
	"os"
)

func InitZap() {
	os.MkdirAll("./log", 0777)
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		"./log/dev.log",
	}
	build, _ := config.Build()
	zap.ReplaceGlobals(build)
	return
}
