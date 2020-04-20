## Step by Step using Go Module

1. Set env variable `GO111MODULE="on"`
2. Run `go mod init` on root package
3. If you're using IntelliJ IDEA, after running command in #2 you will get popup notification 'Go Modules are detected', choose 'Enable Integration'
4. Change module name accordingly (optional)
5. Create main.go file
6. Try adding other pre-existed go-module from github repo, for example I use gin module.
7. In main.go file, type in the code like usual, add the gin module path to the import statement

The current import statement will look like this:
```go
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
```

8. Run `go build main.go` on terminal

It will download and extract all the dependencies if they're not already existed in the 'global cache' (that's what the guy on youtube said) the location of global cache is at $GOPATH/pkg/mod
  
9. Run `go run main.go` to start the application