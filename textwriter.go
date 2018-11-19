package textwriter

import (
	"io"

	"golang.org/x/text/transform"
)

type text struct {
	prev byte
}

func (t *text) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nDst < len(dst) && nSrc < len(src) {
		c := src[nSrc]
		if c == '\n' {
			if nDst+1 >= len(dst) {
				break
			}
			if t.prev != '\r' {
				dst[nDst] = '\r'
				nDst++
			}
			dst[nDst] = '\n'
			nDst++
			nSrc++
		} else {
			dst[nDst] = c
			nDst++
			nSrc++
		}
		t.prev = c
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return
}

func (t *text) Reset() {
	t.prev = 0
}

func NewWriter(w io.Writer) io.WriteCloser {
	return transform.NewWriter(w, &text{})
}
