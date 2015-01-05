> NOTE: This is a prerelease. The API may change.

# Filebase [![wercker status](https://app.wercker.com/status/6438ed03b8e2d1655bef928ba1fe88fc/s "wercker status")](https://app.wercker.com/project/bykey/6438ed03b8e2d1655bef928ba1fe88fc) [![GoDoc](https://godoc.org/github.com/omeid/filebase?status.svg)](https://godoc.org/github.com/omeid/filebase) [![Build Status](https://drone.io/github.com/omeid/filebase/status.png)](https://drone.io/github.com/omeid/filebase/latest)

Version v0.1.0-alpha 

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
	NewDecoder(io.Reader) decoder
	NewEncoder(io.Writer) encoder
}

type decoder interface {
	Decode(v interface{}) error
}

type encoder interface {
	Encode(v interface{}) error
}
```

> NOTE: You can use type casting to enforce a specific type of objects. See [Raw Codec](codec/raw.go) for example, which only accepts `[]byte`.

### Buckets & Objects

Filebase has no concept of table or database, it is buckets and objects. A bucket may have any number of objects and buckets to the limits supported by the underlying file system.


Please see the [API Documentation](https://godoc.org/github.com/omeid/filebase) for more details and refer to [test](filebase_test.go) for some example.



### Contribution

Pull requests are welcome.


### TODO:

 - Finish [todo example](examples/todo)
 - Advisory Codec File for buckets.
 - More test for Bucket.Query
