package main

import (
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/buy/ticket", handleReq1)
	http.ListenAndServe(":3002", nil)
}

//处理请求函数,根据请求将响应结果信息写入日志
func handleReq1(w http.ResponseWriter, r *http.Request) {
	failedMsg :=  "handle in port:"
	writeLog1(failedMsg, "./stat.log")
}

//写入日志
func writeLog1(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "3002")
	buf := []byte(content)
	fd.Write(buf)
}
