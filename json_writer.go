package main

import (
	"encoding/json"
	"io"

	"github.com/hokaccha/go-prettyjson"
)

func writeJSON(w io.Writer, obj interface{}) error {
	if *color && *filename == "" {
		if data, err := prettyjson.Marshal(obj); err != nil {
			return err
		} else if _, err := w.Write(data); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		return e.Encode(obj)
	}
}

func writeJSONFor(objInfo ObjectInfo, obj interface{}) error {
	w, err := writerFor(objInfo)
	if err != nil {
		return err
	}
	return writeJSON(w, obj)
}
