// This file is automatically generated by qtc from "metrics.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line metrics/metrics.qtpl:1
package metrics

//line metrics/metrics.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line metrics/metrics.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line metrics/metrics.qtpl:2
func (m *Metrics) StreamPrometheus(qw422016 *qt422016.Writer) {
	//line metrics/metrics.qtpl:2
	qw422016.N().S(`# TYPE connections counter`)
	//line metrics/metrics.qtpl:3
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:3
	qw422016.N().S(`connections`)
	//line metrics/metrics.qtpl:4
	qw422016.N().S(` `)
	//line metrics/metrics.qtpl:4
	qw422016.N().D(m.Connections)
	//line metrics/metrics.qtpl:4
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:5
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:5
	qw422016.N().S(`# TYPE timeouts counter`)
	//line metrics/metrics.qtpl:6
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:6
	qw422016.N().S(`timeouts`)
	//line metrics/metrics.qtpl:7
	qw422016.N().S(` `)
	//line metrics/metrics.qtpl:7
	qw422016.N().D(m.Timeouts)
	//line metrics/metrics.qtpl:7
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:8
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:8
	qw422016.N().S(`# TYPE errors counter`)
	//line metrics/metrics.qtpl:9
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:9
	qw422016.N().S(`errors`)
	//line metrics/metrics.qtpl:10
	qw422016.N().S(` `)
	//line metrics/metrics.qtpl:10
	qw422016.N().D(m.Errors)
	//line metrics/metrics.qtpl:10
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:11
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:11
	qw422016.N().S(`# TYPE requestSum counter`)
	//line metrics/metrics.qtpl:12
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:12
	qw422016.N().S(`requestSum`)
	//line metrics/metrics.qtpl:13
	qw422016.N().S(` `)
	//line metrics/metrics.qtpl:13
	qw422016.N().D(m.RequestSum)
	//line metrics/metrics.qtpl:13
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:14
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:14
	qw422016.N().S(`# TYPE requestDuration counter`)
	//line metrics/metrics.qtpl:15
	qw422016.N().S(`
`)
	//line metrics/metrics.qtpl:15
	qw422016.N().S(`requestDuration`)
	//line metrics/metrics.qtpl:16
	qw422016.N().S(` `)
	//line metrics/metrics.qtpl:16
	qw422016.N().D(int(m.RequestDuration))
	//line metrics/metrics.qtpl:16
	qw422016.N().S(`
`)
//line metrics/metrics.qtpl:17
}

//line metrics/metrics.qtpl:17
func (m *Metrics) WritePrometheus(qq422016 qtio422016.Writer) {
	//line metrics/metrics.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line metrics/metrics.qtpl:17
	m.StreamPrometheus(qw422016)
	//line metrics/metrics.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line metrics/metrics.qtpl:17
}

//line metrics/metrics.qtpl:17
func (m *Metrics) Prometheus() string {
	//line metrics/metrics.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line metrics/metrics.qtpl:17
	m.WritePrometheus(qb422016)
	//line metrics/metrics.qtpl:17
	qs422016 := string(qb422016.B)
	//line metrics/metrics.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line metrics/metrics.qtpl:17
	return qs422016
//line metrics/metrics.qtpl:17
}
