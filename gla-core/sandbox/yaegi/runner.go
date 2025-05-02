package yaegi
    │   │   import (
    │   │       "context"
    │   │       "time"
    │   │       "github.com/traefik/yaegi/interp"
    │   │       "github.com/traefik/yaegi/stdlib"
    │   │   )
    │   │   func Execute(code string, timeout time.Duration) (string, error) {
    │   │       ctx, cancel := context.WithTimeout(context.Background(), timeout)
    │   │       defer cancel()
    │   │       i := interp.New(interp.Options{})
    │   │       i.Use(stdlib.Symbols)
    │   │       _, err := i.EvalWithContext(ctx, code)
    │   │       if err != nil {
    │   │           return "", err
    │   │       }
    │   │       return "Execution Successful", nil
    │   │   }