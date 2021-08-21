package flatten

func Flatten(array []interface{}) []interface{} {
	res := make([]interface{}, 0)
	for _, value := range array {
		if value == nil {
			continue
		}
		switch v := value.(type) {
		case []interface{}:
			tmp := Flatten(v)
			res = append(res, tmp...)
		default:
			res = append(res, value)
		}
	}
	return res
}
