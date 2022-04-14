package storage

type (
	// Storage is the interface
	// that wraps all methods
	// related to storaging.
	Storage interface {
		Write(key, str string) error
		Read(key string) (string, error)
		Delete(key string) error
		MarshalAndWrite(v interface{}, key string) error
		ReadAndUnmarshal(v interface{}, key string) error
	}
)
