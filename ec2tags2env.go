package main

import "fmt"
import "net/http"

func getInstanceId() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		panic(err)
	}
	return resp.Body()
}

func main() {
	fmt.Println(getInstanceId)
}
