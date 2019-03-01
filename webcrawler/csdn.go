package webcrawler

import (
	"github.com/axgle/mahonia"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"../utils"
)

type Crawler interface {
	IncreaseVisits(blog string, data []string)
	GetArticles(url string, num int, proxyUrl string) []string
}

func NewCsdn() Crawler {
	return &csdn{}
}

type csdn struct {
}

func (c *csdn) IncreaseVisits(blog string, data []string) {
	if data == nil {
		return
	}
	ip := utils.GetIp()
	for _, value := range data {
		utils.GetRep(value, ip)
	}
}

func (*csdn) GetArticles(url string, num int, proxyUrl string) []string {
	var i int
	var data []string
	if num < 0 {
		return nil
	}
	for i = 1; i <= num; i++ {
		newUrl := url + strconv.Itoa(i)
		res := utils.GetRep(newUrl, proxyUrl)
		//res,_:= http.Get(newUrl)
		//body, _ := ioutil.ReadAll(res.Body)
		//fmt.Println(string(body))
		if res == nil {
			proxyUrl = utils.GetIp()
			continue
		}
		dec := mahonia.NewDecoder("utf-8")
		doc := dec.NewReader(res.Body)
		result, _ := goquery.NewDocumentFromReader(doc)
		result.Find(".content").Each(func(i int, selection *goquery.Selection) {
			if i <= 0 {
				return
			}
			v, _ := selection.Find("a").Attr("href")
			data = append(data, v)
		})
	}
	return data
}
