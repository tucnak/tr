package tr

// DefaultEngine is what used when calling package-scope Tr().
var DefaultEngine *Engine

func init() {
	DefaultEngine = &Engine{}
}

// TrimEnd is an optional 3rd argument to trim \n ending.
const TrimEnd = true

// Init is used to set the locales directory, as well as the
// default locale.
func Init(path, defaultLocale string, trimOptional ...bool) error {
	trim := false
	if len(trimOptional) > 0 {
		trim = true
	}

	engine, err := NewEngine(path, defaultLocale, trim)
	if err != nil {
		return err
	}

	DefaultEngine = engine
	return nil
}

// Tr provides default locale's translation of path.
func Tr(path string) string {
	return DefaultEngine.Tr(path)
}

// Lang returns a *Locale by name.
func Lang(localeName string) *Locale {
	return DefaultEngine.Lang(localeName)
}
