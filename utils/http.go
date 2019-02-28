package utils

import (
	"net/http"
	"fmt"
	"math/rand"
	"math"
)
var (
	rd = rand.New(rand.NewSource(math.MaxInt64))
)

func Get(url string,header map[string]string) *http.Response {
	client := &http.Client{}
	//生成要访问的url
	//提交请求
	req, err := http.NewRequest("GET", url, nil)
	//增加header选项
	req.Header.Add("Referer", header["Referer"])
	req.Header.Add("User-Agent", header["User-Agent"])
	req.Header.Add("X-Requested-With",randomIP())
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(req)
	defer response.Body.Close()
	return response
}

func randomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rd.Intn(255), rd.Intn(255), rd.Intn(255), rd.Intn(255))
}