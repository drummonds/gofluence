// Aim is to get the body of a page.
// Testing fat client for ease of use
package main

import (
	"context"
	"fmt"
	"os"

	fat "github.com/drummonds/gofluence/api/fat"
)

func main() {
	fc, err := fat.NewClient(
		os.Getenv("JSM_DOMAIN"),
		os.Getenv("JSM_USER_EMAIL"),
		os.Getenv("JSM_TOKEN"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result %+v and err %v\n", fc, err)

	ctx := context.Background()

	ancestorStr := os.Getenv("CONF_ANCESTOR_ID")
	title := "Test gofluence"
	result, err := fc.GetPageBody(ctx, ancestorStr, title)
	body := result.Body.AtlasDocFormat.Value
	if err != nil {
		fmt.Sprintf("Could'nt find title %s for ancestor %s", title, ancestorStr)
	}
	fmt.Printf("Body: %s", *body)

}