package main

import(
	"./webcrawler"
	"flag"
	"github.com/gpmgo/gopm/modules/goconfig"
	"strconv"
	"time"
	"./utils"
	"fmt"
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
	num,_:=strconv.Atoi(blogNum)
	num = *flag.Int("num",num,"num of blog")

	csdn:=webcrawler.NewCsdn()
	proxyUrl:=utils.GetIp()
	data:=csdn.GetArticles(articlesUrl,num,proxyUrl)
	//重试
	for i:=0;data==nil;i++{
		data=csdn.GetArticles(articlesUrl,num,proxyUrl)
	}
	fmt.Println("文章总数量:",len(data))
	fmt.Println("文章链接:",data)
	for{
		go csdn.IncreaseVisits(blogName,data)
		time.Sleep(time.Duration(30)*time.Second)
	}
}
