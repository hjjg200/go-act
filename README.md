# go-act
Assert, catch, and try

## Example

```go
import (
    . "github.com/hjjg200/go-act" // import . so that you don't have to write "act."
)
func Foo() (bar int, err error) {
    defer Catch(&err) // Should the function panic, the error that caused panic, the error will be captured into err
    
    // Try
    n, err := Foo2()
    Try(err) // Panics if err is not nil
    Try(Foo3()) // If the function returns an error only, it can be used this way
    
    // Assert
    val, ok := SomeMap[key]
    Assert(ok, "key does not exist in the map!")
    
    return
}
```
