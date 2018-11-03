package routing

import (
	"fmt"
	"regexp"
	"strings"
)

type Route struct {
	Methods  []string
	Path     string
	Pattern  *regexp.Regexp
	Segments []Segment
	Handler  RouteHandler
	Name     string
}

type Segment struct {
	Name       string
	Type       string
	Required   bool
	Expression string
}

type RouteHandler func(*Context)

type RouteParams map[string]interface{}

func (r RouteParams) Has(name string) bool {
	_, ok := r[name]

	return ok
}

func (r RouteParams) Get(name string, _default interface{}) interface{} {
	if r.Has(name) {
		return r[name]
	}

	return _default
}

// PHP's list()
func list(arr []string, dest ...*string) {
	for i := range dest {
		if len(arr) > i {
			*dest[i] = arr[i]
		}
	}
}

// resolvePath Resolve and compile the path of the route
func resolvePath(path string) (*regexp.Regexp, []Segment, error) {
	var segments []Segment
	regexpSegment := regexp.MustCompile("({[^}]+})")

	/*
		This pattern: {([a-zA-z_][a-zA-Z_0-9]+)(\??)(?(?=:):(.*)|())}
		Get all parts of params in the following way:
		[0] Full match
		[1] Name
		[2] Required
		[3] Format([0-9], string, int,...)

		PHP understand the conditions But Golang still not support conditions in Regex expressions
		I uploaded an example of this pattern here: https://regex101.com/r/F9PlVb/1
	*/
	matches := regexpSegment.FindAllString(path, -1)

	if len(matches) > 0 {
		fmt.Println(matches)
		for _, match := range matches {
			segment := Segment{}
			_tmp := strings.TrimLeft(match, "{")
			_tmp = strings.TrimRight(_tmp, "}")

			list(strings.Split(_tmp, ":"), &_tmp, &segment.Type)

			result := strings.Split(_tmp, "?")
			segment.Name = result[0]
			segment.Required = len(result) == 1
			segments = append(segments, segment)
			expression := resolveSegmentTypeExpression(&segment)
			path = strings.Replace(path, match, expression, -1)
			//fmt.Println("Pattern segments: " + match)
		}
	}

	path = "^/" + strings.Trim(path, "/") + "$"

	// fmt.Println("Pattern: " + path)

	pathCompiled, err := regexp.Compile(path)

	if err != nil {
		return nil, segments, err
	}

	return pathCompiled, segments, nil
}

func resolveSegmentTypeExpression(segment *Segment) string {
	segment.Expression = "("

	switch segment.Type {
	case "int":
		segment.Expression += "[0-9]"
	case "string":
		segment.Expression += "[a-zA-Z_]"
	case "alpha":
		segment.Expression += "[a-zA-z_0-9]"
	case "any":
	default:
		segment.Type = "any"
		segment.Expression += "[^\\/]"
	}

	if segment.Required {
		segment.Expression += "+)"
	} else {
		segment.Expression += "*)"
	}

	return segment.Expression
}

func NewRoute(methods []string, path string, handler RouteHandler) *Route {
	route := &Route{}
	route.Methods = methods
	route.Path = path
	pattern, segments, err := resolvePath(path)

	if err != nil {
		fmt.Println("Error compilando la ruta: " + path)
		panic(err)
	}

	route.Pattern = pattern
	route.Segments = segments

	route.Handler = handler

	return route
}

func (r *Route) Match(path string) (bool, RouteParams) {
	matches := r.Pattern.FindStringSubmatch(path)

	if matches != nil {
		params := make(RouteParams)

		for index, match := range matches[1:] {
			segment := r.Segments[index]
			params[segment.Name] = match // Falta convertir el valor al tipo definido en "segment.Type"
		}

		return true, params
	}

	return false, nil
}

func (r *Route) HasMethod(method string) bool {
	for _, _method := range r.Methods {
		if strings.ToUpper(_method) == strings.ToUpper(method) {
			return true
		}
	}

	return false
}

func (r *Route) SetName(name string) *Route {
	r.Name = name
	return r
}
