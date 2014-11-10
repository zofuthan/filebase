package codec

import (
	"io"

	"github.com/omeid/filebase/codec/ini"
)

type INI struct{}

func (i INI) NewDecoder(r io.Reader) Decoder {
	return ini_codec{r: r}
}

func (i INI) NewEncoder(w io.Writer) Encoder {
	return ini_codec{w: w}
}

type ini_codec struct {
	r io.Reader
	w io.Writer
}

func (i ini_codec) Decode(v interface{}) error {

	return ini.Decoder(i.r, v)
}

func (i ini_codec) Encode(v interface{}) error {
	return ini.Encoder(v, i.w)
}
