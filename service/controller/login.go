package controller

import (
  "net/http"
  "../dto"
  "../common"
  "../logic"
  "../../config"
)


func GetLoginViewHandler(w http.ResponseWriter, r *http.Request) {
  common.WriteLog(config.DEBUG, "login", r)

  page := new(Page)
  page.Title = "login"

  tmpl, err := common.ViewParse("./view/login/login.html")
  if err != nil {
    common.WriteErrorLog(config.DEBUG, err, nil)
  }

  err = tmpl.Execute(w, page)
  if err != nil {
    common.WriteErrorLog(config.DEBUG, err, nil)
  }
}


func PostLoginViewHandler(w http.ResponseWriter, r *http.Request) {
  common.WriteLog(config.DEBUG, "login", r)
  r.ParseForm()

  if r.Method == "POST" {
    // フォーム入力情報からUserForm型を作成

    user := dto.UserForm{r.Form["username"][0], r.Form["password"][0]}
    page := new(Page)

    // 認証処理
    ok, err := logic.Authenticate(user)
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
    if ok{
      // ログイン成功
      page.Title = "logined"
      page.Message = "login success! "
    } else {
      // ログイン失敗
      page.Title = "can't login"
      page.Message = "login faild!"
    } 

    tmpl, err := common.ViewParse("view/login/logined.html")
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
    err = tmpl.Execute(w, page)
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
  } else {
    page := new(Page)
    page.Title = "wrong access"

    tmpl, err := common.ViewParse("view/login/login.html")
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
    err = tmpl.Execute(w, page)
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
  }
}