package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/json-iterator/go"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		log.Println("main.go,YongHuHandler:err",err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		log.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())
		if part.FileName() == "" { // this is FormData
			bodyByte, err := ioutil.ReadAll(part)
			log.Println("main.go,35---", string(bodyByte)) //传文件的话字符没法单独传了
			if err != nil {
				w.Write([]byte("error,0:参数读取错误" + err.Error()))
				return
			}
			bodyStrArr := []string{}
			err = json.Unmarshal(bodyByte, &bodyStrArr)
			if err != nil {
				w.Write([]byte("error,1:参数格式错误"))
				return
			}
			if len(bodyStrArr) == 0 {
				w.Write([]byte("error,2:请勿传入空数组！"))
				return
			}
		} else { // This is FileData
			dst, _ := os.Create("./" + part.FileName() + ".upload")
			defer dst.Close()
			io.Copy(dst, part)
		}
	}

}

func YongHuHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
  w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
  w.Header().Set("content-type", "application/json")             //返回数据格式是json
	if r.Method == "OPTIONS"{
    w.Write([]byte("200"))
    return
  }
  json := jsoniter.ConfigCompatibleWithStandardLibrary
  datab, err := ioutil.ReadAll(r.Body)
  if err != nil{
    log.Println("datab读取错误",err,r)
  }
  
  dataObj := map[string]interface{}{}
  err = json.Unmarshal(datab,&dataObj)
  if err != nil{
    log.Println("datab解析错误",err,r.Method)
  }
  log.Println("dataObj---",dataObj)
  str,err := json.Marshal(map[string]interface{}{
    "Obj":true,
  })
  if err != nil{
    log.Println("返回数据组装json错误",err)
  }
  w.Write([]byte(str))
  return
}

func StartAPI() {
	http.HandleFunc("/hfxyonghu", YongHuHandler)
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		log.Fatal("服务端报错：", err.Error())
	}

}
func main() {
	StartAPI()
}
