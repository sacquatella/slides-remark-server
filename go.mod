module github.com/sacquatella/slides-remark-server

go 1.12

require (
	github.com/gin-contrib/logger v0.0.1
	github.com/gin-contrib/static v0.0.0-20190511124741-c1cdf9c9ec7b
	github.com/gin-gonic/gin v1.4.0
	github.com/rakyll/statik v0.1.6
	github.com/rs/zerolog v1.14.3
	github.com/stretchr/testify v1.3.0
	golang.org/x/sys v0.0.0-20190610081024-1e42afee0f76 // indirect

)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
