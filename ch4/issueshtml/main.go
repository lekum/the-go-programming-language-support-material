package main

import (
	"github.com/lekum/the-go-programming-language-support-material/ch4/github"
	"html/template"
	"log"
	"os"
)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func main() {
	issues, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var report = template.Must(template.New("issuelist").Parse(templ))
	if report.Execute(os.Stdout, issues); err != nil {
		log.Fatal(err)
	}
}
