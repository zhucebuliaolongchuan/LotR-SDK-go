## Introduction
Do you love The Lord of the Rings? Great, this is a powerful go sdk to interact with an existing Lord of the Rings [API](https://the-one-api.dev/documentation)

### Dependencies
* install [Golang](https://go.dev/doc/install) (>=1.15)

### How to get
```
$ go get github.com/zhucebuliaolongchuan/LotR-SDK-go
```

### Usage
#### Build
```
$ ./LotR-SDK-go --help
Do you love The Lord of the Rings? Great, this is a helpful CLI for you to consume information about the trilogy

Usage:
  LotR-SDK-go [command]

Available Commands:
  book        List books
  chapter     List book chapters
  character   List characters including metadata like name, gender, realm, race and more
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  movie       List movies
  quote       List movie quotes or one specific movie quote

Flags:
  -h, --help   help for LotR-SDK-go

Use "LotR-SDK-go [command] --help" for more information about a command.
```

#### TODO
* Provide feature flags for pagination, sorting and filtering
