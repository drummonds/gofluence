package fat

import (
	"context"
	"fmt"
	"strconv"

	gofluence "github.com/drummonds/gofluence/api"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

type FatClient struct {
	Client *gofluence.ClientWithResponses
}

type UpdatePageResponse struct {
	Id      string
	Version int32
}

// eg domain="my_company.atlassian.net/wiki/api/v2/"
func NewClient(domain, user, token string) (*FatClient, error) {
	fc := new(FatClient)
	basicAuth, err := securityprovider.NewSecurityProviderBasicAuth(user, token)
	if err != nil {
		panic(err)
	}

	fc.Client, err = gofluence.NewClientWithResponses(domain, gofluence.WithRequestEditorFn(basicAuth.Intercept))
	return fc, err
}

func (fc *FatClient) UpdatePage(ctx context.Context, ancestorId, title, body string) (*UpdatePageResponse, error) {
	resp := new(UpdatePageResponse)
	// Test the Ancestor page exists and get spaceID
	spaceId, versionId, exists := fc.PageExistsById(ctx, ancestorId)
	if versionId == nil || spaceId == nil {
		return resp, fmt.Errorf("versionId or spaceId are nil for ancestor %v", ancestorId)
	}
	ids, err := idToIds(*spaceId)
	if err != nil {
		return nil, err
	}
	id, versionId, exists := fc.PageExistsByTitle(ctx, &ids, title)
	if versionId == nil || id == nil {
		return resp, fmt.Errorf("versionId or spaceId are nil for ancestor %v", ancestorId)
	}
	resp.Id = *id
	resp.Version = *versionId
	if !exists {
		createPageParams := gofluence.CreatePageParams{}
		var wiki gofluence.PageBodyWriteRepresentation = "wiki"
		coreBody := gofluence.PageBodyWrite{Representation: &wiki, Value: &body}

		createBody := gofluence.CreatePageJSONRequestBody{SpaceId: *spaceId, Title: &title, Body: &coreBody, ParentId: &ancestorId}
		create_response, err := fc.Client.CreatePageWithResponse(ctx, &createPageParams, createBody)
		if err != nil {
			return resp, err
		}
		if create_response.HTTPResponse.StatusCode != 200 {
			return resp, err
		}
		id := create_response.JSON200.Id
		resp.Id = *id
		versionNumber := create_response.JSON200.Version.Number
		resp.Version = *versionNumber
		return resp, err
	}
	// If it doesn't exist then Update page to this ancestor
	var thisId int64
	thisId, err = strconv.ParseInt(resp.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	newVersionNumber := resp.Version + 1
	var wiki gofluence.PageBodyWriteRepresentation = "wiki"
	updateBody := gofluence.PageBodyWrite{
		Representation: &wiki,
		Value:          &body}
	updateBodyJSON := gofluence.UpdatePageJSONBody{}
	updateBodyJSON.Id = resp.Id
	updateBodyJSON.Status = "current"
	updateBodyJSON.Title = title
	updateBodyJSON.Body = updateBody
	updateBodyJSON.Version.Number = &newVersionNumber
	var updateBodyReq gofluence.UpdatePageJSONRequestBody = gofluence.UpdatePageJSONRequestBody(updateBodyJSON)

	update_response, err := fc.Client.UpdatePageWithResponse(ctx, thisId, updateBodyReq)
	if err != nil {
		return resp, err
	}
	if update_response.HTTPResponse.StatusCode != 200 {
		return resp, err
	}
	versionNumber := update_response.JSON200.Version.Number
	resp.Version = *versionNumber
	return resp, err

}

func idToIds(s string) ([]int64, error) {
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	var ids = make([]int64, 1, 1)
	ids[0] = id
	return ids, nil
}

// return version number and found
func (fc *FatClient) PageExistsById(ctx context.Context, Id string) (*string, *int32, bool) {
	ids, err := idToIds(Id)
	if err != nil {
		return nil, nil, false
	}
	pageParams := gofluence.GetPagesParams{Id: &ids}
	pageResponse, err := fc.Client.GetPagesWithResponse(ctx, &pageParams)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return nil, nil, false
	}
	json := (*pageResponse).JSON200
	if json == nil {
		return nil, nil, false
	}
	results := json.Results
	if results == nil {
		return nil, nil, false
	}
	if len(*results) == 0 {
		return nil, nil, false
	}
	version := (*results)[0].Version
	if version == nil {
		return nil, nil, false
	}
	versionNumber := version.Number
	spaceId := (*results)[0].SpaceId
	if spaceId == nil {
		return nil, versionNumber, false
	}
	return spaceId, versionNumber, true
}

func (fc *FatClient) PageExistsByTitle(ctx context.Context, spacesId *[]int64, title string) (*string, *int32, bool) {
	pageParams := gofluence.GetPagesParams{SpaceId: spacesId, Title: &title}
	pageResponse, err := fc.Client.GetPagesWithResponse(ctx, &pageParams)
	if err != nil {
		panic(err)
	}
	json := (*pageResponse).JSON200
	if json == nil {
		return nil, nil, false
	}
	results := json.Results
	if results == nil {
		return nil, nil, false
	}
	if len(*results) == 0 {
		return nil, nil, false
	}
	id := (*results)[0].Id
	version := (*results)[0].Version
	if version == nil {
		return id, nil, false
	}
	versionNumber := version.Number
	return id, versionNumber, true
}
