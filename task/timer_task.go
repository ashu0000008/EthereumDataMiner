package task

import (
	"github.com/robfig/cron/v3"
	"strconv"
)

func StartTimerTask(timeS int, funcCallBack func()) {
	crontab := cron.New(cron.WithSeconds()) //精确到秒

	//定时任务
	spec := "0 0 " + strconv.Itoa(timeS) + " * * ?"

	// 添加定时任务,
	crontab.AddFunc(spec, funcCallBack)
	// 启动定时器
	crontab.Start()
}
