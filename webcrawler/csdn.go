package webcrawler

import (
	"github.com/axgle/mahonia"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"net/http"
	"fmt"
)
type Crawler interface {
	IncreaseVisits(blog string,data []string)
	GetArticles(url string,num int) []string
}


func NewCsdn() Crawler{
	return &csdn{}
}

type csdn struct {

}

func(c *csdn) IncreaseVisits(blog string,data []string){
	if data ==nil{
		return
	}
	/*header:=make(map[string]string)
	header["Referer"]="https://blog.csdn.net/"+blog
	header["User-Agent"]="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	*/
	for _,value:=range data{
        	_,err:=http.Get(value)
        	if err!=nil{
        		fmt.Println(err.Error())
			}
	}
}

func(* csdn) GetArticles(url string,num int) []string{
	var i int
	var data []string
	if num<0{
		return nil
	}
	for i=1;i<=num;i++{
		newUrl:=url+strconv.Itoa(i)
		res,_:=http.Get(newUrl)
		dec := mahonia.NewDecoder("utf-8")
		doc := dec.NewReader(res.Body)
		result, _ := goquery.NewDocumentFromReader(doc)
		result.Find(".content").Each(func(i int, selection *goquery.Selection) {
			if i<=0 {
				return
			}
			v,_:=selection.Find("a").Attr("href")
			data=append(data,v)
		})
	}
	return data
}


