package source

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrParse = fmt.Errorf("no match")
)

var (
	DefaultParse = Parse
	DefaultRegex = Regex
)

// Regex matches the following pattern:
//
//	123_name.up.ext
//	123_name.down.ext
//
// ^(((V|U|B)[0-9]+)|(R))__(.*)\.(up|down)\.(.*)$
var Regex = regexp.MustCompile(`^(([VUB])([0-9]+)|(R))__(.*)\.(up|down)\.(.*)$`)

// Parse returns Migration for matching Regex pattern.
func Parse(raw string) (*Migration, error) {
	m := Regex.FindStringSubmatch(raw)
	if len(m) == 8 {
		versionUint64, err := strconv.ParseUint(m[3], 10, 64)
		if err != nil {
			return nil, err
		}
		return &Migration{
			Version:    uint(versionUint64),
			Identifier: m[5],
			Direction:  Direction(m[6]),
			Raw:        raw,
		}, nil
	}
	return nil, ErrParse
}
