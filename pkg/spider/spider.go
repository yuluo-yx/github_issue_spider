package spider

import (
	"fmt"
	"github_spider/pkg/constants"
	"github_spider/pkg/model"
	"log"

	"github.com/gocolly/colly/v2"
)

func IssueSpider() *[]model.Issue {

	cc := model.GetCC()
	config := model.GetConfig().Github

	cc.UserAgent = config.Agent

	return getIssues(cc, config.Start, config.End, config.Url)
}

func getIssues(c *colly.Collector, start, end int, url string) *[]model.Issue {

	var issues []model.Issue

	c.OnHTML("div#show_issue", func(e *colly.HTMLElement) {
		child := e.DOM.Find("bdi")
		issue := model.Issue{
			Title: child.Text(),
			// Link:  e.Request.URL.String(),
		}
		issues = append(issues, issue)

		e.ForEach(""+
			"div#discussion_bucket "+
			"div.edit-comment-hide "+
			"td.d-block.comment-body.markdown-body.js-comment-body p",
			func(i int, el *colly.HTMLElement) {
				text := el.Text
				issueIndex := len(issues) - 1
				issues[issueIndex].Info = append(issues[issueIndex].Info, text)
			})
	})

	for i := start; i <= end; i++ {
		url := fmt.Sprintf("%s/%s/%d", url, constants.Type, i)
		fmt.Printf("issue task started: url is %v \n", url)
		err := c.Visit(url)
		if err != nil {
			log.Printf("spider error: %v\n", err.Error())
		}
	}

	return &issues
}
