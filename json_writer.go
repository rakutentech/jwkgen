package main

import (
	"encoding/json"
	"github.com/nwidger/jsoncolor"
	"io"
)

func prettyMarshal(obj interface{}) ([]byte, error) {
	return jsoncolor.MarshalIndent(obj, "", "  ")
}

func writeJSON(w io.Writer, obj interface{}) error {
	if *useColor && *filename == "" {
		if data, err := prettyMarshal(obj); err != nil {
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
