package main

import (
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/buy/ticket", handleReq3)
	http.ListenAndServe(":3004", nil)
}

//处理请求函数,根据请求将响应结果信息写入日志
func handleReq3(w http.ResponseWriter, r *http.Request) {
	failedMsg :=  "handle in port:"
	writeLog3(failedMsg, "./stat.log")
}

//写入日志
func writeLog3(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "3004")
	buf := []byte(content)
	fd.Write(buf)
}
