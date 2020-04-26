package readConfig

import (
	"github.com/magiconair/properties"
)

// get property by name
func ParseField(name string, props *properties.Properties) string {
	result, ok := props.Get(name)
	if ok {
		return result
	} else {
		return ""
	}
}
