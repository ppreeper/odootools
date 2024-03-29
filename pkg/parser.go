// odootools - a cli tool library to access Odoo instances via Json RPC
// Copyright (C) 2023  Peter Preeper

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package pkg

import (
	"errors"
	"strconv"
	"strings"
)

func parseFields(field string) (fields []string) {
	if field != "" {
		fields = strings.Split(field, ",")
	} else {
		fields = []string{}
	}
	return
}

func parseFilter(filter string) (filters []any, err error) {
	filter = strings.TrimSpace(filter)

	// pre-parse
	sqBCount, sqBDepth, parenCount, _ := countBrackets(filter)
	if len(filter) == 0 {
		return
	}
	if !(len(filter) > 4) {
		return nil, errors.New("invalid filter length")
	}
	if !(sqBCount == 0 && sqBDepth == 1) {
		return nil, errors.New("invalid filter format")
	}
	if !(parenCount == 0) {
		return nil, errors.New("invalid filter format")
	}

	// lex
	tokens, err := lexer(filter)
	CheckErr(err)

	parenCount = 0
	var arg []any
	var sList []any
	for i := 1; i < len(tokens)-1; i++ {
		f := tokens[i]
		switch {
		case f == "(":
			parenCount += 1
			if parenCount < 2 {
				arg = []any{}
			} else {
				sList = []any{}
			}
		case f == ")":
			parenCount -= 1
			if parenCount == 1 {
				arg = append(arg, sList)
			} else {
				filters = append(filters, arg)
			}
		case f == "&":
			filters = append(filters, f)
		case f == "|":
			filters = append(filters, f)
		case f == "!":
			filters = append(filters, f)
		default:
			if parenCount > 1 {
				if IsInt(f) {
					fi, _ := strconv.Atoi(f)
					sList = append(sList, fi)
				} else if IsNumeric(f) {
					fi, _ := strconv.ParseFloat(f, 64)
					sList = append(sList, fi)
				} else if IsBool(f) {
					fb, _ := strconv.ParseBool(f)
					sList = append(sList, fb)
				} else {
					sList = append(sList, f)
				}
			} else {
				if IsInt(f) {
					fi, _ := strconv.Atoi(f)
					arg = append(arg, fi)
				} else if IsNumeric(f) {
					fi, _ := strconv.ParseFloat(f, 64)
					arg = append(arg, fi)
				} else if IsBool(f) {
					fb, _ := strconv.ParseBool(f)
					arg = append(arg, fb)
				} else {
					arg = append(arg, f)
				}
			}
		}
	}
	return
}

func countBrackets(ff string) (sqBCount int, sqBDepth int, parenCount int, parenDepth int) {
	for _, f := range ff {
		switch {
		case string(f) == "[":
			sqBCount += 1
			sqBDepth += 1
		case string(f) == "]":
			sqBCount -= 1
		case string(f) == "(":
			parenCount += 1
			parenDepth += 1
		case string(f) == ")":
			parenCount -= 1
		}
	}
	return
}

func lexer(s string) ([]string, error) {
	tokens := []string{}
	bb := []byte(s)
	for i := 0; i < len(bb); i++ {
		b := bb[i]
		switch {
		case string(b) == ",":
			ffwd, sToken := lexToken(bb[i:len(bb)-1], []string{"(", "'", ",", ")"})
			if len(sToken) > 0 {
				tokens = append(tokens, strings.TrimSpace(sToken))
				if string(bb[i+ffwd-1]) == ")" {
					i += ffwd
				}
			}
		case string(b) == "'":
			ffwd, sToken := lexToken(bb[i:len(bb)-1], []string{"'"})
			if len(sToken) > 0 {
				tokens = append(tokens, strings.TrimSpace(sToken))
				i += ffwd
			}
		case string(b) == "[":
			tokens = append(tokens, string(b))
		case string(b) == "]":
			tokens = append(tokens, string(b))
		case string(b) == "(":
			tokens = append(tokens, string(b))
		case string(b) == ")":
			tokens = append(tokens, string(b))
		default:
			continue
		}
	}
	return tokens, nil
}

func lexToken(bb []byte, endTerms []string) (ffwd int, sToken string) {
	for i := 1; i < len(bb); i++ {
		b := bb[i]
		for _, t := range endTerms {
			if string(b) == t {
				return i, string(bb[1:i])
			}
		}
	}
	return
}
