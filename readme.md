# Gofluence

A OPENAPI generated client to Confluence Rest API V2

The endpoints are documented here https://developer.atlassian.com/cloud/confluence/rest/v2/.


## Creation

There is a task file (https://taskfile.dev/) which will do the following steps:

- download the openapi from the Atlassian Confluence page `task get_api`
- edit the api to something that oapi codegen can generate from  `task edit_api`
- generate the go interface from the edited  OpenAPI specification `task gen`

Or just `task dia` to run all the steps in sequence.

## Code

There is /api which is the genereated interface from the specification

There is also /api/fat which is a fatter client to do a specific task.  
So *UpdatePage* will create or update a page given an ancestor page id, a title and a new body with a representation

This was the function I needed to generate documentation automatically.

## Examples
The examples have been tested by locally.

### first.go
This shows the simplest calls to Confluence to confirm it is working.  


## Change Log

# V0.1.0
Introduced fat client