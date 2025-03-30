# GoLang Basics - Day 02

# Topics Covered

- Constants
- Variables
- Data Types

Constants

Constants are values that cannot be changed.

```go
const (
	PI = 3.14159265358979323846
	E  = 2.71828182845904523536
)

println(PI)
println(E)
```

Variables

Variables are used to store values.

```go
var a int
a = 10
println(a)

var b float64
b = 10.5
println(b)

var c string
c = "Hello World"
println(c)

var d bool
d = true
println(d)

var e byte
e = 10
println(e)
```

## Data Types

GoLang has the following data types

1.  Numeric Types

- **Integers**: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- **Floating-point**: `float32`, `float64`
- **Complex Numbers**: `complex64`, `complex128`

2.  Boolean Type

- `bool`

3.  String Type

- `string`

4.  Array Type

- Fixed-size collection: `[size]Type`

5.  Slice Type

- Dynamic-sized array: `[]Type`

6.  Map Type

- Key-value store: `map[keyType]valueType`

7.  Struct Type

- Custom data structure: `struct`

8.  Pointer Type

- Stores memory address: `*Type`

9.  Interface Type

- Defines behavior: `interface {}`

10. Function Type

- First-class function: `func(params) returnType`

11. Channel Type

- Communication between goroutines: `chan Type`

12. Error Type

- Represents an error: `error`

13. Rune Type

- Unicode character (alias for `int32`): `rune`

14. Complex Type

- Complex number support: `complex64`, `complex128`
