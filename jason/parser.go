package jason

import (
	"fmt"
)

func Parse(tokens []interface{}) (interface{}, []interface{}) {
	t := tokens[0]

	if t == string(JSONLeftBracket) {
		return parseArray(tokens[1:])
	} else if t == string(JSONLeftBrace) {
		return parseObject(tokens[1:])
	}

	return t, tokens[1:]
}

func parseArray(tokens []interface{}) ([]interface{}, []interface{}) {
	jsonArray := []interface{}{}

	var json interface{}
	t := tokens[0]
	if t == JSONRightBracket {
		return jsonArray, tokens[1:]
	}

	for {
		json, tokens = Parse(tokens)
		jsonArray = append(jsonArray, json)

		t = tokens[0]
		if t == string(JSONRightBracket) {
			return jsonArray, tokens[1:]
		} else if t != string(JSONComma) {
			panic(fmt.Sprintf("Expected a comma, got: %v", t))
		}

		tokens = tokens[1:]
	}
}

func parseObject(tokens []interface{}) (map[string]interface{}, []interface{}) {
	out := make(map[string]interface{})

	var val interface{}

	t := tokens[0]
	if t == string(JSONRightBrace) {
		return out, tokens[1:]
	}

	for {
		key := tokens[0]
		if isString(key) {
			tokens = tokens[1:]
		} else {
			panic(fmt.Sprintf("expected string key, got: %v", key))
		}

		if tokens[0] != string(JSONColon) {
			panic(fmt.Sprintf("expected colon, got: %v", tokens[0]))
		}

		val, tokens = Parse(tokens[1:])
		out[key.(string)] = val

		t = tokens[0]
		if t == string(JSONRightBrace) {
			return out, tokens[1:]
		} else if t != string(JSONComma) {
			panic(fmt.Sprintf("expected comma after pair in object, got: %v", t))
		}

		tokens = tokens[1:]
	}
}

func isString(k interface{}) bool {
	switch k.(type) {
	case string:
		return true
	}

	return false
}
