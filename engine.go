package tr

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func NewEngine(path, defaultLocale string, trim bool) (*Engine, error) {
	e := &Engine{
		Path:  path,
		Langs: map[string]*Locale{},
	}

	path, _ = filepath.Abs(path)

	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			if !file.IsDir() {
				continue
			}

			name := file.Name()
			paths := []string{}

			err = filepath.Walk(path+"/"+name, func(fpath string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					paths = append(paths, fpath)
				}

				return err
			})

			if err != nil {
				return nil, errors.Wrap(err, "tr: couldn't walk thru files")
			}

			var c *Locale

			c, err = NewLocale(path, name, paths, trim)
			if err != nil {
				return nil, err
			}

			e.Langs[name] = c
		}

		e.DefaultLocale = e.Langs[defaultLocale]
	} else {
		return nil, errors.Wrap(err, "tr: couldn't open locales")
	}

	return e, nil
}

type Engine struct {
	Path          string
	DefaultLocale *Locale

	Langs map[string]*Locale
}

func (e *Engine) Lang(localeName string) *Locale {
	return e.Langs[localeName]
}

func (e *Engine) Tr(path string) string {
	if e.Langs == nil {
		panic("tr: default engine is not sent, see tr.Init()")
	}

	return e.DefaultLocale.Tr(path)
}
