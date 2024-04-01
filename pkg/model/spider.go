package model

import (
	"github.com/gocolly/colly/v2"
)

var (
	cc *colly.Collector
)

func SetCC(collector *colly.Collector) {

	cc = collector
}

func GetCC() *colly.Collector {

	return cc
}
