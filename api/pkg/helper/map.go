package helper

type H map[string]interface{}

func MergeMap[T comparable](params1, params2 map[T]interface{}) map[T]interface{} {
	newParams := make(map[T]interface{})
	for i, v := range params1 {
		for j, w := range params2 {
			if i == j {
				newParams[i] = w
			} else {
				if _, ok := newParams[i]; !ok {
					newParams[i] = v
				}
				if _, ok := newParams[j]; !ok {
					newParams[j] = w
				}
			}
		}
	}
	return newParams
}
