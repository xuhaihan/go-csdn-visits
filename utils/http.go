package utils

import (
	"net/http"
	"fmt"
	"math/rand"
	"math"
	"time"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"strconv"
)
var (
	rd = rand.New(rand.NewSource(math.MaxInt64))
	proxyServer="http://47.107.89.95:8090/get"
	//proxyServer="http://172.17.19.7:8090/get"
)


func Get(url string,header map[string]string) *http.Response {
	client := &http.Client{}
	//生成要访问的url
	//提交请求
	req, err := http.NewRequest("GET", url, nil)
	//增加header选项
	req.Header.Add("Referer", header["Referer"])
	req.Header.Add("User-Agent", header["User-Agent"])
	req.Header.Add("X-Requested-With",RandomIP())
	req.Header.Add("Cookie",RandomIP())
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(req)
	defer response.Body.Close()
	return response
}

func RandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rd.Intn(255), rd.Intn(255), rd.Intn(255), rd.Intn(255))
}

func GetIp() string{
	proxyMap:=make(map[string]interface{})
	res,_:=http.Get(proxyServer)
	body,err:=ioutil.ReadAll(res.Body)
	if err !=nil{
		return ""
	}
	fmt.Println(string(body))
	json.Unmarshal([]byte(string(body)),&proxyMap)
	port:=proxyMap["port"].(float64)
	ip:=(proxyMap["ip"]).(string)
	proxyUrl:="http://"+ip+":"+strconv.Itoa(int(port))
	return proxyUrl
}
/**
* 返回response
*/
func GetRep(urls string,ip string) *http.Response {

	request, _ := http.NewRequest("GET", urls, nil)
	//随机返回User-Agent 信息
	request.Header.Set("User-Agent", GetAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	proxy, err := url.Parse(ip)
	//设置超时时间
	timeout := time.Duration(20* time.Second)
	client := &http.Client{}
	if ip != "localhost"{
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
			Timeout: timeout,
		}
	}
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200{
		//fmt.Println("请求遇到了错误",err.Error())
		return nil
	}
	return response
}

/**
* 随机返回一个User-Agent
*/
func GetAgent() string {
	agent  := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(agent)
	return agent[r.Intn(len)]
}



