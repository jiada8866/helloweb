package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/parnurzeal/gorequest"
)

type caller struct {
	apis     map[string]uris
	counters map[string]counter
}

type uris []string
type counter chan int

var server = "http://localhost:1323"

func (c *caller) call() {
	fmt.Println(time.Now(), "start call")

	for method := range c.apis {
		go c.doMock(method)
	}
}

func (c *caller) doMock(method string) {
	uris := c.apis[method]

	for _, uri := range uris {
		counter := c.getCounter(method, uri)
		go stare(method, uri, counter)
	}
}

func (c *caller) setCounters() {
	if len(c.apis) < 0 {
		fmt.Println("empty apis")
		return
	}
	if c.counters == nil {
		c.counters = map[string]counter{}
	}
	for method, uris := range c.apis {
		for _, uri := range uris {
			key := generateKey(method, uri)
			ct := make(counter)
			c.counters[key] = ct
		}
	}
}

func (c *caller) getCounter(method, uri string) (ct counter) {
	key := generateKey(method, uri)
	ct, ok := c.counters[key]
	if !ok {
		ct = make(counter)
		c.counters[key] = ct
	}
	return ct
}

func generateKey(method, uri string) string {
	return method + "#" + uri
}

func stare(method, uri string, ct counter) {
	url := server + uri
	rand.Seed(time.Now().UnixNano())
	for {
		r := rand.Intn(10)
		for i := 0; i < r; i++ {
			go doReq(method, url)
		}
		// 发送请求次数
		ct <- r
		// sleep for a little while
		time.Sleep(time.Millisecond * 50)
	}
}

func doReq(method, url string) {
	req := gorequest.New()
	switch method {
	case "GET":
		get(url, req)
	case "POST":
		post(url, req)
	}
}

func get(url string, req *gorequest.SuperAgent) {
	//resp, body, err := req.Get(url).End()
	_, _, err := req.Get(url).End()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resp, body)
}

func post(url string, req *gorequest.SuperAgent) {
	p := struct {
		Images  []string `json:"images"`
		Content string   `json:"content"`
	}{
		Images:  []string{"111", "222"},
		Content: "test",
	}
	//resp, body, err := req.Post(url).
	_, _, err := req.Post(url).
		Send(p).
		End()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resp, body)
}

// align 对齐到一秒的刚开始
func align() {
	currentNano := time.Now().UnixNano()
	nextNano := time.Now().Add(time.Second).Unix() * int64(time.Second)
	diff := nextNano - currentNano
	time.Sleep(time.Duration(diff))
}

// doStatistics 统计各个url的qps
func doStatistics(key string, c counter, rotateChan *time.Ticker, exitChan chan bool) {
	count := 0
	for {
		select {
		case n := <-c:
			count += n
		case <-rotateChan.C:
			// 这里打印的时间应该减1秒,因为统计的是上一秒的qps
			fmt.Println(time.Now().Add(-time.Second), key, count)
			count = 0
			// TODO 统计一个总数
		case <-exitChan:
			return
		}
	}
}

func main() {
	// TODO 实现可配置，通过flag和config file
	apis := map[string]uris{
		"GET":  {"/", "/one", "/two"},
		"POST": {"/post"},
	}
	caller := caller{
		apis: apis,
	}

	caller.setCounters()
	align()
	caller.call()

	exitChan := make(chan bool)
	for key, c := range caller.counters {
		rotateChan := time.NewTicker(time.Second)
		go doStatistics(key, c, rotateChan, exitChan)
	}

	time.Sleep(time.Second * 120)
	exitChan <- true
	// TODO 等待各个goroutine结束
}
