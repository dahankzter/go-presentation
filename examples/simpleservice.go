package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Data struct {
	Id   string
	Name string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {

	if !strings.HasPrefix(req.URL.Path[1:], "aok") {
		http.Error(res, "I can't eat that!", http.StatusNotAcceptable)
		return
	}

	d := loadData(req.URL.Path[1:])

	data, err := json.Marshal(d)

	if err != nil {
		http.Error(res, fmt.Sprintf("Something went wrong: [%s]", err.Error()),
			http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-type", "application/json")

	fmt.Println(string(data))
	fmt.Fprintln(res, string(data))
}

func loadData(input string) *Data {
	d := &Data{
		Name: input,
		Id:   createUUID(),
	}
	return d
}

func createUUID() string {
	f, _ := os.Open("/dev/urandom")
	defer f.Close()

	b := make([]byte, 16)
	f.Read(b)

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
