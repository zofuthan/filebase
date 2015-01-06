> NOTE: This is a prerelease. The API may change.

# Filebase [![wercker status](https://app.wercker.com/status/6438ed03b8e2d1655bef928ba1fe88fc/s "wercker status")](https://app.wercker.com/project/bykey/6438ed03b8e2d1655bef928ba1fe88fc) [![GoDoc](https://godoc.org/github.com/omeid/filebase?status.svg)](https://godoc.org/github.com/omeid/filebase) [![Build Status](https://drone.io/github.com/omeid/filebase/status.png)](https://drone.io/github.com/omeid/filebase/latest)

Version v0.1.0-alpha-4

Filebase is a filesystem based Key-Object store with pluggable codec.



### Why?

Filebase is ideal when you want more than config files yet not a database. Because Filebase is using a filesystem and optionally a human readable encoding, you can work with the database with traditional text editing tools like a GUI text editor or command-line tools like `cat`, `grep`, `head`, et al.

The [gob](http://golang.org/pkg/encoding/gob/) codec makes it possible to store any type of Go object and recover it later. This is ideal for storing objects that have state to reload after restarting your application.

Filebase can also be used as a filesystem abstraction, the RAW ([]byte passthrough) codec is idea for this.

### Codecs

Filebase currently ships YAML, JSON, gob, and RAW codecs.

To build a new codec, you just need to satisify the `codec.Codec` interface:


```go
type Codec interface {
	NewDecoder(io.Reader) Decoder
	NewEncoder(io.Writer) Encoder
}

type Decoder interface {
	Decode(v interface{}) error
}

type Encoder interface {
	Encode(v interface{}) error
}
```

> NOTE: You can use type casting to enforce a specific type of objects. See [Raw Codec](codec/raw.go) for example, which only accepts `[]byte`.

### Buckets & Objects

Filebase has no concept of table or database, it is buckets and objects. A bucket may have any number of objects and buckets to the limits supported by the underlying file system.


### Exampe 

```go

    // Open a bucket. Will create if doesn't exists.
	bucket, err := Open("filebase/path/and-name", codec.RAW{})

	if err != nil {
		t.Fatal(err)
	}

	data := []byte(`Hello world. This is some raw data`)

    // Put our data into the bucket.
    // This will create a file `filebase/path/and-name/test` 
    // containing the `data` value.
	err = bucket.Put("test", expected, true, true)
	if err != nil {
		log.Fatal(err)
	}

    // Grab our data by key and put it in fromstorage variable.
	var fromstorage []byte
	err = bucket.Get("test", &fromstorage)
	if err != nil {
		log.Fatal(err)
	}

    // Drop the object from store. This will also delete the underlying file.
    err = bucket.Drop("test")
    if err != nil {
      log.Fatal(err)
    }

   //Delete our bucket.
	err = bucket.Destroy(false) //Expecting empty bucket.
	if err != nil {
		log.Fatal(err)
	}

```
Please see the [API Documentation](https://godoc.org/github.com/omeid/filebase) for querying and more. Refer to [test](filebase_test.go) for some examples.



### Contribution

Pull requests are welcome.


### TODO:

 - Finish [todo example](examples/todo)
 - Advisory Codec File for buckets.
 - More test for Bucket.Query
