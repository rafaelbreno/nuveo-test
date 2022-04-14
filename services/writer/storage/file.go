package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/rafaelbreno/nuveo-test/internal"
)

type (
	// File is responsible for
	File struct {
		in *internal.Internal
	}
)

// NewFile creates an instance
// of File struct with a given
// Internal.
func NewFile(in *internal.Internal) File {
	return File{
		in: in,
	}
}

// Write will receive a given string
// and write it in a file.
func (f *File) Write(key, str string) error {
	file, err := os.Create(key + ".json")

	if err != nil {
		f.in.L.Error(err.Error())
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			f.in.L.Error(err.Error())
		}
	}()

	if _, err := file.WriteString(str); err != nil {
		f.in.L.Error(err.Error())
		return err
	}

	return nil
}

// Read will search for a file with
// given key, and return its content.
func (f *File) Read(key string) (string, error) {
	b, err := os.ReadFile(key + ".json")
	if err != nil {
		f.in.L.Error(err.Error())
		return "", err
	}
	return string(b), nil
}

// Delete will search for a file with
// given key, and delete it.
func (f *File) Delete(key string) error {
	return os.Remove(key + ".json")
}

// MarshalAndWrite is a shortcut between
// json.Marshal and file.Write.
func (f *File) MarshalAndWrite(v interface{}, key string) error {
	b, err := json.Marshal(v)

	if err != nil {
		f.in.L.Error(err.Error())
		return err
	}

	return f.Write(key, string(b))
}

// ReadAndUnmarshal is a shortcut between
// file.Read and json.Unmarshal.
func (f *File) ReadAndUnmarshal(v interface{}, key string) error {
	// check if given "v" is a pointer
	kindOfV := reflect.ValueOf(v).Kind()
	if kindOfV != reflect.Ptr {
		err := fmt.Errorf("%s isn't a pointer", "v")
		f.in.L.Error(err.Error())
		return err
	}

	str, err := f.Read(key)
	if err != nil {
		f.in.L.Error(err.Error())
		return err
	}

	return json.Unmarshal([]byte(str), v)
}
