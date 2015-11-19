package main

import (
	"encoding/json"
	"errors"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
	"github.com/sfodje/swagson/specs"
)

// array of markup nodes to parse for
var NODES = []specs.MarkupNode{specs.APIMETA, specs.APIROUTE, specs.APIMODEL}

// getGoFiles returns an array of all files in the given directory (and subdirectores) with a '.go' extension
func getGoFiles(dir string) (*[]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return err
	})

	return &files, err
}

// extractComments returns a pointer to an array of all go-style comments in the given file
// this function assumes the file string passed to it is an existing file
func extractComments(file string, pkg ...string) (*[]string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	if len(pkg) > 0 && len(pkg[0]) > 0 && strings.ToLower(pkg[0]) != strings.ToLower(f.Name.String()) {
		return &[]string{}, nil
	}
	var comments []string
	for _, cpkg := range f.Comments {
		for _, c := range cpkg.List {
			comments = append(comments, c.Text)
		}
	}
	return &comments, nil
}

// extractMarkup extracts comment text for each markup and populates the markup parameter with the data
// It checks the first line of each comment group for any matching markup parameters (api:meta, api:route)
// If found it links the markup name to the body of the comment in the markup map
func extractMarkup(markup map[specs.MarkupNode][]string, comments *[]string) {
	for _, c := range *comments {
		split_comment := strings.Split(c, "\n")
		for _, node := range NODES {
			n := node.String()
			f_line := strings.ToLower(split_comment[0])
			// swagson comments must have at least two lines
			if strings.Contains(f_line, n) && len(split_comment) > 2 {
				split_comment = split_comment[1 : len(split_comment)-1]
				comment_str := strings.Join(split_comment, "\n")
				markup[node] = append(markup[node], comment_str)
			}
		}
	}
}

// extractSwaggerDoc extracts swagger doc data from the markup parameter,
// populates a specs.SwagDoc struct and returns a pointer to it.
func extractSwaggerDoc(markup *map[specs.MarkupNode][]string) (*specs.SwagDoc, error) {
	var swagDoc = new(specs.SwagDoc)
	var err error
	for node, docs := range *markup {
		switch node {
		default:
		case specs.APIMETA:
			err = handleMeta(swagDoc, docs)
		case specs.APIROUTE:
			err = handleRoute(swagDoc, docs)
		case specs.APIMODEL:
			err = handleModel(swagDoc, docs)
		}
		if err != nil {
			return nil, err
		}
	}
	return swagDoc, err
}

// handleMeta extracts data from api:meta markup
// there should only be one api:meta markup tag per project
// if there is more than one, the first tag will be used
func handleMeta(swagDoc *specs.SwagDoc, docs []string) error {
	swagDoc.Swagger = "2.0"
	y := []byte(docs[0])
	j, err := yaml.YAMLToJSON(y)
	err = json.Unmarshal(j, swagDoc)
	return err
}

// handleRoute extracts data from the api:route markup
func handleRoute(swagDoc *specs.SwagDoc, docs []string) error {
	for _, doc := range docs {
		var paths = make(map[string]specs.SwagPath)
		y := []byte(doc)
		j, err := yaml.YAMLToJSON(y)
		err = json.Unmarshal(j, &paths)
		swagDoc.Paths = &paths
		if err != nil {
			return err
		}
		for k, v := range paths {
			paths[k] = v
		}
	}
	return nil
}

// handleModel extracts data from the api:model markup
func handleModel(swagDoc *specs.SwagDoc, docs []string) error {
	var def = make(map[string]specs.SwagSchema)
	for _, doc := range docs {
		var model map[string]specs.SwagSchema
		y := []byte(doc)
		// for some reason yaml.Unmarshal throws error: "panic: reflect: reflect.Value.Set using unaddressable value"
		// with structs that have nested pointers
		// as a workaround, convert yaml string to json string and unmarshal
		j, err := yaml.YAMLToJSON(y)
		err = json.Unmarshal(j, &model)
		if err != nil {
			return err
		}
		for k, v := range model {
			def[k] = v
		}
	}
	swagDoc.Definitions = &def
	return nil
}

// swagDocToJson converts a swagDoc struct to json and returns a pointer
func swagDocToJson(swagDoc *specs.SwagDoc) (*[]byte, error) {
	y, err := yaml.Marshal(swagDoc)
	if err != nil {
		return nil, err
	}

	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		return nil, err
	}

	return &j, err
}

// swagDocToYaml converts a swagDoc struct to yaml and returns a pointer
func swagDocToYaml(swagDoc *specs.SwagDoc) (*[]byte, error) {
	y, err := yaml.Marshal(swagDoc)
	if err != nil {
		return nil, err
	}

	return &y, err
}

func main() {
	usage := `Swagson.

Usage:
  swagson <projectdir> <outputdir> [--yaml] [--package=<package>]
  swagson -h | --help
  swagson --version

Options:
  -h --help     	 	Show usage.
  -v --version     	 	Show version.
  -y --yaml  		 	Output as yaml format.
  -p --package=<package>  	Package name of project to be parsed.`

	arguments, _ := docopt.Parse(usage, nil, true, "0.0.1", false)
	var dir = arguments["<projectdir>"].(string)
	var outputdir = arguments["<outputdir>"].(string)
	var yaml = arguments["--yaml"].(bool)
	var pkg = arguments["--package"].(string)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("Error: %s does not exist", dir)
	}
	if _, err := os.Stat(outputdir); os.IsNotExist(err) {
		log.Fatalf("Error: %s does not exist", outputdir)
	}
	dir, _ = filepath.Abs(dir)
	outputdir, _ = filepath.Abs(outputdir)

	files, err := getGoFiles(dir)
	if err != nil {
		log.Fatal(err)
	}
	var markup = make(map[specs.MarkupNode][]string)
	for _, f := range *files {
		comment_arr, err := extractComments(f, pkg)
		if err != nil {
			log.Fatal(err)
		}
		extractMarkup(markup, comment_arr)
		if err != nil {
			log.Fatal(err)
		}
	}

	swagDoc, err := extractSwaggerDoc(&markup)
	if err != nil {
		log.Fatal(err)
	}

	if swagDoc.Swagger == "" {
		log.Fatal(errors.New("Missing required property: 'swagger'"))
	} else if swagDoc.Info == nil {
		log.Fatal(errors.New("Missing required property: 'info'"))
	} else if swagDoc.Paths == nil {
		log.Fatal(errors.New("Missing required property: 'paths'"))
	}

	var output *[]byte
	var file string

	if yaml {
		output, err = swagDocToYaml(swagDoc)
		file = filepath.Join(outputdir, "swagger.yaml")
	} else {
		output, err = swagDocToJson(swagDoc)
		file = filepath.Join(outputdir, "swagger.json")
	}

	if err != nil {
		log.Fatal(err)
	}

	perm := os.FileMode(0777)
	ioutil.WriteFile(file, *output, perm)
}
