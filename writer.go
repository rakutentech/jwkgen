package main

import (
	"fmt"
	"io"
	"os"
)

// HeaderWriter is an io.Writer that writes a header before all data
type HeaderWriter struct {
	Base          io.Writer
	Header        []byte
	headerWritten bool
}

// Write writes the specified data (and inserts the header if it has not been
// written yet)
func (w *HeaderWriter) Write(data []byte) (n int, err error) {
	if !w.headerWritten {
		if _, err := w.Base.Write(w.Header); err != nil {
			return 0, err
		}
		w.headerWritten = true
	}

	return w.Base.Write(data)
}

// ObjectInfo is the object info including header and suffix
type ObjectInfo struct {
	Suffix string
	Header string
}

func writerFor(objInfo ObjectInfo) (io.Writer, error) {
	if *filename != "" {
		return os.Create(*filename + objInfo.Suffix)
	} else if bareOutput {
		return os.Stdout, nil
	} else {
		return &HeaderWriter{
			Base:   os.Stdout,
			Header: []byte(fmt.Sprintf("\n===== %s =====\n", objInfo.Header)),
		}, nil
	}
}
