package main

import (
	"bufio"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-http/pkg/util"
	"github.com/xxl6097/go-sse/pkg/sse"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type SSEClient struct {
	URL  string
	done chan struct{}
}

func (c *SSEClient) Connect() {
	for {
		select {
		case <-c.done:
			return
		default:
			req, _ := http.NewRequest("GET", c.URL, nil)
			req.SetBasicAuth("admin", "het002402")
			req.Header.Set("Accept", "text/event-stream")
			client := &http.Client{Timeout: 0} // 无超时限制
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("连接失败: %v，5秒后重试", err)
				time.Sleep(5 * time.Second)
				continue
			}

			scanner := bufio.NewScanner(resp.Body)
			var eventType, data string
			for scanner.Scan() {
				line := scanner.Text()
				switch {
				case strings.HasPrefix(line, "data:"):
					data = strings.TrimSpace(line[5:])
				case strings.HasPrefix(line, "event:"):
					eventType = strings.TrimSpace(line[6:])
				case line == "" && data != "":
					log.Printf("[%s] %s", eventType, data)
					data, eventType = "", ""
				}
			}
			resp.Body.Close()
		}
	}
}

func (c *SSEClient) Close() {
}
func GetLocalMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网络接口失败：", err)
		return ""
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.HardwareAddr != nil {
			devMac := strings.ReplaceAll(iface.HardwareAddr.String(), ":", "")
			fmt.Println(iface.Name, ":", devMac)
			return devMac
		}
	}
	return ""
}

func main() {
	//client := &SSEClient{URL: "http://uuxia.cn:7001/api/sse", done: make(chan struct{})}
	//go client.Connect()
	//time.Sleep(30 * time.Second) // 模拟运行
	//client.Close()               // 主动关闭

	url := "http://uuxia.cn:7001/api/sse"
	sse.NewClient(url).
		BasicAuth("admin", "het002402").
		ListenFunc(func(s string) {
			glog.Debugf("SSE: %s", s)
		}).Header(func(header *http.Header) {
		header.Add("Sse-Event-IP-Address", util.GetHostIp())
		header.Add("Sse-Event-MAC-Address", GetLocalMac())
	}).Done()
	select {}
}
