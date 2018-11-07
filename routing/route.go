package routing

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/juniorocheiruiz/gframework/http"
	"github.com/juniorocheiruiz/gframework/utils"
)

type Route struct {
	Methods  []string
	Path     string
	Pattern  *regexp.Regexp
	Segments []RouteSegment
	Handler  RouteHandler
	Name     string
}

type RouteHandler func(*http.Request, *http.Response)

// resolvePath Resolve and compile the path of the route
func resolvePath(path string) (*regexp.Regexp, []RouteSegment, error) {
	var segments []RouteSegment
	regexpSegment := regexp.MustCompile("({[^}]+})")

	/*
		This pattern: {([a-zA-z_][a-zA-Z_0-9]+)(\??)(?(?=:):(.*)|())}
		Get all parts of params in the following way:
		[0] Full match
		[1] Name
		[2] Required("?"=optional || ""=required)
		[3] Format([0-9], string, int,...)

		PHP support the conditions in Regex expressions But Golang still not support it. I made the comment because in the future would can support it.
		I uploaded an example of this pattern here: https://regex101.com/r/F9PlVb/1
	*/
	matches := regexpSegment.FindAllString(path, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			var name, _type string
			var required bool
			_tmp := strings.TrimLeft(match, "{")
			_tmp = strings.TrimRight(_tmp, "}")

			utils.List(strings.Split(_tmp, ":"), &_tmp, &_type)

			// when we make split if the length is == 1 it means that not contains "?", so is required.
			result := strings.Split(_tmp, "?")
			name = result[0]
			required = len(result) == 1

			segment := NewRouteSegment(name, _type, required)
			segments = append(segments, segment)
			path = strings.Replace(path, match, segment.Expression, -1)
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

func (r *Route) Match(path string) (bool, http.UrlParams) {
	matches := r.Pattern.FindStringSubmatch(path)

	if matches != nil {
		params := make(http.UrlParams)

		for index, match := range matches[1:] {
			segment := r.Segments[index]
			params[segment.Name] = match
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
