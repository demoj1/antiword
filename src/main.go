package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	log.Printf("request\nUser-agent: %v\nLength: %v\n", r.UserAgent(), r.ContentLength)
	log.Println("start convert")

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := []byte(r.PostFormValue("data"))

	tmpName := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))

	file, err := os.OpenFile(tmpName, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	n, err := file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	if n != len(data) {
		log.Println("Write bytes not equal request bytes")
		log.Fatalf("Request: %v | Write: %v", len(data), n)
	}

	antiword := exec.Command("/usr/bin/antiword", tmpName)
	f, err := os.OpenFile("OUT"+tmpName, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("OUT" + tmpName)

	antiword.Stdout = f

	err = antiword.Run()
	if err != nil {
		log.Fatal(err)
	}

	plain, err := ioutil.ReadFile("OUT" + tmpName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("data: %v\n", string(plain))
	log.Println("convert to plain text end.")

	w.Write(plain)
}

func main() {
	http.HandleFunc("/", root)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
