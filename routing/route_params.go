package routing

import (
	"fmt"
)

type RouteParams map[string]interface{}

func (r RouteParams) Has(name string) bool {
	_, ok := r[name]

	return ok
}

func (r RouteParams) Get(name string) interface{} {
	return r[name]
}

func (r RouteParams) GetDefault(name string, _default interface{}) interface{} {
	if r.Has(name) {
		return r[name]
	}

	return _default
}

func (r RouteParams) GetString(name string) (string, error) {
	if r.Has(name) {
		if val, isValid := r.Get(name).(string); isValid {
			return val, nil
		}

		return "", fmt.Errorf("RouteParams: The param \"%s\" cannot be converted to string.", name)
	}

	return "", fmt.Errorf("RouteParams: Not exists the param \"%s\".", name)
}

func (r RouteParams) GetStringDefault(name string, _default string) string {
	if val, err := r.GetString(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetInt(name string) (int, error) {
	if r.Has(name) {
		if val, isValid := r.Get(name).(int); isValid {
			return val, nil
		}

		return 0, fmt.Errorf("RouteParams: The param \"%s\" cannot be converted to int.", name)
	}

	return 0, fmt.Errorf("RouteParams: Not exists the param \"%s\".", name)
}

func (r RouteParams) GetIntDefault(name string, _default int) int {
	if val, err := r.GetInt(name); err == nil {
		return val
	}

	return _default
}
