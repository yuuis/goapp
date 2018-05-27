package main

import (
  "net/http"
  "log"
  "./service/controller"
)

func authenticate(fn http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println("before process") // 前処理
    fn(w, r)
    log.Println("after process") // 後処理
  }
}
func handlerOnlyPost(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    log.Println("POSTじゃない")
    return
  }
  log.Println("POST")
}

func main() {
  http.HandleFunc("/", authenticate(controller.HelloViewHandler))
  http.HandleFunc("/hello", authenticate(controller.HelloViewHandler))
  http.HandleFunc("/login", authenticate(controller.GetLoginViewHandler))
  http.HandleFunc("/loginpost", authenticate(controller.PostLoginViewHandler))
  http.ListenAndServe(":3000", nil)
}
