package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"strings"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("# 目录\n")
	readDir(".", root, &buf)
	buf.WriteString(footer)
	err := ioutil.WriteFile("README.md", buf.Bytes(), 0644)
	checkErr(err)
}

func readDir(dirname string, level int, buf *bytes.Buffer) {
	fileLis, err := ioutil.ReadDir(dirname)
	checkErr(err)
	for _, fi := range fileLis {
		if validName(fi.Name()) {
			if level == root {
				buf.WriteString("\n")
				if fi.IsDir() {
					buf.WriteString("## ")
				}
			} else {
				buf.WriteString(strings.Repeat("  ", level))
				buf.WriteString("- ")
			}
			buf.WriteString(fmt.Sprintf("[%s](%s)\n",
				prettyName(fi.Name()),
				path.Join(url.PathEscape(dirname), url.PathEscape(fi.Name())),
			))
			if fi.IsDir() {
				readDir(path.Join(dirname, fi.Name()), level+1, buf)
			}
		}
	}
}

func validName(name string) bool {
	name = strings.ToLower(name)
	exclude := map[string]bool{
		"license": true,
	}
	return !(exclude[name] ||
		strings.HasPrefix(name, ".") ||
		strings.HasSuffix(name, ".md") ||
		strings.HasSuffix(name, ".go"))
}

func prettyName(name string) string {
	name = strings.ReplaceAll(name, "+", " ")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.Title(name)
	return name
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

const root = 0

const footer = `
## &copy;2018 Shuo. All rights reserved.
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fopenset%2Fbooks.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fopenset%2Fbooks?ref=badge_shield)

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fopenset%2Fbooks.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fopenset%2Fbooks?ref=badge_large)
`
