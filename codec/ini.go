package codec

import (
	"io"

	"github.com/omeid/go-ini"
)

type INI struct{}

func (i INI) NewDecoder(r io.Reader) Decoder {
	return ini.NewDecoder(r)
}

func (i INI) NewEncoder(w io.Writer) Encoder {
	return ini.NewEncoder(w)
}
