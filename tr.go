package tr

import (
	"fmt"
)

// Init is used to instantiate an Engine with its locales directory, as well as the
// default locale.
func Init(path, defaultLocale string, trimOptional ...bool) (*Engine, error) {
	trim := false
	if len(trimOptional) > 0 {
		trim = true
	}

	engine, err := NewEngine(path, defaultLocale, trim)
	if err != nil {
		return nil, fmt.Errorf("Could not load new engine. Error: %v", err)
	}

	return engine, nil
}