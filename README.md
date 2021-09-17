# meow

meow is a CLI based markdown to html generator. But with a few quirks, the markdown sytanx it uses is based on gemini markdown with some extra features.

## Getting Started

Note: meow is developed on and for Linux.

### Requirments

* Go (golang)
* GOPATH setup must be done
* GOPATH/bin in PATH variable

### Compile + Install

Execute clean-install.sh

```
./clean-install.sh
```

Enjoy!!

## Usage

When using meow, it requires a template file which may be generated using:

```
meow --generate
```

Edit the template file to generate perfect html pages according to your liking.

## Features

* Supports batch converting files when a directory is selected.
Note: Files must have '.gmi' extension.
* Output files have same name but '.html' extension instead of '.gmi'
* Output directory can be selected using '-o' flag.

```
meow -o ./out/output-dir file.gmi
```

* By default filename is used in 'title' tag in HTML.
* If the title needs to be changed. '-c' flag followed by the case 'title, upper, lower or custom' may be used.

```
meow -c title ./src/index.gmi
```

* Refer --help option for all the flags.

```
meow --help
```

## gemini Markdown Cheatsheet

### Headings:

```
# This is a <h1> tag
## This is a <h2> tag
### This is a <h3> tag
```

### Paragraphs

```
Random text without line break is always considered paragraph and is word wrapped based on viewport.

Empty lines are left empty.
```

### List

```
* Item 1
* Item 2
* Item 3
```

### Code/ Pre-formatted text

```
triple backtick
This is the code.
	or
Pre-formatted text for ASCII drawings
triple backtick
```

### Blockquote

```
> This is text in a Blockquote
```

### Hyperlink

When 'Displayed Text' is left empty the url is used as the display text too.

```
=> <url of destination> <Displayed Text>
=> https://www.google.com Go to google
```

## Extra Features

### Horizontal Line

I really needed this for my website.

```
_
```

### Footer

Might be used for copyright text. meow ends parsing immediately after recieving footer or EOF

```
^ This is Footer Text
```

## Note

* This project was made to create my static website.
* Extremely simple parsing is used, I promote editing 'helpers.go' file if you need any additional features.
