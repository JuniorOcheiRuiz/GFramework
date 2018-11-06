package routing

import (
	"fmt"
	"strconv"
)

type RouteParams map[string]string

func (r RouteParams) Has(name string) bool {
	_, ok := r[name]

	return ok
}

func (r RouteParams) Get(name string) string {
	return r[name]
}

func (r RouteParams) GetDefault(name string, _default string) string {
	if r.Has(name) {
		return r.Get(name)
	}

	return _default
}

/* int - int8 - int16 - int32 - int64 - uint - uint8 - uint16 - uint32 - uint64 */
func (r RouteParams) GetInt(name string) (int, error) {
	val, err := strconv.Atoi(r.Get(name))

	if err == nil {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int.", name)
}

func (r RouteParams) GetIntDefault(name string, _default int) int {
	if val, err := r.GetInt(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetInt8(name string) (int8, error) {
	val, err := strconv.ParseInt(r.Get(name), 10, 8)

	if err == nil {
		return int8(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to int8.", name)
}

func (r RouteParams) GetInt8Default(name string, _default int8) int8 {
	if val, err := r.GetInt8(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetInt16(name string) (int16, error) {
	val, err := strconv.ParseInt(r.Get(name), 10, 16)

	if err == nil {
		return int16(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to int16.", name)
}

func (r RouteParams) GetInt16Default(name string, _default int16) int16 {
	if val, err := r.GetInt16(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetInt32(name string) (int32, error) {
	val, err := strconv.ParseInt(r.Get(name), 10, 32)

	if err == nil {
		return int32(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to int32.", name)
}

func (r RouteParams) GetInt32Default(name string, _default int32) int32 {
	if val, err := r.GetInt32(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetInt64(name string) (int64, error) {
	val, err := strconv.ParseInt(r.Get(name), 10, 64)

	if err == nil {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int32.", name)
}

func (r RouteParams) GetInt64Default(name string, _default int64) int64 {
	if val, err := r.GetInt64(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetUint(name string) (uint, error) {
	val, err := strconv.ParseUint(r.Get(name), 10, 0)

	if err == nil {
		return uint(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint.", name)
}

func (r RouteParams) GetUintDefault(name string, _default uint) uint {
	if val, err := r.GetUint(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetUint8(name string) (uint8, error) {
	val, err := strconv.ParseUint(r.Get(name), 10, 8)

	if err == nil {
		return uint8(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint8.", name)
}

func (r RouteParams) GetUint8Default(name string, _default uint8) uint8 {
	if val, err := r.GetUint8(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetUint16(name string) (uint16, error) {
	val, err := strconv.ParseUint(r.Get(name), 10, 16)

	if err == nil {
		return uint16(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint16.", name)
}

func (r RouteParams) GetUint16Default(name string, _default uint16) uint16 {
	if val, err := r.GetUint16(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetUint32(name string) (uint32, error) {
	val, err := strconv.ParseUint(r.Get(name), 10, 32)

	if err == nil {
		return uint32(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint32.", name)
}

func (r RouteParams) GetUint32Default(name string, _default uint32) uint32 {
	if val, err := r.GetUint32(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetUint64(name string) (uint64, error) {
	val, err := strconv.ParseUint(r.Get(name), 10, 64)

	if err == nil {
		return val, nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint32.", name)
}

func (r RouteParams) GetUint64Default(name string, _default uint64) uint64 {
	if val, err := r.GetUint64(name); err == nil {
		return val
	}

	return _default
}

/* float32 - float64 */
func (r RouteParams) GetFloat32(name string) (float32, error) {
	val, err := strconv.ParseFloat(r.Get(name), 32)

	if err == nil {
		return float32(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to float32.", name)
}

func (r RouteParams) GetFloat32Default(name string, _default float32) float32 {
	if val, err := r.GetFloat32(name); err == nil {
		return val
	}

	return _default
}

func (r RouteParams) GetFloat64(name string) (float64, error) {
	val, err := strconv.ParseFloat(r.Get(name), 64)

	if err == nil {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to float64.", name)
}

func (r RouteParams) GetFloat64Default(name string, _default float64) float64 {
	if val, err := r.GetFloat64(name); err == nil {
		return val
	}

	return _default
}

/* bool */
func (r RouteParams) GetBool(name string) (bool, error) {
	val, err := strconv.ParseBool(r.Get(name))

	if err == nil {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to bool.", name)
}

func (r RouteParams) GetBoolDefault(name string, _default bool) bool {
	if val, err := r.GetBool(name); err == nil {
		return val
	}

	return _default
}
