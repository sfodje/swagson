/* api:route
"/pet/{petId}":
    Get:
        Description: Returns a single pet
        OperationId: getPetById
        Parameters:
            - Description: ID of pet to return
              Format: int64
              In: path
              Name: petId
              Required: true
              Type: integer
        Produces:
            - application/xml
            - application/json
        Responses:
            200:
                Description: successful operation
                Schema:
                    "$ref": "#/definitions/Pet"
            400:
                Description: Invalid ID supplied
            404:
                Description: Pet not found
        Security:
            - api_key: []
        Summary: Find pet by ID
        Tags:
            - pet
*/
package testapi

func test() {

}
