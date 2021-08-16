package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// Global Variables
var (
	output string
	pre    bool
	block  bool
	list   bool
	err    error

	filePath    string
	outPath     string
	titleCase   string
	customTitle string
)

func features(x uint) {
	switch os.Args[x] {
	case "-t", "--title":
		if len(os.Args)-int(x) != 2 {
			fmt.Println(noInput)
			os.Exit(0)
		}
		titleCase = "custom"
		filePath = os.Args[x+1]

	case "-c", "--setcase":
		if len(os.Args)-int(x) != 3 {
			fmt.Println(noInput)
			os.Exit(0)
		}
		titleCase = os.Args[x+1]
		filePath = os.Args[x+2]

	default:
		if strings.HasPrefix(os.Args[1], "-") || len(os.Args) != 2 {
			fmt.Println(noInput)
			os.Exit(0)
		}
		titleCase = "none"
		filePath = os.Args[1]
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(noInput)
		os.Exit(0)
	}
	switch os.Args[1] {
	case "--help":
		// Print Help
		if len(os.Args) > 2 {
			fmt.Println("WARNING: --help doesn't expect any more arguments.")
		}
		fmt.Println(help)
		os.Exit(0)

	case "--generate":
		// Write template and css variable to file.
		if len(os.Args) > 2 {
			fmt.Println("WARNING: --generate option don't expect any more arguments.")
		}
		err = os.WriteFile("html_template.tmpl", []byte(templateFile), 0644)
		if err != nil {
			fmt.Println("ERROR: failed to generate template file")
			os.Exit(1)
		}
		err = os.WriteFile("style.css", []byte(cssFile), 0644)
		if err != nil {
			fmt.Println("ERROR: failed to generate css file")
			os.Exit(1)
		}
		os.Exit(0)

	case "-o", "--output":
		if len(os.Args) < 4 {
			fmt.Println(noInput)
			os.Exit(0)
		}
		outPath = os.Args[2]
		_, err = os.Stat(outPath)
		if os.IsNotExist(err) {
			fmt.Println("ERROR: output directory does not exist")
			os.Exit(1)
		}
		if strings.HasPrefix(os.Args[3], "-") {
			features(3)
		} else {
			filePath = os.Args[3]
			titleCase = "none"
		}

	default:
		features(1)
	}

	f, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("ERROR: input file/directory doesn't exist")
		os.Exit(1)
	}
	if !f.IsDir() {
		generateOutput()
		os.Exit(0)
	}

	files, err := os.ReadDir(filePath)
	if err != nil {
		fmt.Println("ERROR: can't read input directory")
	}
	copyPath := filePath
	for _, file := range files {
		if !file.IsDir() {
			filePath = path.Join(copyPath, file.Name())
			generateOutput()
		}
	}
}
