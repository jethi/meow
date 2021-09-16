package main

const help = `
		meow - convert gemini markup to html

USAGE:
		meow FILE/DIRECTORY
		meow [-o OUTPUT_DIR] [-t TEMPLATE_FILE_PATH] FILE/DIRECTORY
		meow [-o OUTPUT_DIR] [-c title | lower | upper | custom] FILE/DIRECTORY
		meow [-o OUTPUT_DIR] [-t TEMPLATE_FILE_PATH] [-c title | lower | upper | custom] FILE/DIRECTORY

NOTE:
		meow needs a template and optionally css to work
		use --generate option to generate the default template and
		css. Like this-

			meow --generate

OPTIONS:

		--help
			print help page (this page)

		--generate
			generate config (default template and css)

		-o, --output
			set output directory

		-t, --template
			specify the template file

			Default: if '-t' is not used meow looks for 'html_template.tmpl' file in current directory

			Note: User is trusted to provide a valid template file.

		-c, --setcase [title | upper | lower | custom | none]
			Converts filename to specified case for title tag
				'custom' option prompts user to enter title for each file

			Default: if '-c' is not used, it defaults to 'none' which uses file name as title

`

const templateFile = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="./style.css">
	<link rel="shortcut icon" type="image/png" href="./assets/favicon.png"/>
{{.Title}}
</head>
<body>
	<nav id="navbar">
		<a class="button" href="/">Home</a>
		<a class="button" href="./blog.html">Blog</a>
		<a class="button" href="./doc.html">Doc</a>
		<a class="button" href="./about.html">About</a>
	</nav>
	<div id="data">
{{.Data}}
	</div>
	<footer>
		<p id="bottomtext">
			{{.Footer}}
		</p>
	</footer>
</body>
</html>
`

const cssFile = `:root {
	--body-bg:#ECECEC;
	--nav-foot-bg:#C5DEDD;
	--button-bg:#FFF1E6;
	--button-hover:#B7B7E1;
	--block-pre-bg:#F1ECCE;
	--html-bg:#F2F2F2;
	--border-col:#3B3B46;
	--all-text:#3B3B46;
}

html {
	padding: 0px;
	background: var(--html-bg);
}

body {
	margin: 0 auto;
	outline: solid 1px var(--border-col);
	max-width: 700px;
	color: var(--all-text);
	background-color: var(--body-bg);
}

#navbar {
	padding: 7px 7px;
	margin-bottom: 17px;
	background-color: var(--nav-foot-bg);
}

.button {
	padding: 2px 10px;
	border: solid 1px var(--border-col);
	border-radius: 10px;
	text-decoration: none;
	display: inline-block;
	position: relative;
	cursor: pointer;
	font-family: serif;
	font-size: 1.25rem;
	color: var(--all-text);
	background-color: var(--button-bg);
}

h1, h2, h3 {
	margin: 0;
	font-family: serif;
}

p, ul, a {
	margin: 0;
	font-family: sans-serif;
}

a {
	text-decoration: none;
	margin: 0px 20px;
	color: var(--all-text);
	font-weight: bold;
}

a:hover {
	background-color: var(--button-hover)
}

#navbar a {
	margin: 0px;
}

#navbar a:hover {
	background-color: var(--button-hover);
}

blockquote {
	padding: 5px 10px;
	margin: 0px 30px;
	font-family: monospace;
	border: dashed 1px var(--border-col);
	background-color: var(--block-pre-bg);
}

pre {
	padding: 5px 10px;
	margin: 0px;
	font-family: monospace;
	background-color: var(--block-pre-bg);
	border: solid 0.5px;
	overflow-x: scroll;
}

#data {
	padding: 0px 10px;
}

footer {
	margin-top: 17px;
	background-color: var(--nav-foot-bg);
}

hr {
	margin: 11px 0px 0px;
}

#bottomtext {
	margin: 0px;
	text-align: center;
	padding: 10px;
}
`
