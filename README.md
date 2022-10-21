# Intrajasa golang library

## Installation
### 1.1 Using go module
Run this command on your project to initialize Go mod (if you haven't):
```go
go mod init
```
then reference intrajasa-go in your project file using `import`:
```go
import (
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/coreapi"
    "github.com/midtrans/midtrans-go/snap"
    "github.com/midtrans/midtrans-go/iris"
)
```

### 1.2 Using go get
Also, the alternative way you can use `go get` the package into your project
```go
go get -u github.com/Fermekoo/intrajasa-go
```