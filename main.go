package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func YongHuHandler(w http.ResponseWriter, r *http.Request) {
	// 	datab, _ := ioutil.ReadAll(r.Body)
	// log.Println(r.Form["arr"])
	// var defaultMaxMemory int64
	//  defaultMaxMemory= 32 << 20 // 32 MB
	//  r.ParseMultipartForm(defaultMaxMemory)
	// log.Println("multiPart---",r.MultipartForm)
	// log.Println("multiPart---",r.Form)
	reader, err := r.MultipartReader()
	if err != nil {
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
func StartAPI() {
	http.HandleFunc("/hfxxt0xitong", YongHuHandler)
	err := http.ListenAndServeTLS(":8888", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("服务端报错：", err.Error())
	}

}
func main() {
	StartAPI()
}
