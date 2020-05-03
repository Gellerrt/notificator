package config

import (
	"github.com/magiconair/properties"
)

// read .properties file
func ReadPropsConfig(path string, enc properties.Encoding) *properties.Properties {
	props, err := properties.LoadFile(path, enc)
	if err != nil {
		return nil
	} else {
		return props
	}
}
