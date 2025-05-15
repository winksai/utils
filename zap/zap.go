package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LoggerConfig 日志配置结构体
type LoggerConfig struct {
	Level      string // 日志级别: debug, info, warn, error
	Path       string // 日志路径
	MaxSize    int    // 每个日志文件最大大小 (MB)
	MaxBackups int    // 保留日志文件个数
	MaxAge     int    // 保留天数
	Compress   bool   // 是否压缩
}

// InitLogger 初始化 Zap 日志
func InitLogger(config LoggerConfig) *zap.Logger {
	// 创建日志目录
	logDir := filepath.Dir(config.Path)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		_ = os.MkdirAll(logDir, os.ModePerm)
	}

	// Lumberjack 配置
	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.Path,
		MaxSize:    config.MaxSize, // MB
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   // days
		Compress:   config.Compress, // 是否压缩
	}

	// Zap 日志级别转换
	level := zapcore.InfoLevel
	if err := level.UnmarshalText([]byte(config.Level)); err != nil {
		level = zapcore.InfoLevel
	}

	// 编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 编码器：JSON 格式
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 构建核心
	core := zapcore.NewCore(encoder, zapcore.AddSync(lumberjackLogger), level)

	// 构建 Logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// 全局替换 Logger
	zap.ReplaceGlobals(logger)

	return logger
}
