// go-jsonb は JSON ファイルの読み書きを提供します.
package jsonb

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path"
)

// Load は JSON ファイルを読み込みます
func Load[V any](name string, v *V) error {
	r, err := os.Open(name)
	if err != nil {
		return err
	}
	defer r.Close()

	fi, err := r.Stat()
	if err != nil {
		return err
	}

	b := bytes.NewBuffer(make([]byte, 0, fi.Size()))
	if _, err = io.Copy(b, r); err != nil {
		return err
	}
	return json.Unmarshal(b.Bytes(), v)
}

// Save は JSON ファイルに保存します.
func Save[V any](name string, v *V) error {
	if err := os.MkdirAll(path.Dir(name), 0755); err != nil {
		return err
	}
	w, err := os.Create(name)
	if err != nil {
		return err
	}
	defer w.Close()

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
