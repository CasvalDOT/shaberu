# Dialects

Dialects simple implements an interface with 2 methods:
- `Read`for reading file contents
- `Write`for write the converted data

If you need to create your own dialect just copy the content 
of the template above, and replace the `customProvider`string with the name you want.


```go
package dialects

// NOTE!!
// This is a template. For use it, just
// replace the string "customProvider" with the name
// you want. 

import (
	"encoding/csv"
	"os"
)

type customProvider struct {
	dialect
}

func (d *customProvider) Write(path *string, content map[string]string) error {
	// The write logics here
	return nil
}

func (d *customProvider) Read(path *string, separator *string) (map[string]string, error) {
	// The read logics here
    return nil,nil
}

func NewCustomProvider() IDialect {
	return &customProvider{
		dialect: dialect{},
	}
}
```


Also make sure to edit the file `dialect.go` to handle the new dialect created.

```go
package dialect
// ... the contents of the files
func New(format string) IDialect {
	switch format {
	// ... Default case statements ...
	case "<provider>":
	 return NewCustomProvider()
	default:
		return nil
	}
}

```
