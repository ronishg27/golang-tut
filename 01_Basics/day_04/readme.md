# GoLang Basics - Day 04

### Topics Covered

- Slice
- Array
- Map

## Array

- Collection of elements of the same type.
- Fixed size.
- Array are passed by value.

```go

arr:= [5]int{1, 2, 3, 4, 5}

```

## Slice

- Also same as array but can be resized.
- Slice are passed by reference.

```go
slc := []int{1, 2, 3, 4, 5}

```

when size is given when declaring it is array (fixed size) otherwise it is slice (dynamic size).

```go
arr:= [5]int{1, 2, 3, 4, 5} // array - fixed size
slc := []int{1, 2, 3, 4, 5} // slice - dynamic size
```

### Copying a slice

- copying with reference

```go
copySlice := slc
```

- copying with value
  - copySlice is a new slice with same elements as slc
  - slc is not changed

```go
// copyslice := CopySlice(slc)
copySlice := make([]int, len(slc))
copy(copySlice, slc)
```

## Map

- Collection of key-value pairs.
- Map are passed by value.

```go
person := map[key_type]value_type{
	key1: value1,
	key2: value2,
}

```

### Accessing Map elements

```go
fmt.Println(person["name"])
```

### Merging two maps

- two maps can be merged using `copy()` function

```go
mergedMap := make(map[string]any)
maps.Copy(mergedMap, map1)
maps.Copy(mergedMap, map2)
```
