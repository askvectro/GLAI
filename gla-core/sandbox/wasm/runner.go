package wasm
    │   │   import (
    │   │       "context"
    │   │       "github.com/tetratelabs/wazero"
    │   │   )
    │   │   func ExecuteWASM(module []byte) error {
    │   │       ctx := context.Background()
    │   │       runtime := wazero.NewRuntime(ctx)
    │   │       defer runtime.Close(ctx)
    │   │       _, err := runtime.InstantiateModuleFromBinary(ctx, module)
    │   │       return err
    │   │   }