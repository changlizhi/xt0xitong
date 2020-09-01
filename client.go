package main

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter1, _ := bodyWriter.CreateFormFile("files", "file1.txt")
	file1, _ := os.Open("file1.txt")
	defer file1.Close()
	io.Copy(fileWriter1, file1)

	fileWriter2, _ := bodyWriter.CreateFormFile("files", "123.mp4")
	file2, _ := os.Open("123.mp4")
	defer file2.Close()
	io.Copy(fileWriter2, file2)

	canShu := "[\"1\",\"2\",\"3\"]"
	bodyWriter.WriteField("arr", canShu)

	bodyWriter.Close()
	req, err := http.NewRequest("POST", "https://localhost:8888/hfxxt0xitong", bodyBuffer)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 忽略客户端安全tls
	}
	client := &http.Client{
		Transport: tr,
	}
	contentType := bodyWriter.FormDataContentType()
	req.Header.Set("Content-Type", contentType)

	// req.Body=ioutil.NopCloser(bodyBuffer)

	//发起请求

	resp, err := client.Do(req)
	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status, err, string(resp_body))
}
