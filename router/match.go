package router

import "strings"

func match(pattern, path string) (map[string]string, bool) {
	patternPart := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternPart) != len(pathParts) {
		return nil, false
	}

	params := map[string]string{}
	for i := range patternPart {
		pattern := patternPart[i]
		part := pathParts[i]
		if strings.HasPrefix(pattern, "{") && strings.HasSuffix(pattern, "}") {
			key := pattern[1 : len(pattern)-1]
			params[key] = part

		} else if pattern != part {
			return nil, false
		}

	}
	return params, true

}
