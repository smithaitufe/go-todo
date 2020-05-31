//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
package schema

import (
	"bytes"
	"log"
)

func GetRootSchema() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b, err := Asset(name)
		if err != nil {
			log.Fatalf("Error occurred %v\n", err)
		}
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
