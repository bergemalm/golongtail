package longtailutils

import (
	"log"
	"regexp"

	"github.com/DanEngelbrecht/golongtail/longtaillib"
)

type regexPathFilter struct {
	compiledIncludeRegexes []*regexp.Regexp
	compiledExcludeRegexes []*regexp.Regexp
}

// MakeRegexPathFilter ...
func MakeRegexPathFilter(includeFilterRegEx string, excludeFilterRegEx string) (longtaillib.Longtail_PathFilterAPI, error) {
	regexPathFilter := &regexPathFilter{}
	if includeFilterRegEx != "" {
		compiledIncludeRegexes, err := splitRegexes(includeFilterRegEx)
		if err != nil {
			return longtaillib.Longtail_PathFilterAPI{}, err
		}
		regexPathFilter.compiledIncludeRegexes = compiledIncludeRegexes
	}
	if excludeFilterRegEx != "" {
		compiledExcludeRegexes, err := splitRegexes(excludeFilterRegEx)
		if err != nil {
			return longtaillib.Longtail_PathFilterAPI{}, err
		}
		regexPathFilter.compiledExcludeRegexes = compiledExcludeRegexes
	}
	if len(regexPathFilter.compiledIncludeRegexes) > 0 || len(regexPathFilter.compiledExcludeRegexes) > 0 {
		return longtaillib.CreatePathFilterAPI(regexPathFilter), nil
	}
	return longtaillib.Longtail_PathFilterAPI{}, nil
}

func (f *regexPathFilter) Include(rootPath string, assetPath string, assetName string, isDir bool, size uint64, permissions uint16) bool {
	for _, r := range f.compiledExcludeRegexes {
		if r.MatchString(assetPath) {
			log.Printf("INFO: Skipping `%s`", assetPath)
			return false
		}
	}
	if len(f.compiledIncludeRegexes) == 0 {
		return true
	}
	for _, r := range f.compiledIncludeRegexes {
		if r.MatchString(assetPath) {
			return true
		}
	}
	log.Printf("INFO: Skipping `%s`", assetPath)
	return false
}

func splitRegexes(regexes string) ([]*regexp.Regexp, error) {
	var compiledRegexes []*regexp.Regexp
	m := 0
	s := 0
	for i := 0; i < len(regexes); i++ {
		if (regexes)[i] == '\\' {
			m = -1
		} else if m == 0 && (regexes)[i] == '*' {
			m++
		} else if m == 1 && (regexes)[i] == '*' {
			r := (regexes)[s:(i - 1)]
			regex, err := regexp.Compile(r)
			if err != nil {
				return nil, err
			}
			compiledRegexes = append(compiledRegexes, regex)
			s = i + 1
			m = 0
		} else {
			m = 0
		}
	}
	if s < len(regexes) {
		r := (regexes)[s:]
		regex, err := regexp.Compile(r)
		if err != nil {
			return nil, err
		}
		compiledRegexes = append(compiledRegexes, regex)
	}
	return compiledRegexes, nil
}