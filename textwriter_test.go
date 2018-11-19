package textwriter

import (
	"bytes"
	"log"
	"testing"
)

func TestWriter(t *testing.T) {
	tests := []struct {
		in   []string
		want string
	}{
		{
			in:   []string{"foo\nbar\n"},
			want: "foo\r\nbar\r\n",
		},
		{
			in:   []string{"foo\r\nbar\r\n"},
			want: "foo\r\nbar\r\n",
		},
		{
			in:   []string{"foo\nbar\r"},
			want: "foo\r\nbar\r",
		},
		{
			in:   []string{"foo\nbar"},
			want: "foo\r\nbar",
		},
		{
			in:   []string{""},
			want: "",
		},
		{
			in:   []string{"\n"},
			want: "\r\n",
		},
		{
			in:   []string{"\n\n"},
			want: "\r\n\r\n",
		},
		{
			in:   []string{"foo\r", "\nbar\r\n"},
			want: "foo\r\nbar\r\n",
		},
		{
			in:   []string{"foo", "\r\nbar\r\n"},
			want: "foo\r\nbar\r\n",
		},
		{
			in:   []string{"foo\r\nbar\r", "\n"},
			want: "foo\r\nbar\r\n",
		},
	}
	for _, test := range tests {
		var buf bytes.Buffer
		w := NewWriter(&buf)
		for _, in := range test.in {
			_, err := w.Write([]byte(in))
			if err != nil {
				log.Fatal(err)
			}
		}
		got := buf.String()
		if got != test.want {
			t.Fatalf("want %q for %q but got %q", test.want, test.in, got)
		}
	}
}
