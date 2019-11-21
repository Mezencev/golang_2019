package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// const FILE for json
const (
	FILE = "./users.json"
	URL  = "https://dog.ceo/api/breeds/image/random"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// NameReader interface
type NameReader interface {
	Read(source string, key string) string
}

// file struct
type file struct{}

// file struct
type url struct{}

// DataJSON interface for data
type DataJSON map[string]interface{}

func (f *file) Read(source string, key string) string {
	plan, _ := ioutil.ReadFile(source)
	var data DataJSON
	err := json.Unmarshal(plan, &data)
	checkError(err)
	return searchValue(data, key)
}

func (u *url) Read(source string, key string) string {
	res, err := myClient.Get(source)
	checkError(err)
	var data DataJSON
	err = json.NewDecoder(res.Body).Decode(&data)
	checkError(err)
	return searchValue(data, key)
}

func searchValue(data map[string]interface{}, key string) string {
	for k, v := range data {
		if k == key {
			return fmt.Sprintf("Key: %s Value: { %v }", key, v.(string))
		}
	}
	return fmt.Sprintf("Key: not found")
}

func main() {
	var f, u NameReader = &file{}, &url{}
	fmt.Println("file=", f.Read(FILE, "firstName"))
	fmt.Println("url=", u.Read(URL, "status"))
}

func checkError(e error) {
	if e != nil {
		log.Fatal("Fatal error:", e)
		panic(e)
	}
}
