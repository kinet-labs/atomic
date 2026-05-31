package atomic

//go:generate bin/gen-atomicwrapper -name=String -type=string -wrapped Value -pack packString -unpack unpackString -compareandswap -swap -file=string.go

func packString(s string) interface{} {
	return s
}

func unpackString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// String returns the wrapped value.
func (s *String) String() string {
	return s.Load()
}

// MarshalText encodes the wrapped string into a textual form.
//
// This makes it encodable as JSON, YAML, XML, and more.
func (s *String) MarshalText() ([]byte, error) {
	return []byte(s.Load()), nil
}

// UnmarshalText decodes text and replaces the wrapped string with it.
//
// This makes it decodable from JSON, YAML, XML, and more.
func (s *String) UnmarshalText(b []byte) error {
	s.Store(string(b))
	return nil
}
