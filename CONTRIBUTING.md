# Server development

## Main Go packages used
- gin as web  framework  https://github.com/gin-gonic/gin
- logrus for logs
- statik for statics content css, js ... https://github.com/rakyll/statik.
- go-bindata for html templates https://github.com/jteeuwen/go-bindata.

Only one binary; js, css and html templates are integrated into binary through statik and go-bindata packages.  


## Build

first get go-bindata et statik packages

```bash
go get -u github.com/jteeuwen/go-bindata/...
go get github.com/rakyll/statik
```

Generate go files from:
 - statics files with `statik` 
 - html templates files with `go-bindata`. 

```bash
statik -src=./public
go-bindata -o templates.go templates
```

Build 
```bash
go build -o slide-server
```

### Cross compilation

For Linux
```bash
GOOS=linux GOARCH=amd64 go build -o slide-server
```

For Windows
```bash
GOOS=windows GOARCH=amd64 go build -o slide-server
```

For OSX
```bash
GOOS=darwin GOARCH=amd64 go build -o slide-server
```

For Raspberry pi
```bash
GOOS=linux GOARCH=arm GOARM=5 go build -o slide-server
```

### logs
To switch in release mode (Gin)
```bash
GIN_MODE=release slide-server 
```