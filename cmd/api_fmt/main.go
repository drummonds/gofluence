package main

import (
	"fmt"
	"os"

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

// Doesn't create enum structures
// Workaround is to replace with strings
// panics on error
func FixEnum(g *gabs.Container, dotPath string) {
	err := g.DeleteP(dotPath + ".enum")
	if err != nil {
		panic(fmt.Errorf("Deleting dotPath %s.enum %v", dotPath, err))
	}
}

func main() {
	dat, err := os.ReadFile("openapi-v2.v3.json")
	if err != nil {
		panic(err)
	}
	jsonObj, _ := gabs.ParseJSON(dat)

	// Replace "default": "my, team", with  "default": "team",
	_, err = jsonObj.SetP("team", "paths./spaces/{id}/content/labels.get.parameters.1.schema.default")
	if err != nil {
		panic(err)
	}
	_, err = jsonObj.SetP("team", "paths./spaces/{id}/labels.get.parameters.1.schema.default")
	if err != nil {
		panic(err)
	}
	// Remove deprecated properties
	err = jsonObj.DeleteP("components.schemas.InlineCommentModel.properties.properties.properties.inline-marker-ref")
	if err != nil {
		panic(err)
	}
	err = jsonObj.DeleteP("components.schemas.InlineCommentModel.properties.properties.properties.inline-original-selection")
	if err != nil {
		panic(err)
	}
	err = jsonObj.DeleteP("components.schemas.InlineCommentProperties.properties.inline-marker-ref")
	if err != nil {
		panic(err)
	}
	err = jsonObj.DeleteP("components.schemas.InlineCommentProperties.properties.inline-original-selection")
	if err != nil {
		panic(err)
	}
	// Remove Polymorphism - golang and the openapi doesn't handle polymorphic end points for one or many using oneOF
	// SO in the insterface I am getting rid of the grouped functions.  In a later date they can be put back
	// as another client function
	RemoveOneOff(jsonObj, "components.requestBodies.BlogPostUpdateRequest.content.application/json.schema.properties.body")
	RemoveOneOff(jsonObj, "components.requestBodies.BlogPostCreateRequest.content.application/json.schema.properties.body")
	RemoveOneOff(jsonObj, "components.requestBodies.CustomContentCreateRequest.content.application/json.schema.properties.body")
	RemoveOneOff(jsonObj, "components.requestBodies.CustomContentUpdateRequest.content.application/json.schema.properties.body")
	RemoveOneOff(jsonObj, "components.requestBodies.PageCreateRequest.content.application/json.schema.properties.body")
	RemoveOneOff(jsonObj, "components.requestBodies.PageUpdateRequest.content.application/json.schema.properties.body")
	RemoveOneOff(jsonObj, "components.schemas.CreateFooterCommentModel.properties.body")
	RemoveOneOff(jsonObj, "components.schemas.CreateInlineCommentModel.properties.body")
	RemoveOneOff(jsonObj, "components.schemas.UpdateFooterCommentModel.properties.body")
	RemoveOneOff(jsonObj, "components.schemas.UpdateInlineCommentModel.properties.body")
	// Must use strings rather than ints
	RemoveAnyOff(jsonObj, "components.requestBodies.ContentIdToContentTypeRequest.content.application/json.schema.properties.contentIds.items")
	// Work around for I think bug issue https://github.com/oapi-codegen/oapi-codegen/issues/467
	// Simplest is to replace enum with string but can do enum later
	FixEnum(jsonObj, "components.requestBodies.BlogPostCreateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.BlogPostUpdateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.ContentClassificationLevelDeleteRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.ContentClassificationLevelUpdateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.CustomContentCreateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.CustomContentUpdateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.PageCreateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.PageUpdateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.LiveEditContentClassificationLevelResetRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.LiveEditContentClassificationLevelUpdateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.TaskUpdateRequest.content.application/json.schema.properties.status")
	FixEnum(jsonObj, "components.requestBodies.SpaceDefaultClassificationLevelUpdateRequest.content.application/json.schema.properties.status")

	jsonOutput := jsonObj.StringIndent("", "  ")
	// Becomes `{"outer":{"values":{"first":10,"second":11}},"outer2":"hello world"}`
	err = os.WriteFile("confluence.json", []byte(jsonOutput), 0644)
	if err != nil {
		panic(err)
	}
}
