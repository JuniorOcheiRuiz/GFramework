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
}

type Segment struct {
	Name     string
	Type     string
	Required bool
}

type RouteHandler func(*Context)

type RouteParams map[string]interface{}

// PHP's list()
func list(arr []string, dest ...*string) {
	for i := range dest {
		if len(arr) > i {
			*dest[i] = arr[i]
		}
	}
}

// Resuelve y compila la path de la ruta
func resolvePath(path string) (*regexp.Regexp, []Segment, error) {
	var segments []Segment
	regexpSegment := regexp.MustCompile("{([^}]+)}")

	//{([a-zA-z_][a-zA-Z_0-9]+)(\??):?([^}$]+)}

	matches := regexpSegment.FindStringSubmatch(path)

	if len(matches) > 1 {
		for _, match := range matches[1:] {
			segment := Segment{}

			list(strings.Split(match, ":"), &match, &segment.Type)

			result := strings.Split(match, "?")
			segment.Name = result[0]
			segment.Required = len(result) == 1
			segments = append(segments, segment)
			_type := resolveTypeToRegex(segment.Type)

			if segment.Required {
				_type = _type + "+"
			} else {
				_type = _type + "*"
			}

			path = strings.Replace(path, match, "("+_type+")", -1)
		}
	}

	pathCompiled, err := regexp.Compile(path)

	if err != nil {
		return nil, segments, err
	}

	return pathCompiled, segments, nil
}

func resolveTypeToRegex(_type string) string {
	var value string

	switch _type {
	case "int":
		value = "[0-9]"
	case "string":
		value = "[a-zA-z_]"
	case "alpha":
		value = "[a-zA-z_0-9]"
	case "any":
	default:
		value = "."
	}

	return value
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
