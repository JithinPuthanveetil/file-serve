package config

// Root is the root directory
var Root = ""

// Template is the html template
const Template = `
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<!-- Modified from lighttpd directory listing -->
<head>
<title>Index of {{.Root}}</title>
<style type="text/css">
a, a:active {text-decoration: none; color: blue;}
a:visited {color: #48468F;}
a:hover, a:focus {text-decoration: underline; color: red;}
body {background-color: #F5F5F5;}
h2 {margin-bottom: 12px;}
table {margin-left: 12px;}
th, td { font: 120% monospace; text-align: left;}
th { font-weight: bold; padding-right: 14px; padding-bottom: 3px;}
td {padding-right: 14px;}
td.s, th.s {text-align: right;}
div.list { display: table; background-color: white; border-top: 1px solid #646464; border-bottom: 1px solid #646464; padding-top: 10px; padding-bottom: 14px;}
</style>
</head>
<body>
<h2>Index of {{.Root}}</h2>
<div class="list">
<table summary="Directory Listing" cellpadding="2" cellspacing="1">
<thead><tr><th class="n">Name</th><th class="t">Type</th><th class="dl">Options</th></tr></thead>
<tbody>
<tr><td class="n"><a href="../">Parent Directory</a>/</td><td class="t">Directory</td><td class="dl"></td></tr>
{{range .FolderName}}
<tr><td class="n"><a href="{{.}}/">{{.}}</a></td><td class="t">Directory</td><td class="dl"><a href="{{.}}?download=true">Download</a></td></tr>
{{end}}
{{range .FileName}}
<tr><td class="n"><a href="{{.}}">{{.}}</a></td><td class="t">&nbsp;</td><td class="dl"><a href="{{.}}?download=true">Download</a></td></tr>
{{end}}
</tbody>
</table>
</div>
</body>
</html>`
