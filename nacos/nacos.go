package nacos

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
)

type Config struct {
	Mysql struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"redis"`
}

var configData Config

// Nacos 初始化并监听 Nacos 配置
func Nacos(nameSpace, username, password, addr string, port uint64, dataId, group string) (*Config, error) {
	// 创建日志目录
	os.MkdirAll("./tmp/nacos/log", os.ModePerm)
	os.MkdirAll("./tmp/nacos/cache", os.ModePerm)

	// Nacos 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         nameSpace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            username,
		Password:            password,
	}

	// Nacos 服务配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      addr,
			ContextPath: "/nacos",
			Port:        port,
			Scheme:      "http",
		},
	}

	// 创建 Nacos 客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Nacos 客户端初始化失败: %w", err)
	}

	// 获取初次配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return nil, fmt.Errorf("获取 Nacos 配置失败: %w", err)
	}

	// 解析初次配置
	if err := json.Unmarshal([]byte(content), &configData); err != nil {
		return nil, fmt.Errorf("初次解析配置失败: %w", err)
	}

	fmt.Println("初次加载配置参数:", configData)

	// 监听配置变更
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置变更，重新加载配置...")

			var newConfig Config
			if err := json.Unmarshal([]byte(data), &newConfig); err != nil {
				fmt.Println("配置更新失败:", err)
				return
			}

			configData = newConfig
			fmt.Println("配置更新成功:", configData)
		},
	})
	if err != nil {
		fmt.Println("监听配置失败:", err)
		os.Exit(1)
	}

	return &configData, nil
}

// GetConfig 获取当前最新的配置
func GetConfig() *Config {
	return &configData
}
