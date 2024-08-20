package main

import (
	"context"
	"fmt"
	"os"

	gofluence "github.com/drummonds/gofluence/api"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

func main() {
	// var (
	// 	mail  = os.Getenv("JSM_USER_EMAIL")
	// 	token = os.Getenv("JSM_TOKEN")
	// )
	// eg  JSM_DOMAIN="my_company.atlassian.net/wiki/api/v2/"

	host := fmt.Sprintf("https://%s", os.Getenv("JSM_DOMAIN"))

	basicAuth, err := securityprovider.NewSecurityProviderBasicAuth(
		os.Getenv("JSM_USER_EMAIL"),
		os.Getenv("JSM_TOKEN"))
	if err != nil {
		panic(err)
	}

	nc, err := gofluence.NewClient(host, gofluence.WithRequestEditorFn(basicAuth.Intercept))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result %+v and err %v\n", nc, err)

	ctx := context.Background()

	blogStatus := []gofluence.GetBlogPostsParamsStatus{gofluence.GetBlogPostsParamsStatusCurrent}
	blogParams := gofluence.GetBlogPostsParams{Status: &blogStatus}
	blogs, err := nc.GetBlogPosts(ctx, &blogParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("List of blogs %+v and err %v\n", blogs, err)

	status := gofluence.GetSpacesParamsStatusCurrent
	params := gofluence.GetSpacesParams{Status: &status}
	spaces, err := nc.GetSpaces(ctx, &params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("List of spaces %+v and err %v\n", spaces, err)

	pageParams := gofluence.GetPageByIdParams{}
	page, err := nc.GetPageById(ctx, 98566152, &pageParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pages %+v and err %v\n", page, err)

}
