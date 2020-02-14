# client-go-client

A client-go implementation of the Manifestival Client 

Usage
-----

```go
import (
	mfc "github.com/manifestival/client-go-client"
	mf  "github.com/manifestival/manifestival"
	"k8s.io/client-go/rest"
)

func main() {
    var config *rest.Config = ...
    
    manifest, err := mfc.NewManifest("dir/", config, mf.Recursive)
	if err != nil {
		panic("Failed to load manifest")
	}
    
    manifest.ApplyAll()
}
```

The `NewManifest` function in this library delegates to the function
of the same name in the `manifestival` package after constructing a
`manifestival.Client` implementation from the `*rest.Config`.
