package controller

import (
  "net/http"
  "../common"
  "../../config"
)

type Page struct {
  Title string
  Message string
}

func HelloViewHandler(w http.ResponseWriter, r *http.Request) {
  common.WriteLog(config.INFO, "hello", r)

  page := new(Page)
  page.Title = "hello"
  page.Message = "hello, go"

  tmpl, err := common.ViewParse("./view/hello/hello.html")
  if err != nil {
    common.WriteErrorLog(config.DEBUG, err, nil)
  }

  err = tmpl.Execute(w, page)
  if err != nil {
    common.WriteErrorLog(config.DEBUG, err, nil)
  }
}