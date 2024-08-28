// Aim is to create a page, update and then delete it.
// Switching to using a client with responses
// Highlighting JSON200 unmarshalling of return values
// Not using the with body versions which use an unmarshalled body
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
