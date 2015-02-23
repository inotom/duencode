// data uri scheme encoder
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const DATA_URI_SCHEME_TEMPLATE = "data:image/%s;base64,"
const MAX_NOT_FLAG_COUNT = 1

// encoding to base64 data from image file
func base64encode(file os.File) string {
	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

// make data uri scheme string from file extention
func dataUriScheme(file os.File) string {
	fi, _ := file.Stat()
	ext := strings.Trim(filepath.Ext(fi.Name()), ".")
	return fmt.Sprintf(DATA_URI_SCHEME_TEMPLATE, ext)
}

func encode(filePath string, isPlain bool) {
	scheme := ""

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	if !isPlain {
		scheme = dataUriScheme(*file)
	}

	fmt.Printf("%s%s\n", scheme, base64encode(*file))
}

func main() {
	var isPlainFormat bool

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
  %s [OPTIONS] FILE

Options
  -h: Show this message
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&isPlainFormat, "p", false, "output with plain format")
	flag.Parse()

	if flag.NArg() != MAX_NOT_FLAG_COUNT {
		flag.Usage()
		os.Exit(1)
	}

	encode(flag.Arg(0), isPlainFormat)
}
