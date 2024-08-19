package main

import (
	"fmt"
	"os"

	gofluence "github.com/drummonds/gofluence/api"
)

func main() {
	// var (
	// 	mail  = os.Getenv("JSM_USER_EMAIL")
	// 	token = os.Getenv("JSM_TOKEN")
	// )
	host := fmt.Sprintf("https://%s", os.Getenv("JSM_DOMAIN"))

	nc, err := gofluence.NewClient(host)

	fmt.Printf("Result %+v and err %v", nc, err)
}
