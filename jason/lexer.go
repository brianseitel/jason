package jason

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// def lex(string):
//     tokens = []

//     while len(string):
//         json_string, string = lex_string(string)
//         if json_string is not None:
//             tokens.append(json_string)
//             continue

//         # TODO: lex booleans, nulls, numbers

//         if string[0] in JSON_WHITESPACE:
//             string = string[1:]
//         elif string[0] in JSON_SYNTAX:
//             tokens.append(string[0])
//             string = string[1:]
//         else:
//             raise Exception('Unexpected character: {}'.format(string[0]))

//     return tokens

// Lex a string
func Lex(input string) []interface{} {
	tokens := []interface{}{}

	var err error
	var jsonString string
	var jsonBool bool
	var jsonNull bool
	var jsonNumber interface{}

	rawString := input
	i := 0
	for {
		if len(rawString) == 0 {
			break
		}

		jsonString, rawString, err = lexString(rawString)
		if err == nil {
			tokens = append(tokens, jsonString)
			continue
		}

		jsonNumber, rawString = lexNumber(rawString)
		if jsonNumber != nil {
			tokens = append(tokens, jsonNumber)
			continue
		}

		jsonBool, rawString, err = lexBool(rawString)
		if err == nil {
			tokens = append(tokens, jsonBool)
			continue
		}

		jsonNull, rawString = lexNull(rawString)
		if jsonNull {
			tokens = append(tokens, nil)
			continue
		}

		if isWhitespace(rawString[0]) {
			rawString = rawString[1:]
		} else if isJSONSyntax(rawString[0]) {
			tokens = append(tokens, string(rawString[0]))
			rawString = rawString[1:]
		} else {
			panic(fmt.Sprintf("Unexpected character: %v", string(rawString[0])))
		}
		i++
	}

	return tokens
}

// JSON constants
const (
	JSONQuote        = '"'
	JSONTrue         = "true"
	JSONFalse        = "false"
	JSONNull         = "null"
	JSONNumber       = "01234567890-e."
	JSONComma        = ','
	JSONColon        = ':'
	JSONLeftBracket  = '['
	JSONRightBracket = ']'
	JSONLeftBrace    = '{'
	JSONRightBrace   = '}'
)

func isWhitespace(input byte) bool {
	return input == ' ' || input == '\t' || input == '\n' || input == '\r'
}

func isJSONSyntax(input byte) bool {
	return input == JSONQuote ||
		input == JSONComma ||
		input == JSONColon ||
		input == JSONLeftBrace ||
		input == JSONRightBrace ||
		input == JSONLeftBracket ||
		input == JSONRightBracket
}

func isNumber(input byte) bool {
	validNumberValues := strings.Split(JSONNumber, "")

	numMap := map[string]bool{}

	for _, num := range validNumberValues {
		numMap[num] = true
	}

	_, ok := numMap[string(input)]

	return ok
}

func lexString(input string) (string, string, error) {
	out := ""
	if input[0] == JSONQuote {
		input = input[1:]
	} else {
		return "", input, errors.New("not a string")
	}

	for i := 0; i < len(input); i++ {
		if input[i] == JSONQuote {
			return out, input[len(out)+1:], nil
		} else {
			out += string(input[i])
		}
	}

	panic("Unexpected end-of-string quote")
}

func lexNumber(input string) (interface{}, string) {
	out := ""

	for i := 0; i < len(input); i++ {
		if isNumber(input[i]) {
			out += string(input[i])
		} else {
			break
		}
	}

	rest := input[len(out):]

	if len(out) == 0 {
		return nil, input
	}

	if strings.Contains(out, ".") {
		f, _ := strconv.ParseFloat(out, 64)
		return f, rest
	}

	i, _ := strconv.Atoi(out)
	return i, rest
}

func lexBool(input string) (bool, string, error) {
	if len(input) >= len(JSONTrue) && input[:len(JSONTrue)] == JSONTrue {
		return true, input[len(JSONTrue):], nil
	}

	if len(input) >= len(JSONFalse) && input[:len(JSONFalse)] == JSONFalse {
		return false, input[len(JSONFalse):], nil
	}

	return false, input, errors.New("not a bool")
}

func lexNull(input string) (bool, string) {
	if len(input) >= len(JSONNull) && input[:len(JSONNull)] == JSONNull {
		return true, input[len(JSONNull):]
	}
	return false, input
}
