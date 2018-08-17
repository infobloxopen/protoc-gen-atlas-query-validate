package main

import (
	"github.com/gogo/protobuf/vanity/command"
	"github.com/infobloxopen/protoc-gen-atlas-query-perm/plugin"
)

func main() {
	plugin := &plugin.PermPlugin{}
	response := command.GeneratePlugin(command.Read(), plugin, ".pb.perm.go")
	plugin.CleanFiles(response)
	command.Write(response)
}
