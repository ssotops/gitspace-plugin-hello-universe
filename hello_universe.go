package main

import (
	"fmt"
)

type HelloUniversePlugin struct{}

func (h HelloUniversePlugin) Run() error {
	fmt.Println("Hello, World!")
	return nil
}

func (h HelloUniversePlugin) Name() string {
	return "hello_world"
}

func (h HelloUniversePlugin) Version() string {
	return "1.0.0"
}

// Export the plugin
var Plugin HelloUniversePlugin
