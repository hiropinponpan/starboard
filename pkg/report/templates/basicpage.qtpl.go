// Code generated by qtc from "basicpage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line pkg/report/templates/basicpage.qtpl:3
package templates

//line pkg/report/templates/basicpage.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pkg/report/templates/basicpage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pkg/report/templates/basicpage.qtpl:4
type Page interface {
//line pkg/report/templates/basicpage.qtpl:4
	Title() string
//line pkg/report/templates/basicpage.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line pkg/report/templates/basicpage.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line pkg/report/templates/basicpage.qtpl:4
	Body() string
//line pkg/report/templates/basicpage.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line pkg/report/templates/basicpage.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line pkg/report/templates/basicpage.qtpl:4
}

// Page prints a page implementing Page interface.

//line pkg/report/templates/basicpage.qtpl:12
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line pkg/report/templates/basicpage.qtpl:12
	qw422016.N().S(`
<html>
	<head>
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css" integrity="sha384-9aIt2nRpC12Uk9gS9baDl411NQApFmC26EwAOH8WgZl5MYYxFfc+NcPb1dKGj7Sk" crossorigin="anonymous">
		<link href="https://fonts.googleapis.com/css2?family=Lato:wght@300;700&display=swap" rel="stylesheet">
		<title>`)
//line pkg/report/templates/basicpage.qtpl:17
	p.StreamTitle(qw422016)
//line pkg/report/templates/basicpage.qtpl:17
	qw422016.N().S(`</title>
	</head>
	<body style="font-family: 'Lato', sans-serif;">
		`)
//line pkg/report/templates/basicpage.qtpl:20
	p.StreamBody(qw422016)
//line pkg/report/templates/basicpage.qtpl:20
	qw422016.N().S(`
	</body>
</html>
`)
//line pkg/report/templates/basicpage.qtpl:23
}

//line pkg/report/templates/basicpage.qtpl:23
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line pkg/report/templates/basicpage.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/basicpage.qtpl:23
	StreamPageTemplate(qw422016, p)
//line pkg/report/templates/basicpage.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/basicpage.qtpl:23
}

//line pkg/report/templates/basicpage.qtpl:23
func PageTemplate(p Page) string {
//line pkg/report/templates/basicpage.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/basicpage.qtpl:23
	WritePageTemplate(qb422016, p)
//line pkg/report/templates/basicpage.qtpl:23
	qs422016 := string(qb422016.B)
//line pkg/report/templates/basicpage.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/basicpage.qtpl:23
	return qs422016
//line pkg/report/templates/basicpage.qtpl:23
}
