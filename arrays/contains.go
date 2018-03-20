package arrays

// Contains 渡されたarraynの中にvalが存在したらtrueを返す
func Contains(arr interface{}, val interface{}) bool {
	switch art := arr.(type) {
	case []interface{}:
		return false
	case []int:
		v, ok := val.(int)
		if !ok {
			return false
		}
		for _, key := range art {
			if key == v {
				return true
			}
		}
	case []string:
		v, ok := val.(string)
		if !ok {
			return false
		}
		for _, key := range art {
			if key == v {
				return true
			}
		}
	}
	return false
}
