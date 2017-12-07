# tr
> Easy drop-in i18n solution for Go applications.

[![GoDoc](https://godoc.org/github.com/tucnak/tr?status.svg)](https://godoc.org/github.com/tucnak/tr)

I was looking for a real easy way to provide i18n support for my Telegram
bot, for which the data is pretty much a set of 20 different text messages.
I couldn't find a single solution that would utilize the file system.
Here's how `tr` works:

1. You have to create a locales directory, e.g. `$ tree lang`:
   ```
   lang
   ├── en
   │   ├── hello.txt
   │   └── inner
   │       └── text.txt
   ├── fr
   │   ├── hello.txt
   │   └── inner
   │       └── text.txt
   └── ru
       ├── hello.txt
       └── inner
           └── text.html

   6 directories, 6 files
   ```

   Your files could be of any extension, it doesn't really matter,
   since `tr` ignores extensions anyway.

2. Init `tr` properly in your program:
    ```go
	package main

	import (
		"fmt"
		"os"

		"github.com/tucnak/tr"
	)

	func init() {
		// tr.Init(localesDirectory, defaultLocale)
		if err := tr.Init("lang", "en"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
    ```

3. Use simple syntax for i18n:
	```go
    // Inline syntax:
	fmt.Println("In English:", tr.Lang("en").Tr("hello"))
	fmt.Println("In French:", tr.Lang("fr").Tr("hello"))
	fmt.Println("In Russian:", tr.Lang("ru").Tr("hello"))

	// Shadowing
	tr := tr.Lang("fr")
	fmt.Println(tr.Tr("inner/text"))
	```

Pass an optional third `true` argument to `tr.Init()` if you wish
to trim all `\n`s from the end of the string returned.
