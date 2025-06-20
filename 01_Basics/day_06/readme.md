# Golang Basics - Day 06

## Topics Covered

- Error Handling in Go

### Error in GoLang

Error handling is a crucial aspect of Go programming. Unlike many other languages that use exceptions, Go uses explicit error handling through return values.

#### 1. Error Interface

```go
type error interface {
    Error() string
}
```

- The `error` interface is a built-in interface in Go
- It requires only one method: `Error() string`
- This method returns a string describing the error

#### 2. Creating Custom Errors

There are several ways to create errors in Go:

1. Using `errors.New()`:

```go
err := errors.New("something went wrong")
```

2. Using `fmt.Errorf()`:

```go
err := fmt.Errorf("failed to process %v: %v", input, reason)
```

3. Creating custom error types:

```go
type ValidationError struct {
    Field string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

#### 3. Error Handling Patterns

1. Basic Error Checking:

```go
result, err := someFunction()
if err != nil {
    return err
}
```

2. Error Wrapping (Go 1.13+):

```go
import "errors"

if err != nil {
    return errors.Wrap(err, "failed to process request")
}
```

3. Error Type Assertion:

```go
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    // Handle validation error specifically
}
```

#### 4. Best Practices

1. Always check errors:

```go
// Bad
result, _ := someFunction()

// Good
result, err := someFunction()
if err != nil {
    return err
}
```

2. Return errors early:

```go
// Good
if err := validate(); err != nil {
    return err
}
if err := process(); err != nil {
    return err
}
```

3. Add context to errors:

```go
// Bad
return err

// Good
return fmt.Errorf("failed to process request: %w", err)
```

#### 5. Common Error Handling Patterns

1. Sentinel Errors:

```go
var (
    ErrNotFound = errors.New("not found")
    ErrTimeout  = errors.New("timeout")
)
```

2. Error Types:

```go
type NotFoundError struct {
    Resource string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s not found", e.Resource)
}
```

3. Error Handling with defer:

```go
func processFile() error {
    file, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer file.Close()
    // Process file
    return nil
}
```

#### 6. Testing Errors

```go
func TestProcess(t *testing.T) {
    err := Process()
    if err != nil {
        t.Errorf("Process() error = %v", err)
        return
    }
}
```

#### 7. Common Pitfalls

1. Ignoring errors
2. Not adding context to errors
3. Using panic for error handling
4. Not handling specific error types
5. Not cleaning up resources in error cases

#### 8. Error Handling Libraries

1. `github.com/pkg/errors` - Enhanced error handling
2. `github.com/hashicorp/errwrap` - Error wrapping utilities
3. `github.com/rotisserie/eris` - Error handling and logging

Remember: Good error handling is crucial for building robust and maintainable Go applications. Always handle errors explicitly and provide meaningful context when returning errors.
