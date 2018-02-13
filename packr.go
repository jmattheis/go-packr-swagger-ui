package swaggerui

import "github.com/gobuffalo/packr"

// GetBox returns the box with the swagger ui files
func GetBox() packr.Box {
	return packr.NewBox("./swagger-ui/dist")
}
