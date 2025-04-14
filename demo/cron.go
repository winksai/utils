package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {

	// 启用秒级解析
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("* * * * * *", func() {
		fmt.Println("每秒执行一次，当前时间:", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	}
	c.Start()
	defer c.Stop()

	// 防止主程序退出
	select {}
}
