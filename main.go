package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"html/template"
	"io/ioutil"
	"net/http"
	upload_delivery "ocr/app/upload/delivery"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("fileUp.gtpl")
		t.Execute(w, nil)
	} else {
		// parse input
		r.ParseMultipartForm(10 << 20)
		// 2. retrieve file
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		if !strings.Contains(handler.Header.Get("Content-Type"), "image") {
			w.WriteHeader(400)
			w.Write([]byte("please upload image"))
			return
		}
		// write file on our server

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		_uuid := uuid.New()
		err = ioutil.WriteFile("./images/"+_uuid.String()+".jpg", fileBytes, 0644)
		if err != nil {
			fmt.Fprintf(w, "Uploaded file failed\n"+err.Error())
			return
		}
		// update database
		// send message
		// return result
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	}
}
func downloadFile(w http.ResponseWriter, r *http.Request) {
	//f, err := os.Open(r.)
	//if err !=nil {
	//	w.WriteHeader(404)
	//	return
	//}
	//var _ io.Reader = (*os.File)(nil)
}
func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":2020", nil)
}
func main() {
	router := gin.Default()
	uploadDelivery, _ := upload_delivery.NewUploadDelivery()
	router.POST("/upload", uploadDelivery.UploadHandle)
	router.Run(":2020")
}
