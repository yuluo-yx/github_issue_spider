package core

import (
	"github_spider/init/config"
	"github_spider/init/spider"
	"github_spider/init/task"
)

func Init() {

	config.InitConfig()

	spider.InitColly()

	task.SyncIssueSpiderTask()
}
