# Gofluence

A OPENAPI generated client to Confluence Rest API V2

The endpoints are documented here https://developer.atlassian.com/cloud/confluence/rest/v2/.

In addition a Atlassian Document Format api is also contained.  It has been generated from the schema and then further hand edited.


## Creation

There is a task file (https://taskfile.dev/) which will do the following steps:

- download the openapi from the Atlassian Confluence page `task get_api`
- edit the api to something that oapi codegen can generate from  `task edit_api`
- generate the go interface from the edited  OpenAPI specification `task gen`

Or just `task dia` to run all the steps in sequence.

## Code

There is /api which is the generated interface from the specification

There is also /api/fat which is a fatter client to do a specific task.  
So *UpdatePage* will create or update a page given an ancestor page id, a title and a new body with a representation

This was the function I needed to generate documentation automatically.

## Examples
The examples have been tested by locally.

### first.go
This shows the simplest calls to Confluence to confirm it is working.  

## Atlassian Document Format (ADF, atlas_doc_format)

One snag with this is that I was using a nice simple Wiki format but internally if you create
a page with Wiki format it gets converted inside Confluence to ADF.  So you if you get a page
created with Wiki you will read back an ADF document.

The ADF format can be helpful.  For instance the Wiki format does not allow varying column widths in tables which can be done in ADF.

The source file is here https://unpkg.com/@atlaskit/adf-schema@44.4.0/dist/json-schema/v1/full.json.  Due to time constraints I haven't made
a full download and update cycle like I did with the api.

The JSON schema was downloaded and turned into a go file with https://github.com/omissis/go-jsonschema.

The adf.go file ahs been hand edited, eg some content should be a list of pointers rather than a list of values.  With values further edits do not mutate the set.  In addition there are some 
additional helper functions in new_adf.go.  These should make the code using this more readable eg see example fatadfpageupdate.go

## Change Log

# V0.3.2
Introduced ADF document format with example on how to produce variable width columns.

# V0.1.0
Introduced fat client