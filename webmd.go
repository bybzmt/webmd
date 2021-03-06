package main

import (
	//"encoding/base64"
	"flag"
	"github.com/russross/blackfriday"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"regexp/syntax"
	"strings"
)

var basedir = flag.String("dir", "./", "markdown files dir")
var layout = flag.String("layout", "./assets/layout.tpl", "markdown files dir")
var addr = flag.String("addr", ":8080", "Listen addr:port")
var dot = flag.String("dot", "dot", "graphviz path")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filename := path.Join(*basedir, path.Clean(r.URL.Path))

		//处理dot
		regx, err := syntax.Parse(".dot$", syntax.PerlX|syntax.MatchNL|syntax.UnicodeGroups)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		//log.Println(regx.String())
		reg, err := regexp.Compile(regx.String())
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		if reg.MatchString(filename) {
			parseDot(filename, w)
			return
		}

		f, err := os.Open(filename)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		if fi.IsDir() {
			http.ServeFile(w, r, filename)
			return
		}

		if strings.ToLower(path.Ext(filename)) != ".md" {
			http.ServeFile(w, r, filename)
			return
		}

		in, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		//处理markdown
		out := parseMarkdown(in)

		data := make(map[string]interface{})
		data["md"] = template.HTML(string(out))

		tpl, err := template.ParseFiles(*layout)
		tpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func parseDot(file string, w io.Writer) {
	c := exec.Command(*dot, "-Tpng", file)

	/*
		stdin, err := c.StdinPipe()
		if err != nil {
			w.Write("exec dot error:" + err.Error())
			return
		}

		go func() {
			defer stdin.Close()

			start := len("<dot>")
			end := len(in) - len("</dot>")

			_, err := stdin.Write(in[start:end])
			if err != nil {
				log.Println("exec dot error:" + err.Error())
			}
		}()
	*/

	stdout, err := c.StdoutPipe()
	if err != nil {
		w.Write([]byte("exec dot error:" + err.Error()))
		return
	}

	go func() {
		defer stdout.Close()

		io.Copy(w, stdout)
	}()

	err = c.Run()
	if err != nil {
		w.Write([]byte("exec dot error:" + err.Error()))
	}
}

func parseMarkdown(in []byte) []byte {

	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_DASHES
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_HEADER_IDS
	extensions |= blackfriday.EXTENSION_BACKSLASH_LINE_BREAK
	extensions |= blackfriday.EXTENSION_DEFINITION_LISTS

	//htmlFlags |= blackfriday.HTML_TOC
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

	return blackfriday.Markdown(in, renderer, extensions)
}
