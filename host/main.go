package main

import (
	"context"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
)

func main() {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: "lib.wasm",
				Name: "lib",
			},
			extism.WasmFile{
				Path: "main.wasm",
				Name: "main",
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	_, out, err := plugin.Call("greet", []byte("mo"))
	if err != nil {
		fmt.Printf("Failed to call plugin: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Output: %s\n", out)
}
