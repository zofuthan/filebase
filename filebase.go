package filebase

import (
	"errors"
	"os"
	"path"

	"github.com/omeid/filebase/codec"
)

const (
	ObjectPerm os.FileMode = 0640
	BucketPerm os.FileMode = 0750
)

// You should expect the following errors.
// the fault is an error type so you should
// treat them like so.

var (
	ErrorKeyEmpty       = errors.New("Empty Key.")
	ErrorNotObjectKey   = errors.New("Key is a bucket.")
	ErrorObjectNotFound = errors.New("No Such Object.")
	ErrorLocationEmpty  = errors.New("Location Empty.")
)

// Returns a new bucket object, it does not touch
// the underlying filesystem if it already exists.
// The codec is used for Marshling and Unmarshaling Objects.
// Currently there is, codec.YAML, codec.JSON, codec.GOB.
// To add your own. see https://godoc.org/github.com/omeid/filebase/codec.
func New(location string, codec codec.Codec) (*Bucket, error) {
	location, name := path.Split(location)
	b := newBucket(location, name, codec)

	//As most error are likely to happen when the first bucket is retrived
	//it makes sense to return it's error.
	return b, b.Error()
}
