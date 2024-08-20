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

	pageParams := gofluence.GetPageByIdParams{}
	page, err := nc.GetPageById(ctx, 98566152, &pageParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pages %+v and err %v\n", page, err)

	// Instead of return HTTP requests is probably better to parse them
	ncr, err := gofluence.NewClientWithResponses(host, gofluence.WithRequestEditorFn(basicAuth.Intercept))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result client with response%+v and err %v\n", ncr, err)

	status := gofluence.GetSpacesParamsStatusCurrent
	var limit int32 = 99
	params := gofluence.GetSpacesParams{Status: &status, Limit: &limit}
	spacesResponse, err := ncr.GetSpacesWithResponse(ctx, &params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("List of spaces\n")
	spaces := spacesResponse.JSON200.Results
	for i, v := range *spaces {
		fmt.Printf("  %v %v %v\n", i, *v.Id, *v.Name)
	}

}
