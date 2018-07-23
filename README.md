# Logger

Logger personalizado para GoLang

## Installation

To install try package, you need to install Go and set your Go workspace first.

1. Download and install it:

```sh
$ go get -u github.com/c019/logger
```

2. Import it in your code:

```go
import "github.com/c019/logger"
```

## API examples:

```go
package main

import (
	"github.com/c019/logger"
)

func init() {
	logger.SetStatusExit(0)            // Caso queria mudar o status do os.Exit
	logger.SetDebug(true)              // Para que os logger.Debug seja mostrado
	logger.SetFileLocal("./teste.log") // Para que ele salve em um arquivo externo
}

func main() {
	logger.Info("Teste")
	logger.Debug("Teste")
	logger.Warn("Teste")
	logger.Error("Teste")
	logger.Panic("Teste") // Para mostrar o logger.Fatal, comente esta linha de código
	logger.Fatal("Teste")

	logger.CloseFile()
}
```