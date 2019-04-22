/**
 *  @author: yanKoo
 *  @Date: 2019/4/21 19:58
 *  @Description:
 */
package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)
var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	//request.Header.Add("Host", "album.zhenai.com")
	//request.Header.Add("Cookie", "xxxxxx")
	request.Header.Add("User-Agent", getUserAgent())

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}


func getUserAgent() string {

	headers := []string{ "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36",
	}
	//以当前的系统时间作为种子
	rand.Seed(time.Now().UnixNano())
	return headers[rand.Intn(5)]

}

// 猜测字节流的编码方式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
