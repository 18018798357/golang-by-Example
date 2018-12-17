package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main(){
	// 初始化colly
	c := colly.NewCollector(
		// 只采集规定的域名下的内容
		colly.AllowedDomains("zhihu.com","api.zhihu.com"),
		// 最大深度
		colly.MaxDepth(1),
	)

	// 认证
	err := c.Post("https://api.zhihu.com/sign_in",map[string]string{"username":"a","password":"admin"})
	if err != nil {
		log.Fatal("==========",err)
	}

	c.OnHTML("a[href]", func(e *colly.HTMLElement){
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		 // 访问网站
		c.Visit(e.Request.AbsoluteURL(link))
	})


	c.OnRequest(func (r *colly.Request){
		fmt.Println("Visiting",r.URL.String())
	})

	//开始抓取
	er := c.Visit("https://www.zhihu.com/")
	fmt.Println(er)
}

