package routing

type HttpMethod string

const (
	HTTP_GET     HttpMethod = "GET"
	HTTP_POST    HttpMethod = "POST"
	HTTP_PUT     HttpMethod = "PUT"
	HTTP_PATCH   HttpMethod = "PATCH"
	HTTP_DELETE  HttpMethod = "DELETE"
	HTTP_OPTIONS HttpMethod = "OPTIONS"
	HTTP_ANY     HttpMethod = "*"
)

type RouteCollection struct {
	Routes []*Route
}

func (r *RouteCollection) AddRoute(route *Route) {
	r.Routes = append(r.Routes, route)
}

func (r *RouteCollection) Match(methods []HttpMethod, path string, handler RouteHandler) *Route {
	route := NewRoute(methods, path, handler)
	r.AddRoute(route)
	return route
}

func (r *RouteCollection) Get(path string, handler RouteHandler) *Route {
	return r.Match([]HttpMethod{HTTP_GET}, path, handler)
}

func (r *RouteCollection) Post(path string, handler RouteHandler) *Route {
	return r.Match([]HttpMethod{HTTP_POST}, path, handler)
}

func (r *RouteCollection) Put(path string, handler RouteHandler) *Route {
	return r.Match([]HttpMethod{HTTP_PUT}, path, handler)
}

func (r *RouteCollection) Patch(path string, handler RouteHandler) *Route {
	return r.Match([]HttpMethod{HTTP_PATCH}, path, handler)
}

func (r *RouteCollection) Delete(path string, handler RouteHandler) *Route {
	return r.Match([]HttpMethod{HTTP_DELETE}, path, handler)
}

func (r *RouteCollection) Options(path string, handler RouteHandler) *Route {
	return r.Match([]HttpMethod{HTTP_OPTIONS}, path, handler)
}

func (r *RouteCollection) Any(path string, handler RouteHandler) *Route {
	return r.Match([]HttpMethod{HTTP_ANY}, path, handler)
}
