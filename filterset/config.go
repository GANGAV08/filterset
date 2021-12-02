package filterset

import (
	"fmt"

	"github.com/GANGAV08/regexp/regexp"
	"github.com/GANGAV08/strict/strict"
)

type MatchType string

const (
	Regexp MatchType = "regexp"

	Strict MatchType = "strict"

	MatchTypeFieldName = "match_type"
)

var (
	validMatchTypes = []MatchType{Regexp, Strict}
)

type Config struct {
	MatchType    MatchType      `mapstructure:"match_type"`
	RegexpConfig *regexp.Config `mapstructure:"regexp"`
}

func CreateFilterSet(filters []string, cfg *Config) (FilterSet, error) {
	switch cfg.MatchType {
	case Regexp:
		return regexp.NewFilterSet(filters, cfg.RegexpConfig)
	case Strict:

		return strict.NewFilterSet(filters), nil
	default:
		return nil, fmt.Errorf("unrecognized %v: '%v', valid types are: %v", MatchTypeFieldName, cfg.MatchType, validMatchTypes)
	}
}
