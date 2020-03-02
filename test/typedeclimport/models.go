package typedeclimport

import subpkg "github.com/bamzi/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
