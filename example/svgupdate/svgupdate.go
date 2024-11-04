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

func getSVG() string {
	new := `<?xml version='1.0' standalone='no'?>
<!DOCTYPE svg PUBLIC '-//W3C//DTD SVG 1.1//EN'
  'http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd'>
<svg width='100%' height='100%' xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink'>

   <title>SVG Table</title>

   <g id='columnGroup'>
      <rect x='65' y='10' width='75' height='110' fill='gainsboro'/>
      <rect x='265' y='10' width='75' height='110' fill='gainsboro'/>

      <text x='30' y='30' font-size='18px' font-weight='bold' fill='crimson'>
         <tspan x='30' dy='1.5em'>Q1</tspan>
         <tspan x='30' dy='1em'>Q2</tspan>
         <tspan x='30' dy='1em'>Q3</tspan>
         <tspan x='30' dy='1em'>Q4</tspan>
      </text>

      <text x='100' y='30' font-size='18px' text-anchor='middle'>
         <tspan x='100' font-weight='bold' fill='crimson'>Sales</tspan>
         <tspan x='100' dy='1.5em'>$ 223</tspan>
         <tspan x='100' dy='1em'>$ 183</tspan>
         <tspan x='100' dy='1em'>$ 277</tspan>
         <tspan x='100' dy='1em'>$ 402</tspan>
      </text>

      <text x='200' y='30' font-size='18px' text-anchor='middle'>
         <tspan x='200' font-weight='bold' fill='crimson'>Expenses</tspan>
         <tspan x='200' dy='1.5em'>$ 195</tspan>
         <tspan x='200' dy='1em'>$ 70</tspan>
         <tspan x='200' dy='1em'>$ 88</tspan>
         <tspan x='200' dy='1em'>$ 133</tspan>
      </text>

      <text x='300' y='30' font-size='18px' text-anchor='middle'>
         <tspan x='300' font-weight='bold' fill='crimson'>Net</tspan>
         <tspan x='300' dy='1.5em'>$ 28</tspan>
         <tspan x='300' dy='1em'>$ 113</tspan>
         <tspan x='300' dy='1em'>$ 189</tspan>
         <tspan x='300' dy='1em'>$ 269</tspan>
      </text>
   </g>
</svg>`
	return new
}

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
	// Add atachement

}
