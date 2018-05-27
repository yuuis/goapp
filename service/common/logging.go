package common

import (
  "../../config"
  "log"
  "net/http"
  "runtime"
)

func getUserAgent(r *http.Request) (string, string){
  if r == nil {
    return "", ""
  }
  return r.RemoteAddr, r.URL.Path
}

func WriteLog(errorLevel int, message string, r *http.Request) {
  _, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
  if ok {
    remoteAddr, path := getUserAgent(r)
    log.Printf("[%v] %v %v %v %s:%d", getLogName(errorLevel), message, remoteAddr, path, sourceFileName, sourceFileLineNum)
  }
}

func WriteErrorLog(errorLevel int, message error, r *http.Request) {
  _, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
  if ok {
    remoteAddr, path := getUserAgent(r)
    log.Printf("[%v] %v %v %v %s:%d", getLogName(errorLevel), message, remoteAddr, path, sourceFileName, sourceFileLineNum)
  }
}

func getLogName(errorLevel int) string {
  switch errorLevel {
    case config.DEBUG:
    return config.DEBUGNAME
  case config.INFO:
    return config.INFONAME
  case config.WARN:
    return config.WARNNAME
  case config.ERROR:
    return config.ERRORNAME
  case config.FATAL:
    return config.FATALNAME
  default:
    return ""
  }
}