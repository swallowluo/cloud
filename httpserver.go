package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// flag.Set("v", "4")
	log.Println("Http Server starting....")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	statusCode := 200
	processRequest(w, r, statusCode)
	io.WriteString(w, "MyTestHttpServer")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	statusCode := 200
	processRequest(w, r, statusCode)
	io.WriteString(w, "OK")
}

func processRequest(w http.ResponseWriter, r *http.Request, statusCode int) {
	//将request header中的内容，写入response的header中
	for key := range r.Header {
		value := r.Header.Get(key)
		// log.Printf("key=%s;   value=%s\n", key, value)
		w.Header().Set(key, value)
	}
	//读取环境变量中的Version，并写入reponse头重
	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)
	//读取客户端请求的URI
	requestUri := r.RequestURI
	//读取客户端IP、HTTP 返回码，并输出到日志中
	clientIP := r.RemoteAddr
	if clientIP == "" {
		clientIP = r.Header.Get("Remote_addr")
	}
	log.Printf("ClientIP[%s], RequestURI[%s], Status[%d]\n", clientIP, requestUri, statusCode)
}
