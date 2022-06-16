// package main

// import (
// 	"fmt"
// )

// //*****How to run
// // Open the terminal or cmd prompt, CD to helloworld.go dir, Run the below command
// // go run HelloWorld.go

// //alias print method
// var print = fmt.Println

// func main() {
// 	display("hello world!")
// }

// func display(text string){
// 	print(text)
// }
package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"os"

	"bufio"
	"bytes"
	"compress/gzip"

	"github.com/golang/snappy"
	"io/ioutil"
)

var err error
var client *mongo.Client
var collection *mongo.Collection
var ctx = context.Background()

type Persister struct {
	client       *mongo.Client
	db           *mongo.Database
	traceEnabled bool
	txSession    mongo.Session
}

type Migration struct {
	Id       int    `json:"id"`
	FileName string `json:"filename"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

func main() {
	fileName := "snapshot"
	var err error
	var data []byte
	p := Migration{Id: 1, Name: "test", Value: "some random string"}

	if data, err = json.Marshal(p); err != nil {
		fmt.Println(err)
	}

	err = OutputGzipped(data, fileName)
	if err != nil {
		fmt.Println(err)
	}
	reader := NewFileReader(fileName)
	if data, err = ioutil.ReadAll(reader); err != nil {
		return err
	}
	var job Migration
	json.Unmarshal(data, &job)
	fmt.Println(job.Name)
	fmt.Println(job.Id)
}

func OutputGzipped(b []byte, filename string) error {
	var zbuf bytes.Buffer
	var n int
	var err error
	gz := gzip.NewWriter(&zbuf)
	nw := 0
	for nw < len(b) {
		if n, err = gz.Write(b[nw:]); err != nil {
			return err
		}
		nw += n
	}
	gz.Close() // flushing the bytes to the buffer.
	return ioutil.WriteFile(filename, zbuf.Bytes(), 0644)
}

func NewFileReader(filename string) (*bufio.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return NewReader(file)
}

func NewReader(file *os.File) (*bufio.Reader, error) {
	var buf []byte
	var err error
	var reader *bufio.Reader

	reader = bufio.NewReader(file)
	if buf, err = reader.Peek(10); err != nil {
		return reader, err
	}
	file.Seek(0, 0)
	bs, err := hex.DecodeString("ff060000734e61507059")
	if string(bs) == string(buf) {
		reader = bufio.NewReader(snappy.NewReader(file))
	} else if buf[0] == 31 && buf[1] == 139 {
		var zreader *gzip.Reader
		if zreader, err = gzip.NewReader(file); err != nil {
			return reader, err
		}
		reader = bufio.NewReader(zreader)
	} else {
		reader = bufio.NewReader(file)
	}

	return reader, nil
}
