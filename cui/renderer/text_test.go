package renderer_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/suru.v0/cui/renderer"
	"gopkg.in/suru.v0/cui/view"
)

var _ = bufio.SplitFunc(renderer.ExportTextSplitter().Split)

type tCase struct {
	w, h  int
	cases []withText
}
type withText struct {
	given, expect string
	expectn       int64
}

func TestText(t *testing.T) {
	t.Helper()
	for _, wh := range []tCase{{
		72, 100,
		[]withText{{
			"", "", 0,
		}, {
			"hello", "hello",
			5,
		}, {
			"\ttab\ttest",
			"\ttab\ttest",
			9,
		}, {
			`
hello
new
lines`[1:],
			`
hello
new
lines`[1:],
			15,
		}, {
			strings.Repeat("hello", 20),
			strings.Repeat("hello", 20)[:72] + "\n" +
				strings.Repeat("llohe", 10)[:28],
			101,
		}},
	}, {
		20, 5,
		[]withText{{
			"hello", "hello",
			5,
		}, {
			strings.Repeat("hello ", 10),
			strings.Join([]string{
				"hello hello hello",
				"hello hello hello",
				"hello hello hello",
				"hello ",
			}, "\n"),
			60,
		}, {
			strings.Repeat("hello ", 20),
			`
hello hello hello
hello hello hello
hello hello hello
hello hello hello
hello hello hello`[1:],
			89,
		}},
	}, {
		40, 40,
		[]withText{{
			"hello", "hello",
			5,
		}, {
			strings.Repeat("lorum ipsum blah blah blahhhhhhh", 10) + `
A rather long segment of text with `[1:] + `
several spaces and `[1:] + `
repeats and things like that `[1:] + `
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`[1:],
			`
lorum ipsum blah blah blahhhhhhhlorum
ipsum blah blah blahhhhhhhlorum ipsum
blah blah blahhhhhhhlorum ipsum blah
blah blahhhhhhhlorum ipsum blah blah
blahhhhhhhlorum ipsum blah blah
blahhhhhhhlorum ipsum blah blah
blahhhhhhhlorum ipsum blah blah
blahhhhhhhlorum ipsum blah blah
blahhhhhhhlorum ipsum blah blah
blahhhhhhhA rather long segment of text 
with several spaces and repeats and
things like that aaaaaaaaaaaaaaaaaaaaaaa
aaaaaaaaaaaaaaaa`[1:],
			444,
		}},
	}, {
		40, 12,
		[]withText{{
			"hello", "hello",
			5,
		}, {
			strings.Repeat("hello ", 100),
			strings.Repeat(`
hello hello hello hello hello hello
`[1:], 11) + "hello hello hello hello hello hello",
			431,
		}},
	}} {
		t.Run(fmt.Sprintf("Size %dx%d", wh.w, wh.h), wh.test)
	}
}

func (tt tCase) test(t *testing.T) {
	for i, tc := range tt.cases {
		logExp := tc.given
		if ltc := len(logExp); ltc > 20 {
			logExp = string(append(
				[]byte(logExp[:17]),
				[]byte("...")...,
			))
		}
		t.Logf("case %d: %s", i, logExp)

		given := renderer.Text{tt.w, tt.h, view.Text(tc.given)}
		buf := new(bytes.Buffer)
		n, err := given.WriteTo(buf)
		assert.Equal(t, io.EOF, err, "error should always be io.EOF")
		assert.Equal(t, tc.expectn, n, "Length should match")
		assert.Equal(t, tc.expect, buf.String(), "contents should match")
	}
}
