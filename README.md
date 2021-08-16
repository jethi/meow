# meow

meow is a CLI based markdown to html generator. But with a few quirks, the markdown sytanx it uses is based on gemini markdown with some extra featurs.

When using meow, it requires a template file which may be generated using:
```
meow --generate
```

Edit the template file to generate perfect html pages according to you liking.

## Features

* Supports batch converting files when a directory is selected.
Note: Files must have '.gmi' extension.
* Output files have same name but '.html' extension instead of '.gmi'
* Output directory can be selected using '-o' flag.

```
meow -o ./out/output-dir file.gmi
```
* Filename is used in 'title' tag in HTML.
Note: '-t' flag may be used to override this and title for each file is explicitly asked.

* Also, if the case of title (based on filename) needs to be changed. '-c' flag followed by the case (title, upper or lower). may be used.
Note: '-c' and '-t' flags are mutually exclusive.

```
meow -c title ./src/index.gmi
```

* You can also refer --help option for all the flags.

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

## Note

* This project was made to create my static website.
* There is no support for reading STDIN when file is absent.
* Extremely simple parsing is used, I promote editing 'helpers.go' file if you need any additional feature.
* Unexpected errors might occur if a file already exists. So, I recommend deleting/moving any files you generated; before re-generating them in the same directory after editing.

