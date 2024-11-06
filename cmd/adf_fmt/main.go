package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/Jeffail/gabs/v2"
)

// Where a path contains a polymorphic selector return first
// of list
// panics on error
func RemoveOneOff(g *gabs.Container, dotPath string) {
	ref, err := g.ArrayElementP(0, dotPath+".oneOf")
	if err != nil {
		panic(fmt.Errorf("dotPath %s.oneOf %v", dotPath, err))
	}
	_, err = g.SetP(ref, dotPath)
	if err != nil {
		panic(fmt.Errorf("dotPath %s.oneOf %v", dotPath, err))
	}
}

func RemoveAnyOff(g *gabs.Container, dotPath string) {
	ref, err := g.ArrayElementP(0, dotPath+".anyOf")
	if err != nil {
		panic(fmt.Errorf("dotPath %s.anyOf %v", dotPath, err))
	}
	_, err = g.SetP(ref, dotPath)
	if err != nil {
		panic(fmt.Errorf("dotPath %s.anyOf %v", dotPath, err))
	}
}

func Remove(g *gabs.Container, dotPath, suffix string) {
	ref, err := g.ArrayElementP(0, dotPath+"."+suffix)
	if err != nil {
		panic(fmt.Errorf("dotPath %s.%s %v", dotPath, suffix, err))
	}
	_, err = g.SetP(ref, dotPath)
	if err != nil {
		panic(fmt.Errorf("dotPath %s.%s %v", dotPath, suffix, err))
	}
}

// Doesn't create enum structures
// Workaround is to replace with strings
// panics on error
func Delete(g *gabs.Container, dotPath string) {
	err := g.DeleteP(dotPath)
	if err != nil {
		panic(fmt.Errorf("Deleting dotPath %s %v", dotPath, err))
	}
}

func main() {
	dat, err := os.ReadFile("AtlasDocFormat.json")
	if err != nil {
		panic(err)
	}
	jsonObj, err := gabs.ParseJSON(dat)
	if err != nil {
		panic(err)
	}

	// Iterate all children
	// S is shorthand for Search
	keys := make([]string, 0)
	for key, _ := range jsonObj.S("definitions").ChildrenMap() {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if i == 10 || i == 46 || i == 49 || i == 58 || i == 76 {
			fmt.Printf("Delete(jsonObj, \"definitions.%s\")\n", k)
		}
	}

	// Lists are not working

	Delete(jsonObj, "definitions.bulletList_node")          // Fail 10
	Delete(jsonObj, "definitions.listItem_node")            // Fail 46
	Delete(jsonObj, "definitions.mediaSingle_caption_node") // Fail 49
	Delete(jsonObj, "definitions.orderedList_node")         // Fail 58
	Delete(jsonObj, "definitions.taskList_node")            // Fail 76

	// Delete(jsonObj, "definitions.textColor_mark")
	// Delete(jsonObj, "definitions.text_node")
	// Delete(jsonObj, "definitions.underline_mark")

	jsonOutput := jsonObj.StringIndent("", "  ")
	err = os.WriteFile("adf.json", []byte(jsonOutput), 0644)
	if err != nil {
		panic(err)
	}
}
