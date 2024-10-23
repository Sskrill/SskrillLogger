My custom logger , package **slog** is taken as a basis.

**How create** 
```go
logger:=NewCSLogger(opts *slog.HandlerOptions)
```
**How use**
```go
logger.PrintInfo("Some Info","someArgs")
```

