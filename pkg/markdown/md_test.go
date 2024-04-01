package markdown

import (
	"github_spider/pkg/model"
	"testing"
)

func TestParse(t *testing.T) {

	issues := []model.Issue{
		{
			Title: "Test1",
			Link:  "http://123.com",
			Info:  []string{"1223", "32"},
		},
		{
			Title: "Test2",
			Link:  "http://123.com",
			Info:  []string{"1223", "32"},
		},
	}

	for _, issue := range issues {
		Parse(issue, "template/issue.tpl")
	}

}
