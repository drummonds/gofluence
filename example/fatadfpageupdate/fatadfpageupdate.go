// Aim is to create a page or update it usigg atlas_doc_format
// Testing fat client for ease of use
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/drummonds/gofluence/api/adf"
	"github.com/drummonds/gofluence/api/fat"
)

func getOldBody() string {
	newBody := fmt.Sprintf(`h3. Hello new page
h1. From the FAT client
On live demo
Has been updated again , at %s
AUTO GENERATED So don't bother editing

|| Key || Meaning ||
|🖥️| host DNS entry works is available|
|❌~DNS~| No host DNS entry, may not exist|
|🟢| API endpoint working|
`, time.Now().Format(time.RFC3339))
	return newBody
}

func getBody() string {
	adfDoc := adf.NewDocNode()
	adfDoc.Add(adf.NewHeadingNode(3, "Hello page with Atlassian Document Format"))
	// adfDoc.Content = append(adfDoc.Content, adf.NewTextNode(fmt.Sprintf("Updated at %s", time.Now().Format(time.RFC3339))))
	// adfDoc.Content = append(adfDoc.Content, adf.NewTextNode(fmt.Sprintf("AUTO GENERATED So don't bother editing")))
	// adfDoc.Content = append(adfDoc.Content, adf.NewTextNode(fmt.Sprintf("AUTO GENERATED So don't bother editing")))
	pn := adf.NewParagraphNode()
	adfDoc.Add(pn)
	pn.Add(adf.NewTextNode(fmt.Sprintf("AUTO GENERATED at %s. So don't bother editing", time.Now().Format(time.RFC3339))))
	pn.Add(adf.NewHardBreakNode())
	pn.Add(adf.NewTextNode("Now a demo of a table:"))
	tn := adf.NewTableNode()
	adfDoc.Add(tn)
	tr := adf.NewTableRowNode()
	tn.Add(tr)
	tr.Add(adf.NewTableHeader("Key"))
	tr.Add(adf.NewTableHeader("Meaning"))
	tr = adf.NewTableRowNode()
	tn.Add(tr)
	tr.Add(adf.NewTableCell("🖥️"))
	tr.Add(adf.NewTableCell("host DNS entry works is available"))
	tr = adf.NewTableRowNode()
	tn.Add(tr)
	tr.Add(adf.NewTableCell("❌~DNS~"))
	tr.Add(adf.NewTableCell("No host DNS entry, may not exist"))
	b, err := json.Marshal(adfDoc)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	newBody := fmt.Sprint(string(b))
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
	fc.UpdatePage(ctx, ancestorStr, title, getBody(), "atlas_doc_format")

}
