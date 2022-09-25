package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("my server start")
	nsm := http.NewServeMux()
	nsm.HandleFunc("/health", health)
	nsm.HandleFunc("/ok", ok)
	http.ListenAndServe(":8080", nsm)
}

func ok(writer http.ResponseWriter, request *http.Request) {

}

func health(writer http.ResponseWriter, request *http.Request) {
	// 接收客户端 request，并将 request 中带的 header 写入 response header
	header := request.Header
	headerName := header.Get("hname")
	//name := request.FormValue("userName")
	writer.Header().Set("recivedheadername", "recived username is "+headerName)
	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	remoteIp := request.RemoteAddr
	if ip := request.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = request.Header.Get("X-Forword-IP"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}

	fmt.Sprintln("来自 %s 的请求访问", remoteIp)

	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	writer.Header().Set("version", version)

	// 当访问 localhost/healthz 时，应返回 200
	//fmt.Fprintln(writer, "Header全部数据:", request.Header)
	writer.WriteHeader(200)
	io.WriteString(writer, "succeed：来自 "+remoteIp+"的请求")

}
