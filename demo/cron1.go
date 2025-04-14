package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("* * * * * *", cron1)
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	} //每五分钟执行一次
	_, err = c.AddFunc("* * * * * *", cron2)
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	} //每天凌晨一点清理过期的预约
	c.Start()
	defer c.Stop()

	// 防止主程序退出
	select {}
}
func cron1() {
	// 获取当前时间
	currentTime := time.Now()
	fmt.Println("当前时间:", currentTime.Format("2006-01-02 15:04:05"))

	// 加30分钟（返回新时间，原时间不变）
	futureTime1 := currentTime.Add(30 * time.Minute)
	fmt.Println("30分钟后:", futureTime1.Format("2006-01-02 15:04:05"))

	// 加1小时（3600秒）
	futureTime2 := currentTime.Add(time.Hour) // 等同于 Add(3600 * time.Second)
	fmt.Println("1小时后:", futureTime2.Format("2006-01-02 15:04:05"))

	// 计算时间差
	duration := futureTime1.Sub(currentTime)
	fmt.Printf("时间差: %.0f分钟\n", duration.Minutes())
}
func cron2() {

}
