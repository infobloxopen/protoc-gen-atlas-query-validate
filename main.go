package main

import (
	"github.com/gogo/protobuf/vanity/command"
	"github.com/infobloxopen/protoc-gen-atlas-query-validate/plugin"
)

func main() {
	plugin := &plugin.QueryValidatePlugin{}
	response := command.GeneratePlugin(command.Read(), plugin, ".pb.atlas.query.validate.go")
	plugin.CleanFiles(response)
	command.Write(response)
}
