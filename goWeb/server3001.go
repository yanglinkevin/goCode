package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/buy/ticket", handleReq)
	http.ListenAndServe(":3001", nil)
}

//处理请求函数,根据请求将响应结果信息写入日志
func handleReq(w http.ResponseWriter, r *http.Request) {
	failedMsg :=  "handle in port:"
	time.Sleep(10000)
	writeLog(failedMsg, "./stat.log")
	fmt.Fprintf(w, "hello world")

}

//写入日志
func writeLog(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "3001")
	buf := []byte(content)
	fd.Write(buf)
}
