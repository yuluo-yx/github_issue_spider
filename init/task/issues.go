package task

import (
	"fmt"
	"github_spider/pkg/constants"
	"github_spider/pkg/markdown"
	"github_spider/pkg/model"
	"github_spider/pkg/spider"
)

func SyncIssueSpiderTask() {

	config := model.GetConfig().Github
	issueSpider := spider.IssueSpider()

	for _, issue := range *issueSpider {
		markdown.Parse(issue, constants.TmplPath)
	}

	fmt.Printf("%d 到 %d issue 爬取完成！检查 issue.md 文件. \n", config.Start, config.End)
}
