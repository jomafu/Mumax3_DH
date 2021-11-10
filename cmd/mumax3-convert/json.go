package main

import (
	"encoding/json"
	"io"

	"github.com/jomafu/Mumax3_DH/data"
)

func dumpJSON(f *data.Slice, info data.Meta, out io.Writer) {
	w := json.NewEncoder(out)
	w.Encode(f.Tensors())
}
