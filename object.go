package filebase

import (
	"os"
	"path"
	"sync"
	"syscall"

	"github.com/omeid/filebase/codec"
)

type object struct {
	sync.RWMutex

	location string
	key      string
	unique   bool
	perm     os.FileMode
}

func (o *object) Write(codec codec.Codec, data interface{}, sync bool) error {

	if o.key == "" {
		return ErrorKeyEmpty
	}

	o.Lock()
	defer o.Unlock()

	mode := os.O_WRONLY | os.O_CREATE | os.O_TRUNC // overwrite if exists
	if o.unique {
		mode = mode | os.O_EXCL
	}

	file, err := os.OpenFile(path.Join(o.location, o.key), mode, o.perm)
	if err != nil {
		return err
	}

	defer file.Close()

	//Get an exclusive lock on the file.
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	if err != nil {
		return err
	}
	defer func() {
		err = syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	}()

	if sync {
		defer func() {
			//Don't sync if we have encournted an error.
			if err == nil {
				err = file.Sync()
			}
		}()
	}

	return codec.NewEncoder(file).Encode(data)
}

//Read an object from file.
func (o *object) Read(codec codec.Codec, out interface{}) error {

	if o.key == "" {
		return ErrorKeyEmpty
	}

	o.RLock()
	defer o.RUnlock()

	file, err := os.Open(path.Join(o.location, o.key))

	if os.IsNotExist(err) {
		return ErrorObjectNotFound
	}

	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
	}()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return ErrorNotObjectKey
	}

	return codec.NewDecoder(file).Decode(out)
}

func (o *object) Drop() error {

	o.Lock()
	defer o.Unlock()

	return os.Remove(path.Join(o.location, o.key))
}
