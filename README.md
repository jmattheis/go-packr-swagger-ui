# Go Packr Swagger UI
This is just a go file containing the dist folder of the [swagger-api/swagger-ui](https://github.com/swagger-api/swagger-ui).

## Usage
``` go
package main

import "github.com/jmattheis/go-packr-swagger-ui"

func main() {
   box := swaggerui.GetBox()
}
```

## Usage with [gin-gonic/gin](https://github.com/gin-gonic/gin)
``` go
package main

import (
    "github.com/jmattheis/go-packr-swagger-ui"
    "net/http"
)

func main() {
   box := swaggerui.GetBox()
   router := gin.New()
   g.GET("/docs/*any", gin.WrapH(http.StripPrefix("/docs/", http.FileServer(box))))
   g.GET("/swagger", yourSwaggerDefinition)
   router.Run()
}
```