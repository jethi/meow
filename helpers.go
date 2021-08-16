package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

type Content struct {
	Title string
	Data  string
}

func set(saveb bool, savel bool) {
	if block && !saveb {
		block = !block
		output += "\n\t\t</blockquote>\n"
	}
	if list && !savel {
		list = !list
		output += "\t\t</ul>\n"
	}
}

func generateOutput() {
	if !fileExists(filePath) {
		os.Exit(0)
	}
	if !strings.HasSuffix(filePath, ".gmi") {
		return
	}
	var container Content
	var fileName string
	fileName = strings.TrimSuffix(path.Base(filePath), ".gmi")
	titleCase = strings.ToLower(titleCase)
	switch titleCase {
	case "title":
		container.Title = strings.Title(fileName)
	case "lower":
		container.Title = strings.ToLower(fileName)
	case "upper":
		container.Title = strings.ToUpper(fileName)
	case "custom":
		fmt.Printf("Please enter the title for %s.gmi:", fileName)
		in := bufio.NewReader(os.Stdin)
		customTitle, _ = in.ReadString('\n')
		customTitle = customTitle[:len(customTitle)-1]
		container.Title = customTitle
	case "none":
		container.Title = fileName
	default:
		fmt.Println("meow: please enter a valid case (title | lower | upper)")
		fmt.Println("Try 'meow --help' for more information.")
		os.Exit(0)
	}
	container.Title = "\t<title>" + container.Title + "</title>\n"

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("ERROR: could not read the file:", err)
		os.Exit(1)
	}
	data := strings.Split(string(file), "\n")
	for i, j := range data {
		if pre == false {
			j = strings.TrimSpace(j)
		}
		if strings.HasPrefix(j, "```") {
			if pre == false {
				set(false, false)
				pre = true
				output = output + "<pre>\n"
				continue
			} else {
				pre = false
				output = output + "</pre>\n"
				continue
			}
		}
		if pre == true {
			output = output + j + "\n"
		} else if strings.HasPrefix(j, ">") {
			set(true, false)
			if block == false {
				output = output + "\t\t<blockquote>\n" + "\t\t\t" + strings.TrimSpace(j[1:])
				block = true
			} else {
				output += "<br/>\n" + "\t\t\t" + strings.TrimSpace(j[1:])
			}
		} else if strings.HasPrefix(j, "*") {
			set(false, true)
			if list == false {
				output = output + "\t\t<ul>\n"
				list = true
			}
			output = output + "\t\t\t<li>" + strings.TrimSpace(j[1:]) + "</li>\n"
		} else if strings.HasPrefix(j, "_") {
			set(false, false)
			output = output + "\t\t<hr/>\n"
		} else if strings.HasPrefix(j, "# ") {
			set(false, false)
			output = output + "\t\t<h1>" + strings.TrimSpace(j[1:]) + "</h1>\n"
		} else if strings.HasPrefix(j, "## ") {
			set(false, false)
			output = output + "\t\t<h2>" + strings.TrimSpace(j[2:]) + "</h2>\n"
		} else if strings.HasPrefix(j, "### ") {
			set(false, false)
			output = output + "\t\t<h3>" + strings.TrimSpace(j[3:]) + "</h3>\n"
		} else if strings.HasPrefix(j, "=>") {
			set(false, false)
			matter := strings.SplitN(strings.TrimSpace(j[2:]), " ", 2)
			if len(matter) == 1 {
				output += "\t\t<a href=\"" + matter[0] + "\">" + matter[0] + "</a><br/>\n"
			} else {
				output += "\t\t<a href=\"" + matter[0] + "\">" + matter[1] + "</a><br/>\n"
			}
		} else if i != len(data)-1 {
			// Its a Paragraph
			set(false, false)
			if len(j) == 0 {
				// This one goes for blank lines.
				output = output + "\t\t<br/>\n"
			} else {
				output = output + "\t\t<p>\n\t\t\t" + strings.TrimSpace(j) + "\n\t\t</p>\n"
			}
		}
	}
	set(false, false)
	fileName = fileName + ".html"
	container.Data = output
	// TODO: Check hat happens if directory doesn't exist.
	t, err := template.ParseFiles("html_template.tmpl")
	if err != nil {
		fmt.Println("ERROR: no template file exists")
		fmt.Println("Try 'meow --generate' to generate default template.")
		os.Exit(1)
	}
	f, err := os.OpenFile(path.Join(path.Clean(outPath), fileName), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR: failed to open file:", err)
		os.Exit(1)
	}
	err = t.Execute(f, container)
	if err != nil {
		fmt.Println("ERROR: couldn't create html file:", err)
		os.Exit(1)
	}
	err = f.Close()
	if err != nil {
		fmt.Println("ERROR: counldn't close file:", err)
		os.Exit(1)
	}
	output = ""
}

func fileExists(filePath string) bool {
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("ERROR: please provide a valid input file")
		fmt.Println("Try 'meow --help' for more information.")
		return false
	}
	return true
}
