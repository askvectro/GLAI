adapter := adapters.NewOpenAIAdapter("your-api-key", "gpt-4")
filters := []pipeline.Filter{
    pipeline.NewSimpleFilter([]string{"panic", "unsafe"}),
}
formatter := pipeline.NewGoFormatter()

gen := pipeline.NewGenerator(adapter, filters, formatter)
output, err := gen.Run("Write a Go function that reverses a string")
if err != nil {
    log.Fatal(err)
}
fmt.Println(output)
