package routing

type RouteSegment struct {
	Name       string
	Type       string
	Required   bool
	Expression string
}

func NewRouteSegment(name string, _type string, required bool) RouteSegment {
	segment := RouteSegment{}
	segment.Name = name
	segment.Type = _type
	segment.Required = required

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

	return segment
}
