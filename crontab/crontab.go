package crontab

import cron2 "github.com/robfig/cron"

func InitCrontab() {
	cron := cron2.New()

	// 添加定时任务

	// demo
	_ = cron.AddJob(newCrontabDemo().getSpec(), newCrontabDemo())

	cron.Start()
}
