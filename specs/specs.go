package specs

// MarkupNode is an enum of all the markup nodes to parse for
type MarkupNode int

const (
	APIMETA MarkupNode = iota
	APIROUTE
	APIMODEL
)

var markupNodes = [...]string{"api:meta", "api:route", "api:model"}

func (m MarkupNode) String() string {
	return markupNodes[m]
}

type Scheme int

const (
	HTTP Scheme = iota
	HTTPS
	WS
	WSS
)

var schemes = [...]string{"http", "https", "ws", "wss"}

func (s Scheme) String() string {
	return schemes[s]
}

type SecurityType int

const (
	BASIC SecurityType = iota
	APIKEY
	OAUTH2
)

var securityType = [...]string{"basic", "apiKey", "oauth2"}

func (s SecurityType) String() string {
	return securityType[s]
}

type CollectionFormat int

const (
	CSV CollectionFormat = iota
	SSV
	TSV
	PIPES
)

var collectionFormats = [...]string{"csv", "ssv", "tsv", "pipes"}

func (c CollectionFormat) String() string {
	return collectionFormats[c]
}

type ParamLocation int

const (
	QUERY ParamLocation = iota
	HEADER
	PATH
	FORMDATA
	BODY
)

var paramLocation = [...]string{"query", "header", "path", "formdata", "body"}

func (p ParamLocation) String() string {
	return paramLocation[p]
}

type SwagDoc struct {
	Swagger             string                   `json:"swagger"`
	Info                *SwagInfo                `json:"info"`
	Host                string                   `json:"host,omitempty"`
	BasePath            string                   `json:"basePath,omitempty"`
	Schemes             *[]string                `json:"schemes,omitempty"`
	Consumes            *[]string                `json:"consumes,omitempty"`
	Produces            *[]string                `json:"produces,omitempty"`
	Paths               *map[string]SwagPath     `json:"paths"`
	Definitions         *map[string]SwagSchema   `json:"definitions,omitempty"`
	Parameters          *map[string]SwagParam    `json:"parameters,omitempty"`
	Responses           *map[string]SwagResponse `json:"responses,omitempty"`
	SecurityDefinitions *map[string]SwagSecDef   `json:"securityDefinitions,omitempty"`
	Security            *map[string][]string     `json:"security,omitempty"`
	Tags                *[]SwagTag               `json:"tags,omitempty"`
	ExternalDocs        *SwagExtDoc              `json:"externalDocs,omitempty"`
}

type SwagInfo struct {
	Title          string       `json:"title"`
	Description    string       `json:"description,omitempty"`
	TermsOfService string       `json:"termsOfService,omitempty"`
	Contact        *SwagContact `json:"contact,omitempty"`
	License        *SwagLicense `json:"license,omitempty"`
	Version        string       `json:"version"`
}

type externalReference struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type SwagContact struct {
	externalReference
	Email string `json:"email, omitempty"`
}

type SwagLicense struct {
	externalReference
}

type SwagPath struct {
	Ref        string         `json:"$ref,omitempty"`
	Get        *SwagOperation `json:"get,omitempty"`
	Put        *SwagOperation `json:"put,omitempty"`
	Post       *SwagOperation `json:"post,omitempty"`
	Delete     *SwagOperation `json:"delete,omitempty"`
	Options    *SwagOperation `json:"options,omitempty"`
	Head       *SwagOperation `json:"head,omitempty"`
	Patch      *SwagOperation `json:"patch,omitempty"`
	Parameters *[]SwagParam   `json:"parameters,omitempty"`
}

type SwagOperation struct {
	Tags         *[]string                `json:"tags,omitempty"`
	Summary      string                   `json:"summary,omitempty"`
	Description  string                   `json:"description,omitempty"`
	ExternalDocs *SwagExtDoc              `json:"externalDocs,omitempty"`
	OperationId  string                   `json:"operationId,omitempty"`
	Consumes     *[]string                `json:"consumes,omitempty"`
	Produces     *[]string                `json:"produces,omitempty"`
	Parameters   *[]SwagParam             `json:"parameters,omitempty"`
	Responses    *map[string]SwagResponse `json:"responses"`
	Schemes      *[]string                `json:"schemes,omitempty"`
	Deprecated   *[]string                `json:"deprecated,omitempty"`
	Security     *[]map[string][]string   `json:"security,omitempty"`
}

