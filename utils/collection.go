package utils

import "fmt"

type Collection map[string]interface{}

func (c Collection) Has(name string) bool {
	_, ok := c[name]

	return ok
}

func (c Collection) Get(name string) interface{} {
	return c[name]
}

func (c Collection) GetDefault(name string, _default interface{}) interface{} {
	if c.Has(name) {
		return c.Get(name)
	}

	return _default
}

/* string */
func (c Collection) GetString(name string) (string, error) {
	val, isValid := c.Get(name).(string)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to string.", name)
}

func (c Collection) GetStringDefault(name string, _default string) string {
	if val, err := c.GetString(name); err == nil {
		return val
	}

	return _default
}

/* int - int8 - int16 - int32 - int64 - uint - uint8 - uint16 - uint32 - uint64 */
func (c Collection) GetInt(name string) (int, error) {
	val, isValid := c.Get(name).(int)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int.", name)
}

func (c Collection) GetIntDefault(name string, _default int) int {
	if val, err := c.GetInt(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetInt8(name string) (int8, error) {
	val, isValid := c.Get(name).(int8)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int8.", name)
}

func (c Collection) GetInt8Default(name string, _default int8) int8 {
	if val, err := c.GetInt8(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetInt16(name string) (int16, error) {
	val, isValid := c.Get(name).(int16)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int16.", name)
}

func (c Collection) GetInt16Default(name string, _default int16) int16 {
	if val, err := c.GetInt16(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetInt32(name string) (int32, error) {
	val, isValid := c.Get(name).(int32)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int32.", name)
}

func (c Collection) GetInt32Default(name string, _default int32) int32 {
	if val, err := c.GetInt32(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetInt64(name string) (int64, error) {
	val, isValid := c.Get(name).(int64)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to int32.", name)
}

func (c Collection) GetInt64Default(name string, _default int64) int64 {
	if val, err := c.GetInt64(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetUint(name string) (uint, error) {
	val, isValid := c.Get(name).(uint)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to uint.", name)
}

func (c Collection) GetUintDefault(name string, _default uint) uint {
	if val, err := c.GetUint(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetUint8(name string) (uint8, error) {
	val, isValid := c.Get(name).(uint8)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to uint8.", name)
}

func (c Collection) GetUint8Default(name string, _default uint8) uint8 {
	if val, err := c.GetUint8(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetUint16(name string) (uint16, error) {
	val, isValid := c.Get(name).(uint16)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to uint16.", name)
}

func (c Collection) GetUint16Default(name string, _default uint16) uint16 {
	if val, err := c.GetUint16(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetUint32(name string) (uint32, error) {
	val, isValid := c.Get(name).(uint32)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to uint32.", name)
}

func (c Collection) GetUint32Default(name string, _default uint32) uint32 {
	if val, err := c.GetUint32(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetUint64(name string) (uint64, error) {
	val, isValid := c.Get(name).(uint64)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to uint32.", name)
}

func (c Collection) GetUint64Default(name string, _default uint64) uint64 {
	if val, err := c.GetUint64(name); err == nil {
		return val
	}

	return _default
}

/* float32 - float64 */
func (c Collection) GetFloat32(name string) (float32, error) {
	val, isValid := c.Get(name).(float32)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to float32.", name)
}

func (c Collection) GetFloat32Default(name string, _default float32) float32 {
	if val, err := c.GetFloat32(name); err == nil {
		return val
	}

	return _default
}

func (c Collection) GetFloat64(name string) (float64, error) {
	val, isValid := c.Get(name).(float64)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to float64.", name)
}

func (c Collection) GetFloat64Default(name string, _default float64) float64 {
	if val, err := c.GetFloat64(name); err == nil {
		return val
	}

	return _default
}

/* bool */
func (c Collection) GetBool(name string) (bool, error) {
	val, isValid := c.Get(name).(bool)

	if isValid {
		return val, nil
	}

	return val, fmt.Errorf("The value \"%s\" cannot be converted to bool.", name)
}

func (c Collection) GetBoolDefault(name string, _default bool) bool {
	if val, err := c.GetBool(name); err == nil {
		return val
	}

	return _default
}
