package main

import(
	"./webcrawler"
	"flag"
	"github.com/gpmgo/gopm/modules/goconfig"
	"strconv"
	"time"
)

func main(){
	confUrl:=flag.String("c","conf/config.ini","url of configure")
	cfg, err := goconfig.LoadConfigFile(*confUrl)
	if err != nil{
		panic("加载配置文件错误，请输入正确的配置文件地址")
	}
	blogName, err := cfg.GetValue("CSDN", "blogName")
	blogNum, err := cfg.GetValue("CSDN", "blogNum")
	articlesUrl, err := cfg.GetValue("CSDN", "articleUrl")
	times, err := cfg.GetValue("CSDN", "times")
	retryTimes,_:=strconv.Atoi(times)
	num,_:=strconv.Atoi(blogNum)
	num = *flag.Int("num",num,"num of blog")

	csdn:=webcrawler.NewCsdn()
	data:=csdn.GetArticles(articlesUrl,num)
	//重试
	for i:=0;data==nil&&i<retryTimes;i++{
		data=csdn.GetArticles(articlesUrl,num)
		if i==retryTimes&&data==nil{
			return
		}
	}
	for{
		go csdn.IncreaseVisits(blogName,data)
		time.Sleep(time.Duration(30)*time.Second)
	}
}
