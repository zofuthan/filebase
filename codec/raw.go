package codec

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
)

type RAW struct{}

func (y RAW) NewDecoder(r io.Reader) Decoder {
	return raw_codec{r: r}
}

func (y RAW) NewEncoder(w io.Writer) Encoder {
	return raw_codec{w: w}
}

type raw_codec struct {
	r io.Reader
	w io.Writer
}

func (y raw_codec) Decode(v interface{}) error {

	b, ok := v.(*[]byte)

	if !ok {
		return errors.New("Raw codec only accept *[]byte")
	}

	d, err := ioutil.ReadAll(y.r)

	*b = d
	return err
}

func (y raw_codec) Encode(v interface{}) error {

	b, ok := v.([]byte)

	if !ok {
		return errors.New("Raw codec only accept []byte")
	}

	_, err := bytes.NewBuffer(b).WriteTo(y.w)
	return err
}
