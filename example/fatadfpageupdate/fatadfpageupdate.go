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
	width := 1200.0
	layout := adf.TableNodeAttrsLayoutFullWidth
	tn.Attrs = &adf.TableNodeAttrs{Layout: &layout, Width: &width}
	adfDoc.Add(tn)

	tr := adf.NewTableRowNode()
	th := adf.NewTableHeader("Key")
	th.Attrs.Colwidth = []float64{250.0}
	tr.Add(th)
	th = adf.NewTableHeader("Meaning")
	th.Attrs.Colwidth = []float64{950.0}
	tr.Add(th)
	tn.Add(tr)

	// If you don't set te column width on the next rows it follows the header row
	tr = adf.NewTableRowNode()
	tr.Add(adf.NewTableCell("🖥️"))
	tr.Add(adf.NewTableCell("host DNS entry works is available"))
	tn.Add(tr)

	tr = adf.NewTableRowNode()
	tr.Add(adf.NewTableCell("❌~DNS~"))
	tr.Add(adf.NewTableCell("No host DNS entry, may not exist"))
	tn.Add(tr)
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
