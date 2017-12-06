package tr

var DefaultEngine *Engine

func init() {
	DefaultEngine = &Engine{}
}

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

func Tr(path string) string {
	return DefaultEngine.Tr(path)
}

func Lang(localeName string) *Locale {
	return DefaultEngine.Lang(localeName)
}
