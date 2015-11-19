/* api:model
Pet:
    Xml:
        Name: Pet
    Type: object
    Required:
        - name
        - photoUrls
    Properties:
        category:
            $ref: "#/definitions/Category"
        id:
            Type: integer
            Format: int64
        name:
            Type: string
            Example: doggie
        photoUrls:
            Type: array
            Xml:
                Name: photoUrl
                Wrapped: true
            Items:
                Type: string
        status:
            Description: pet status in the store
            Enum:
                - available
                - pending
                - sold
            Type:
                string
        tags:
            Items:
                $ref: "#/definitions/Tag"
            Type: array
            Xml:
                Name: tag
                Wrapped: true
*/
/* api:model
Category:
    Type: object
    Xml:
        Name: Category
    Properties:
        id:
            Type: integer
            Format: int64
        name:
            Type: string
*/
/* api:model
Tag:
    Type: object
    Xml:
        Name: Tag
    Properties:
        id:
            Type: integer
            Format: int64
        name:
            Type: string
*/
package testapi

func test() {

}
