package core

import (
	"errors"
)

var (
	ErrorNoKeyFound = errors.New("no key found")
	ErrorWrongType  = errors.New("get wrong type")
)

type TtfInfo map[string]interface{}

func (me TtfInfo) PushString(key string, val string) {
	me[key] = val
}

func (me TtfInfo) PushBytes(key string, val []byte) {
	me[key] = val
}

func (me TtfInfo) PushInt64(key string, val int64) {
	me[key] = val
}

func (me TtfInfo) PushInt(key string, val int) {
	me[key] = val
}

func (me TtfInfo) PushUInt64(key string, val uint) {
	me[key] = val
}

func (me TtfInfo) PushBool(key string, val bool) {
	me[key] = val
}

func (me TtfInfo) PushInt64s(key string, val []int) {
	me[key] = val
}

func (me TtfInfo) PushMapIntInt64(key string, val map[int]int) {
	me[key] = val
}

func (me TtfInfo) GetBool(key string) (bool, error) {
	if val, ok := me[key]; ok {

		if m, ok := val.(bool); ok {
			/* act on str */
			return m, nil
		} else {
			return false, ErrorWrongType
		}
	} else {
		return false, ErrorNoKeyFound
	}
}

func (me TtfInfo) GetString(key string) (string, error) {
	if val, ok := me[key]; ok {

		if m, ok := val.(string); ok {
			/* act on str */
			return m, nil
		} else {
			return "", ErrorWrongType
		}
	} else {
		return "", ErrorNoKeyFound
	}
}

func (me TtfInfo) GetInt64(key string) (int, error) {
	if val, ok := me[key]; ok {

		if m, ok := val.(int); ok {
			/* act on str */
			return m, nil
		} else {
			return 0, ErrorWrongType
		}
	} else {
		return 0, ErrorNoKeyFound
	}
}

func (me TtfInfo) GetInt64s(key string) ([]int, error) {
	if val, ok := me[key]; ok {

		if m, ok := val.([]int); ok {
			/* act on str */
			return m, nil
		} else {
			return nil, ErrorWrongType
		}
	} else {
		return nil, ErrorNoKeyFound
	}
}

func (me TtfInfo) GetMapIntInt64(key string) (map[int]int, error) {
	if val, ok := me[key]; ok {

		if m, ok := val.(map[int]int); ok {
			/* act on str */
			return m, nil
		} else {
			return nil, ErrorWrongType
		}
	} else {
		return nil, ErrorNoKeyFound
	}
}

func NewTtfInfo() TtfInfo {
	info := make(TtfInfo)
	return info
}
