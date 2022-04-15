package storage

import (
	"encoding/json"
	"testing"

	"github.com/rafaelbreno/nuveo-test/config"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFileStorage(t *testing.T) {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level.SetLevel(zap.FatalLevel)
	l, _ := cfg.Build()

	type (
		Foo struct {
			Bar string `json:"bar"`
		}
	)

	{
		in := &internal.Internal{
			Cfg: &config.Config{
				Service: config.Service{
					NewClients: "tmp/",
				},
			},
			L: l,
		}

		t.Run("NewFile", func(t *testing.T) {
			assert := assert.New(t)
			_, err := NewFile(in)
			assert.Nil(err)

			in.Cfg.Service.NewClients = ""

			_, err = NewFile(in)
			assert.NotNil(err)
		})
	}

	foo := Foo{
		Bar: "Bar",
	}

	key := "foo"
	b, err := json.Marshal(foo)

	if err != nil {
		t.Fatal(err)
	}

	{
		in := &internal.Internal{
			Cfg: &config.Config{
				Service: config.Service{
					NewClients: "tmp/",
				},
			},
			L: l,
		}

		f, err := NewFile(in)

		if err != nil {
			t.Fatal(err)
		}

		t.Run("File.Write", func(t *testing.T) {
			assert := assert.New(t)

			err = f.Write(key, string(b))
			assert.Nil(err)

			f.in.Cfg.Service.NewClients = "$%#"

			err = f.Write(key, string(b))
			assert.NotNil(err)
		})
	}

	{
		in := &internal.Internal{
			Cfg: &config.Config{
				Service: config.Service{
					NewClients: "tmp/",
				},
			},
			L: l,
		}

		f, err := NewFile(in)

		if err != nil {
			t.Fatal(err)
		}

		t.Run("File.Read", func(t *testing.T) {
			assert := assert.New(t)

			gotStr, err := f.Read(key)
			assert.Nil(err)
			assert.Equal(string(b), gotStr)

			_, err = f.Read("random_key")
			assert.NotNil(err)
		})
	}

	{
		in := &internal.Internal{
			Cfg: &config.Config{
				Service: config.Service{
					NewClients: "tmp/",
				},
			},
			L: l,
		}

		f, err := NewFile(in)

		if err != nil {
			t.Fatal(err)
		}

		t.Run("File.Delete", func(t *testing.T) {
			assert := assert.New(t)

			err := f.Delete(key)
			assert.Nil(err)

			err = f.Delete("random_key")
			assert.NotNil(err)
		})
	}

	{
		in := &internal.Internal{
			Cfg: &config.Config{
				Service: config.Service{
					NewClients: "tmp/",
				},
			},
			L: l,
		}

		f, err := NewFile(in)

		if err != nil {
			t.Fatal(err)
		}

		t.Run("File.MarshalAndWrite", func(t *testing.T) {
			assert := assert.New(t)

			err = f.MarshalAndWrite(foo, key)
			assert.Nil(err)

			f.in.Cfg.Service.NewClients = "$%#"

			err = f.MarshalAndWrite(foo, key)
			assert.NotNil(err)
		})
	}

	{
		in := &internal.Internal{
			Cfg: &config.Config{
				Service: config.Service{
					NewClients: "tmp/",
				},
			},
			L: l,
		}

		f, err := NewFile(in)

		if err != nil {
			t.Fatal(err)
		}

		t.Run("File.ReadAndUnmarshal", func(t *testing.T) {
			assert := assert.New(t)

			fooVar := Foo{}

			err := f.ReadAndUnmarshal(&fooVar, key)
			assert.Nil(err)
			assert.Equal(fooVar, foo)

			err = f.ReadAndUnmarshal(fooVar, key)
			assert.NotNil(err)

			err = f.ReadAndUnmarshal(&fooVar, "random_key")
			assert.NotNil(err)
		})
	}

	{
		in := &internal.Internal{
			Cfg: &config.Config{
				Service: config.Service{
					NewClients: "tmp/",
				},
			},
			L: l,
		}

		f, err := NewFile(in)

		if err != nil {
			t.Fatal(err)
		}

		t.Run("File.getFilePath", func(t *testing.T) {
			assert := assert.New(t)

			expected := in.Cfg.Service.NewClients + "/" + key + ".json"

			got := f.getFilePath(key)

			assert.Equal(expected, got)
		})
	}
}
