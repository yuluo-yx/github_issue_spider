package markdown

import (
	"fmt"
	"html/template"
	"os"

	"github_spider/pkg/model"
)

func Parse(issue model.Issue, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("Failed to load template:", err)
		return
	}

	file, err := os.OpenFile("issue.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open output file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, issue)
	if err != nil {
		fmt.Println("Failed to write to output file:", err)
		return
	}

}
