package renderer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode"
)

// Text is a text formatter enclosing a Stringer.  It encloses the
// Stringer's content and word-wraps it.
// TODO: Add a concept of scrolling / framing.
type Text struct {
	W, H int
	fmt.Stringer
}

func errOrEOF(e error) error {
	if e != nil && e != io.EOF {
		return e
	}
	return io.EOF
}

func writeOmitNewlines(to io.Writer, some []byte) (int, error) {
	totalNewlines := bytes.Count(some, []byte("\n"))
	n, err := to.Write(some)
	return n - totalNewlines, err
}

type textSplitter struct {
	W, H int

	Row, Col int

	token []byte
}

func (t *textSplitter) Split(
	in []byte,
	atEOF bool,
) (advance int, token []byte, err error) {

	// Did we run out of rows?
	if t.Row == t.H {
		// Yes, stop scanning.
		return 0, nil, io.EOF
	}

	// Set advance to be input size; ask for more if needed.
	if advance = len(in); advance == 0 {
		return 0, nil, nil
	}

	// How much space is left?
	rowRemain := t.W - t.Col
	if advance > rowRemain {
		// Not enough, truncate to row.
		advance = rowRemain
	}

	// Let's see if this is our token.
	token = in[:advance]

	// Do we have a newline within the row?
	nextSplit := bytes.IndexByte(token, '\n')
	if nextSplit != -1 {
		// If we found a newline, return truncated token,
		// including the newline, and reset to next row.
		advance = nextSplit + 1
		t.Row, t.Col = t.Row+1, 0
		return advance, token[:advance], err
	}

	// Did we run out of input (without a newline?)
	if advance < rowRemain {
		// Yes, just use what we have.  No newline needed.
		t.Col += advance
		return
	}

	// We're finishing the row, so try not to split on a word.
	if !unicode.IsSpace(rune(token[advance-1])) {
		// Look for the last space in the last 14 bytes.
		from, to := advance-14, advance
		if from < 0 {
			from = 0
		}
		seek := bytes.LastIndexAny(token[from:to], " \t")

		// If space found, split there and truncate the token.
		if seek != -1 {
			advance = from + seek + 1
			// Don't include the space.
			token = token[:advance-1]
		}
	}

	// If there are no remaining rows, skip adding newline.
	if t.Row >= t.H-1 {
		t.Row, t.Col = t.Row+1, 0
		return
	}

	// It's not safe to append since it might overwrite the input
	// buffer, so make a copy.  Try to reuse the internal copy
	// buffer.
	lt := len(token)
	if cap(t.token) < lt+1 {
		t.token = make([]byte, lt, lt+1)
	} else {
		t.token = t.token[:lt]
	}
	copy(t.token, token)
	token = append(t.token, '\n')

	t.Col, t.Row = 0, t.Row+1

	return
}

// WriteTo implements io.WriterTo on Text, splitting on newlines and
// trying to avoid splitting words.
func (t Text) WriteTo(wr io.Writer) (total int64, err error) {
	s := bufio.NewScanner(bytes.NewBufferString(t.String()))
	s.Split((&textSplitter{W: t.W, H: t.H}).Split)

	var n int

	for s.Scan() {
		n, err = wr.Write(s.Bytes())
		total += int64(n)
		if err != nil {
			return
		}
	}

	return total, io.EOF
}
