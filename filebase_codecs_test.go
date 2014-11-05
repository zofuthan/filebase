package filebase

import (
	"encoding/gob"
	"reflect"
	"testing"

	"github.com/omeid/filebase/codec"
)

type TestObject struct {
	Hello string
	Tag   []string
	Key   string
}

func init() {
}

var (
	o = TestObject{"World.",
		[]string{
			"This",
			"is",
			"JSON.",
		},
		"",
	}

	TestKeys = []string{"key1", "key with space", "key-1", "0key", "test"}

	TestQuerys = map[string][]string{
		//May need sorting for this, the order of return depends on system!
		"*":     []string{"key1", "key with space", "key-1", "0key", "test"},
		"key?":  []string{"key1"},
		"?key*": []string{"0key"},
		"k*":    []string{"key1", "key with space", "key-1"},
		"test":  []string{"test"},
	}
)

func TestWrite(t *testing.T) {

	gob.Register(TestObject{})

	for _, codec := range []codec.Codec{codec.JSON{}, codec.YAML{}, codec.GOB{}} {
		fb := New("test-db", codec)
		codec_name := reflect.TypeOf(codec).Name()

		for _, key := range TestKeys {

			o.Key = key

			fb.Collection(codec_name).Put(key, o, false, false)
			r := TestObject{}
			fb.Collection(codec_name).Get(key, &r)

			if !reflect.DeepEqual(o, r) {
				t.Fatal(o, r)
			}
		}

		for query, expected := range TestQuerys {
			keys, err := fb.Collection(codec_name).Query(query)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(keys, expected) {
				t.Fatalf("\nCodec:   %s\n\nQuery:    [%+v]\nExpected: %+v, \nGot:      %+v", codec_name, query, expected, keys)
			}
		}
	}
}