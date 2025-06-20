# Day 05

## Topics Covered

- Struct
- Interface

## Struct

Struct is a composite data type that allows grouping of data fields of different types into a single entity.

Structs are passed by value.

```go
type Student struct {
	Name string
	Age  int
}
```

## Interface

Interface is a collection of methods.

Interface is passed by value.

```go
type Shape interface {
	Area() float64
	Perimeter() float64
}
```

## Struct and Interface

Struct and Interface can be used together.

```go
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length  float64
	Width   float64
	Shape
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*(r.Length + r.Width)
}
```
