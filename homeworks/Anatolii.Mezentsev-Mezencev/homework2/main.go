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
	URL  = "https://gist.githubusercontent.com/Mezencev/95381fceb309ef5bdaccfaafcbcc61d0/raw/11d3d4c6363b9ae631c225711e16a743590fb0b5/users.json"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// NameReader interface
type NameReader interface {
	Read(source string, key string) string
}

// file struct
type file struct{}

// url struct
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

func checkError(e error) {
	if e != nil {
		log.Fatal("Fatal error:", e)
		panic(e)
	}
}

func main() {
	var f, u NameReader = &file{}, &url{}
	fmt.Println("file=", f.Read(FILE, "firstName"))
	fmt.Println("url=", u.Read(URL, "firstName"))
}
