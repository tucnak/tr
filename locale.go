package tr

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/armon/go-radix"
	"github.com/pkg/errors"
)

func NewLocale(root, name string, paths []string, trim bool) (*Locale, error) {
	c := &Locale{
		Root: root,
		Name: name,
		Trim: trim,
		tree: radix.New(),
	}

	cut := len(root + "/" + name + "/")

	for _, filename := range paths {
		if _, err := os.Stat(filename); err == nil {
			var file *os.File

			file, err = os.Open(filename)
			if err != nil {
				return nil, errors.Wrap(err, "tr: can't open file")
			}

			var data []byte
			data, err = ioutil.ReadAll(file)

			if err != nil {
				return nil, errors.Wrap(err, "tr: can't open file")
			}

			rel := filename[cut:]
			for i := len(rel) - 1; i > 0; i-- {
				if rel[i] == '.' {
					rel = rel[:i]
					break
				}

				if rel[i] == '/' || rel[i] == '\\' {
					break
				}
			}

			c.tree.Insert(rel, string(data))
		} else {
			return nil, errors.Wrap(err, "tr: can't open file")
		}
	}

	return c, nil
}

type Locale struct {
	Root string
	Name string
	Trim bool

	tree *radix.Tree
}

func (c *Locale) Tr(path string) string {
	obj, ok := c.tree.Get(path)
	if !ok {
		panic("tr: no translation for " + path)
	}

	text := obj.(string)

	if c.Trim {
		return strings.TrimRight(text, "\n")
	}

	return text
}
