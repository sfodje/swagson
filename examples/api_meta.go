/* api:meta
Info:
    Title: Swagger Petstore
    Description: |
        This is a sample server Petstore server.
        You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.
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
        Url: "http://swagger.io"
*/
// This is a test
package testapi

func test() {

}
