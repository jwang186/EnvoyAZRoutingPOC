package main

import (
   "net/http"
   "net"
   "fmt"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.StatusOK)
   body := "Reached Server in AZ2" + "\n"
   w.Write([]byte(body))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.StatusOK)
   body := "OK" + "\n"
   w.Write([]byte(body))
}

func main() {
   listener, _ := net.Listen("tcp", ":8081")
   defer listener.Close()
   http.HandleFunc("/ok", httpHandler)
   http.HandleFunc("/ping", pingHandler)
   fmt.Println("Server on AZ1 started, supported URL path: /ok")
   http.Serve(listener, nil)
}