package main

import (
	"encoding/json"
	"testing"

	"github.com/sfodje/swagson/specs"
)

const (
	APIMETATEXT = `Info:
    Title: Swagger Petstore
    Description: |
        This is a sample server Petstore server.
        You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key ` + "`special-key`" + ` to test the authorization filters.
    TermsOfService: "http://swagger.io/terms/"
    Contact:
        Name: John Doe
        Url: "http://swagger.io"
        Email: apiteam@swagger.io
    License:
        Name: Apache 2.0
        Url: "http://www.apache.org/licenses/LICENSE-2.0.html"
    Version: 1.0.0
Host: petstore.swagger.io
BasePath: /v2
Consumes:
    - application/json
    - application/xml
Produces:
    - application/json
    - application/xml
Schemes:
    - http
SecurityDefinitions:
    api_key:
        In: header
        Name: api_key
        Type: apiKey
    petstore_auth:
        AuthorizationUrl: "http://petstore.swagger.io/api/oauth/dialog"
        Flow: implicit
        Type: oauth2
        Scopes:
            "read:pets": read your pets
            "write:pets": modify pets in your account
ExternalDocs:
    Description: Find out more about Swagger
    Url: "http://swagger.io"
Tags:
    - Name: pet
      Description: Everything about your Pets
      ExternalDocs:
        Description: Find out more
        Url: "http://swagger.io"`
)

func getSwagDoc() *specs.SwagDoc {
	var swagDoc specs.SwagDoc = specs.SwagDoc{}
	swagDoc.Swagger = "2.0"
	swagDoc.Info = &specs.SwagInfo{}
	swagDoc.Info.Title = "Swagger Petstore"
	swagDoc.Info.Description = `This is a sample server Petstore server.
You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key ` + "`special-key`" + ` to test the authorization filters.
`
	swagDoc.Info.TermsOfService = "http://swagger.io/terms/"
	swagDoc.Info.Contact = &specs.SwagContact{}
	swagDoc.Info.Contact.Name = "John Doe"
	swagDoc.Info.Contact.Url = "http://swagger.io"
	swagDoc.Info.Contact.Email = "apiteam@swagger.io"
	swagDoc.Info.License = &specs.SwagLicense{}
	swagDoc.Info.License.Name = "Apache 2.0"
	swagDoc.Info.License.Url = "http://www.apache.org/licenses/LICENSE-2.0.html"
	swagDoc.Info.Version = "1.0.0"
	swagDoc.Host = "petstore.swagger.io"
	swagDoc.BasePath = "/v2"
	swagDoc.Consumes = &[]string{"application/json", "application/xml"}
	swagDoc.Produces = &[]string{"application/json", "application/xml"}
	swagDoc.Schemes = &[]string{"http"}
	swagDoc.SecurityDefinitions = &map[string]specs.SwagSecDef{}
	(*swagDoc.SecurityDefinitions)["api_key"] = specs.SwagSecDef{
		In:   "header",
		Name: "api_key",
		Type: "apiKey",
	}
	(*swagDoc.SecurityDefinitions)["petstore_auth"] = specs.SwagSecDef{
		AuthorizationUrl: "http://petstore.swagger.io/api/oauth/dialog",
		Flow:             "implicit",
		Type:             "oauth2",
		Scopes: &map[string]string{
			"read:pets":  "read your pets",
			"write:pets": "modify pets in your account",
		},
	}
	swagDoc.ExternalDocs = &specs.SwagExtDoc{
		Description: "Find out more about Swagger",
		Url:         "http://swagger.io",
	}
	swagDoc.Tags = &[]specs.SwagTag{}
	swagDoc.Tags = &[]specs.SwagTag{
		specs.SwagTag{
			Name:        "pet",
			Description: "Everything about your Pets",
			ExternalDocs: &specs.SwagExtDoc{
				Description: "Find out more",
				Url:         "http://swagger.io",
			},
		},
	}
	return &swagDoc
}

func Test_getGoFiles(t *testing.T) {
	if files, err := getGoFiles("./examples/"); len(*files) == 3 && err == nil {
		t.Log("getGoFiles(\"./examples/\") passed.")
	} else {
		t.Log(err)
		t.Error("getGoFiles(\"./examples/\") failed.")
	}

	if files, err := getGoFiles("./non_existent_folder/"); len(*files) == 0 && err != nil {
		t.Log("getGoFiles(\"./non_existent_folder/\") passed.")
	} else {
		t.Log(err)
		t.Error("getGoFiles(\"./non_existent_folder/\") failed.")
	}
}

func Test_extractComments(t *testing.T) {
	if comments, err := extractComments("./examples/api_meta.go"); len(*comments) == 2 && err == nil && (*comments)[1] == "// This is a test" {
		t.Log("extractComments(\"./examples/api_meta.go\") passed.")
	} else {
		t.Log(err)
		t.Error("extractComments(\"./examples/api_meta.go\") failed.")
	}
}

func Test_extractMarkup(t *testing.T) {
	var markup map[specs.MarkupNode][]string = make(map[specs.MarkupNode][]string)
	comments, err := extractComments("./examples/api_meta.go")
	extractMarkup(markup, comments)
	if err != nil || APIMETATEXT != markup[specs.APIMETA][0] {
		t.Log(err)
		t.Error("extractMarkup(markup, comments) failed.")
	} else {
		t.Log("extractMarkup(markup, comments) passed.")
	}
}

func Test_extractSwaggerDoc(t *testing.T) {
	var markup map[specs.MarkupNode][]string = make(map[specs.MarkupNode][]string)
	comments, err := extractComments("./examples/api_meta.go")
	extractMarkup(markup, comments)
	swagDoc, err := extractSwaggerDoc(&markup)
	j1, _ := json.Marshal(swagDoc)
	j2, _ := json.Marshal(getSwagDoc())
	if err != nil || string(j1) != string(j2) {
		t.Log(err)
		t.Error("extractSwaggerDoc(markup) failed.")
	} else {
		t.Log("extractSwaggerDoc(markup) passed.")
	}

}