type SwagParam struct {
	Name             string            `json:"name"`
	In               string            `json:"in"`
	Description      string            `json:"description,omitempty"`
	Required         bool              `json:"required,omitempty"`
	Schema           *SwagSchema       `json:"schema,omitempty"`
	Type             string            `json:"type,omitempty"`
	Format           string            `json:"format,omitempty"`
	AllowEmptyValue  bool              `json:"allowEmptyValue,omitempty"`
	Items            *SwagItems        `json:"items,omitempty"`
	CollectionFormat *CollectionFormat `json:"collectionFormat,omitempty"`
	Default          interface{}       `json:"default,omitempty"`
	Maximum          int               `json:"maximum,omitempty"`
	ExclusiveMaximum bool              `json:"exclusiveMaximum,omitempty"`
	Minimum          int               `json:"minimum,omitempty"`
	ExclusiveMinimum bool              `json:"exclusiveMinimum,omitempty"`
	MaxLength        int               `json:"maxLength,omitempty"`
	MinLength        int               `json:"minLength,omitempty"`
	Pattern          string            `json:"pattern,omitempty"`
	MaxItems         int               `json:"maxItems,omitempty"`
	MinItems         int               `json:"minItems,omitempty"`
	UniqueItems      bool              `json:"uniqueItems,omitempty"`
	Enum             *[]string         `json:"enum,omitempty"`
	MultipleOf       int               `json:"multipleOf,omitempty"`
}

type SwagResponse struct {
	Description string                  `json:"description"`
	Schema      *map[string]interface{} `json:"schema,omitempty"`
	Headers     *map[string]SwagItems   `json:"headers,omitempty"`
	Examples    *map[string]interface{} `json:"examples,omitempty"`
}

type SwagSchema struct {
	Ref              string                 `json:"$ref,omitempty"`
	Format           string                 `json:"format,omitempty"`
	Title            string                 `json:"title,omitempty"`
	Description      string                 `json:"description,omitempty"`
	Default          string                 `json:"default,omitempty"`
	MultipleOf       int                    `json:"multipleOf,omitempty"`
	Maximum          int                    `json:"maximum,omitempty"`
	ExclusiveMaximum int                    `json:"exclusiveMaximum,omitempty"`
	Minimum          int                    `json:"minimum,omitempty"`
	ExclusiveMinimum int                    `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                    `json:"maxlength,omitempty"`
	MinLength        int                    `json:"minlength,omitempty"`
	Pattern          string                 `json:"pattern,omitempty"`
	MaxItems         int                    `json:"maxItems,omitempty"`
	MinItems         int                    `json:"minItems,omitempty"`
	UniqueItems      bool                   `json:"uniqueItems,omitempty"`
	MaxProperties    int                    `json:"maxProperties,omitempty"`
	MinProperties    int                    `json:"minProperties,omitempty"`
	Required         *[]string              `json:"required,omitempty"`
	Enum             *[]string              `json:"enum,omitempty"`
	Type             string                 `json:"type,omitempty"`
	Properties       *map[string]SwagSchema `json:"properties,omitempty"`
	Xml              *SwagXml               `json:"xml,omitempty"`
	Example          string                 `json:"example,omitempty"`
	Items            *SwagItems             `json:"items,omitempty"` // required if type is Array
}

type SwagXml struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}

type SwagItems struct {
	Ref              string            `json:"$ref,omitempty"`
	Description      string            `json:"description,omitempty"`
	Type             string            `json:"type"`
	Format           string            `json:"format,omitempty"`
	Items            *SwagItems        `json:"items,omitempty"` // required if type is Array
	CollectionFormat *CollectionFormat `json:"collectionFormat,omitempty"`
	Default          interface{}       `json:"default,omitempty"`
	Maximum          int               `json:"maximum,omitempty"`
	ExclusiveMaximum bool              `json:"exclusiveMaximum,omitempty"`
	Minimum          int               `json:"minimum,omitempty"`
	ExclusiveMinimum bool              `json:"exclusiveMinimum,omitempty"`
	MaxLength        int               `json:"maxLength,omitempty"`
	MinLength        int               `json:"minLength,omitempty"`
	Pattern          string            `json:"pattern,omitempty"`
	MaxItems         int               `json:"maxItems,omitempty"`
	MinItems         int               `json:"minItems,omitempty"`
	UniqueItems      bool              `json:"uniqueItems,omitempty"`
	Enum             *[]string         `json:"enum,omitempty"`
	MultipleOf       int               `json:"multipleOf,omitempty"`
}

type SwagSecDef struct {
	Type             string             `json:"type,omitempty"`
	Description      string             `json:"description,omitempty"`
	Name             string             `json:"name,omitempty"`
	In               string             `json:"in,omitempty"`
	Flow             string             `json:"flow,omitempty"`
	AuthorizationUrl string             `json:"authorizationUrl,omitempty"`
	TokenUrl         string             `json:"tokenUrl,omitempty"`
	Scopes           *map[string]string `json:"scopes,omitempty"`
}

type SwagTag struct {
	Name         string      `json:"name"`
	Description  string      `json:"description,omitempty"`
	ExternalDocs *SwagExtDoc `json:"externalDocs,omitempty"`
}

type SwagExtDoc struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}
