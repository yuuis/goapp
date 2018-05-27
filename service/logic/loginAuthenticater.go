package logic

import (
  "errors"
  "log"
  "../model"
  "../dto"
)

func Authenticate(user dto.UserForm) (bool, error) {
  ok := model.IsExistUserName(user.Username)
  log.Println("ok : ", ok)
  if ok {
    pwd := model.GetUserPassword(user.Username)
    log.Println("pwd: ", pwd)
    if pwd == user.Password {
      // ログイン成功
      return true, nil
    } else {
      // パスワードが違った場合
      return false, errors.New("wrong password")
    }
  } else {
    // ユーザネームが違った場合
    return false, errors.New("this username is not exist")
  }
}