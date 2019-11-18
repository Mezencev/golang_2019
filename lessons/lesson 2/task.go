package main

// NameReader -
// source can be local file or url of JSON file
// read this file, parse and return "key" field value
type NameReader interface {
	Read(source string, key string) string
}
