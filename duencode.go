// data uri scheme encoder
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const DATA_URI_SCHEME_TEMPLATE = "data:%s;base64,"
const MAX_NOT_FLAG_COUNT = 1

// read []byte data from file
func readFileData(file os.File) []byte {
	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	return data
}

// make data uri scheme string from file data
func dataUriScheme(data []byte) string {
	mimeType := http.DetectContentType(data)
	return fmt.Sprintf(DATA_URI_SCHEME_TEMPLATE, mimeType)
}

func encode(filePath string, isPlain bool, noRet bool) {
	scheme := ""
	ret := "\n"

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	data := readFileData(*file)

	if !isPlain {
		scheme = dataUriScheme(data)
	}

	if noRet {
		ret = ""
	}

	fmt.Printf("%s%s%s", scheme, base64.StdEncoding.EncodeToString(data), ret)
}

func main() {
	var isPlainFormat bool
	var noRet bool

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
  %s [OPTIONS] FILE

Options
  -h: Show this message
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&isPlainFormat, "p", false, "output with plain format")
	flag.BoolVar(&noRet, "n", false, "without line break")
	flag.Parse()

	if flag.NArg() != MAX_NOT_FLAG_COUNT {
		flag.Usage()
		os.Exit(1)
	}

	encode(flag.Arg(0), isPlainFormat, noRet)
}
