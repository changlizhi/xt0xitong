package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/json-iterator/go"
  "xt0xitong/ml9fenfas"
)
func YongHuHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
  w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
  w.Header().Set("content-type", "application/json")             //返回数据格式是json
	if r.Method == "OPTIONS"{
    w.Write([]byte("200"))
    log.Println("OPTIONS嗅探了")
    return
  }
  json := jsoniter.ConfigCompatibleWithStandardLibrary
  datab, err := ioutil.ReadAll(r.Body)
  if err != nil{
    log.Println("datab读取错误",err,r)
  }
  
  dataObj := map[string]interface{}{}
  //从具体操作来说，拿到数据直接丢给分发应该没多大问题，但是如果参数没传上来那就不对应该直接返回错误信息提示
  err = json.Unmarshal(datab,&dataObj)
  if err != nil{
    log.Println("datab解析错误",err)
    str,_ := json.Marshal(map[string]interface{}{
      "Ceng1":"参数错误，请注意约定好的参数格式！",
    })

    w.Write(str)
    return
  }
  ret := ml9fenfas.QuanJuFenFa(dataObj)
  log.Println("ret---",ret)
  retByte,err := json.Marshal(ret)
  if err != nil{
    log.Println("返回数据组装json错误，这个错误不应该发生",err)
    str,_ := json.Marshal(map[string]interface{}{
      "Ceng1":"内部错误，请稍后重试！",
    })
    w.Write(str)
    return
  }
  w.Write(retByte)
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

