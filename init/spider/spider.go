package spider

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github_spider/pkg/model"
	"log"
	"net/http"
	"net/url"
)

func InitColly() {

	collector := colly.NewCollector()
	cfg := model.GetConfig().Proxy

	proxyCfg := "http://" + cfg.Address + cfg.Port
	log.Printf("proxy address info: %v\n", proxyCfg)
	proxyURL, err := url.Parse(proxyCfg)
	if err != nil {
		fmt.Println("Failed to parse proxy URL:", err)
		return
	}

	collector.SetProxyFunc(func(r *http.Request) (*url.URL, error) {
		return proxyURL, nil
	})

	model.SetCC(collector)
}
