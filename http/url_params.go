package http

import (
	"fmt"
	"strconv"
)

type UrlParams map[string]string

func (p UrlParams) Has(name string) bool {
	_, ok := p[name]

	return ok
}

func (p UrlParams) Get(name string) string {
	return p[name]
}

func (p UrlParams) GetDefault(name string, _default string) string {
	if p.Has(name) {
		return p.Get(name)
	}

	return _default
}

/* int - int8 - int16 - int32 - int64 - uint - uint8 - uint16 - uint32 - uint64 */
func (p UrlParams) GetInt(name string) (int, error) {
	val, err := strconv.Atoi(p.Get(name))

	if err == nil {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int.", name)
}

func (p UrlParams) GetIntDefault(name string, _default int) int {
	if val, err := p.GetInt(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetInt8(name string) (int8, error) {
	val, err := strconv.ParseInt(p.Get(name), 10, 8)

	if err == nil {
		return int8(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to int8.", name)
}

func (p UrlParams) GetInt8Default(name string, _default int8) int8 {
	if val, err := p.GetInt8(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetInt16(name string) (int16, error) {
	val, err := strconv.ParseInt(p.Get(name), 10, 16)

	if err == nil {
		return int16(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to int16.", name)
}

func (p UrlParams) GetInt16Default(name string, _default int16) int16 {
	if val, err := p.GetInt16(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetInt32(name string) (int32, error) {
	val, err := strconv.ParseInt(p.Get(name), 10, 32)

	if err == nil {
		return int32(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to int32.", name)
}

func (p UrlParams) GetInt32Default(name string, _default int32) int32 {
	if val, err := p.GetInt32(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetInt64(name string) (int64, error) {
	val, err := strconv.ParseInt(p.Get(name), 10, 64)

	if err == nil {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int32.", name)
}

func (p UrlParams) GetInt64Default(name string, _default int64) int64 {
	if val, err := p.GetInt64(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetUint(name string) (uint, error) {
	val, err := strconv.ParseUint(p.Get(name), 10, 0)

	if err == nil {
		return uint(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint.", name)
}

func (p UrlParams) GetUintDefault(name string, _default uint) uint {
	if val, err := p.GetUint(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetUint8(name string) (uint8, error) {
	val, err := strconv.ParseUint(p.Get(name), 10, 8)

	if err == nil {
		return uint8(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint8.", name)
}

func (p UrlParams) GetUint8Default(name string, _default uint8) uint8 {
	if val, err := p.GetUint8(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetUint16(name string) (uint16, error) {
	val, err := strconv.ParseUint(p.Get(name), 10, 16)

	if err == nil {
		return uint16(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint16.", name)
}

func (p UrlParams) GetUint16Default(name string, _default uint16) uint16 {
	if val, err := p.GetUint16(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetUint32(name string) (uint32, error) {
	val, err := strconv.ParseUint(p.Get(name), 10, 32)

	if err == nil {
		return uint32(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint32.", name)
}

func (p UrlParams) GetUint32Default(name string, _default uint32) uint32 {
	if val, err := p.GetUint32(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetUint64(name string) (uint64, error) {
	val, err := strconv.ParseUint(p.Get(name), 10, 64)

	if err == nil {
		return val, nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to uint32.", name)
}

func (p UrlParams) GetUint64Default(name string, _default uint64) uint64 {
	if val, err := p.GetUint64(name); err == nil {
		return val
	}

	return _default
}

/* float32 - float64 */
func (p UrlParams) GetFloat32(name string) (float32, error) {
	val, err := strconv.ParseFloat(p.Get(name), 32)

	if err == nil {
		return float32(val), nil
	}

	return 0, fmt.Errorf("The value \"%s\" cannot be converted to float32.", name)
}

func (p UrlParams) GetFloat32Default(name string, _default float32) float32 {
	if val, err := p.GetFloat32(name); err == nil {
		return val
	}

	return _default
}

func (p UrlParams) GetFloat64(name string) (float64, error) {
	val, err := strconv.ParseFloat(p.Get(name), 64)

	if err == nil {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to float64.", name)
}

func (p UrlParams) GetFloat64Default(name string, _default float64) float64 {
	if val, err := p.GetFloat64(name); err == nil {
		return val
	}

	return _default
}

/* bool */
func (p UrlParams) GetBool(name string) (bool, error) {
	val, err := strconv.ParseBool(p.Get(name))

	if err == nil {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to bool.", name)
}

func (p UrlParams) GetBoolDefault(name string, _default bool) bool {
	if val, err := p.GetBool(name); err == nil {
		return val
	}

	return _default
}
