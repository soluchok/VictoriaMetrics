// Code generated by qtc from "funcs.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmalert/templates/funcs.qtpl:3
package templates

//line app/vmalert/templates/funcs.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmalert/templates/funcs.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmalert/templates/funcs.qtpl:3
func streamquotesEscape(qw422016 *qt422016.Writer, s string) {
//line app/vmalert/templates/funcs.qtpl:4
	qw422016.N().J(s)
//line app/vmalert/templates/funcs.qtpl:5
}

//line app/vmalert/templates/funcs.qtpl:5
func writequotesEscape(qq422016 qtio422016.Writer, s string) {
//line app/vmalert/templates/funcs.qtpl:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/templates/funcs.qtpl:5
	streamquotesEscape(qw422016, s)
//line app/vmalert/templates/funcs.qtpl:5
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/templates/funcs.qtpl:5
}

//line app/vmalert/templates/funcs.qtpl:5
func quotesEscape(s string) string {
//line app/vmalert/templates/funcs.qtpl:5
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/templates/funcs.qtpl:5
	writequotesEscape(qb422016, s)
//line app/vmalert/templates/funcs.qtpl:5
	qs422016 := string(qb422016.B)
//line app/vmalert/templates/funcs.qtpl:5
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/templates/funcs.qtpl:5
	return qs422016
//line app/vmalert/templates/funcs.qtpl:5
}

//line app/vmalert/templates/funcs.qtpl:7
func streamjsonEscape(qw422016 *qt422016.Writer, s string) {
//line app/vmalert/templates/funcs.qtpl:8
	qw422016.N().Q(s)
//line app/vmalert/templates/funcs.qtpl:9
}

//line app/vmalert/templates/funcs.qtpl:9
func writejsonEscape(qq422016 qtio422016.Writer, s string) {
//line app/vmalert/templates/funcs.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/templates/funcs.qtpl:9
	streamjsonEscape(qw422016, s)
//line app/vmalert/templates/funcs.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/templates/funcs.qtpl:9
}

//line app/vmalert/templates/funcs.qtpl:9
func jsonEscape(s string) string {
//line app/vmalert/templates/funcs.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/templates/funcs.qtpl:9
	writejsonEscape(qb422016, s)
//line app/vmalert/templates/funcs.qtpl:9
	qs422016 := string(qb422016.B)
//line app/vmalert/templates/funcs.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/templates/funcs.qtpl:9
	return qs422016
//line app/vmalert/templates/funcs.qtpl:9
}

//line app/vmalert/templates/funcs.qtpl:11
func streamhtmlEscape(qw422016 *qt422016.Writer, s string) {
//line app/vmalert/templates/funcs.qtpl:12
	qw422016.E().S(s)
//line app/vmalert/templates/funcs.qtpl:13
}

//line app/vmalert/templates/funcs.qtpl:13
func writehtmlEscape(qq422016 qtio422016.Writer, s string) {
//line app/vmalert/templates/funcs.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/templates/funcs.qtpl:13
	streamhtmlEscape(qw422016, s)
//line app/vmalert/templates/funcs.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/templates/funcs.qtpl:13
}

//line app/vmalert/templates/funcs.qtpl:13
func htmlEscape(s string) string {
//line app/vmalert/templates/funcs.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/templates/funcs.qtpl:13
	writehtmlEscape(qb422016, s)
//line app/vmalert/templates/funcs.qtpl:13
	qs422016 := string(qb422016.B)
//line app/vmalert/templates/funcs.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/templates/funcs.qtpl:13
	return qs422016
//line app/vmalert/templates/funcs.qtpl:13
}
