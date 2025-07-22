package main

import (
	"bufio"
	"log"
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
			req.SetBasicAuth("admin", "admin")
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
func main() {
	client := &SSEClient{URL: "http://192.168.1.2:8080/api/sse", done: make(chan struct{})}
	go client.Connect()
	//time.Sleep(30 * time.Second) // 模拟运行
	//client.Close()               // 主动关闭
	select {}
}
