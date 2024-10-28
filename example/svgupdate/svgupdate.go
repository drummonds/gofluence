// Aim is to create a page or update it.  The page will
// an embeded SVG file which will be created on the fly
// Testing fat client for ease of use
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	fat "github.com/drummonds/gofluence/api/fat"
)

func getBody() string {
	newBody := fmt.Sprintf(`h3. Hello new page
h1. From the FAT client
Has been updated again , at %s
AUTO GENERATED So don't bother editing

|| Key || Meaning ||
|🖥️| host DNS entry works is available|
|❌~DNS~| No host DNS entry, may not exist|
|🟢| API endpoint working|
`, time.Now().Format(time.RFC3339))
	return newBody
}

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
	fc.UpdatePage(ctx, ancestorStr, title, getBody())

}
